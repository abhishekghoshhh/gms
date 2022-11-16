package com.tw.gms.connector;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.client.circuitbreaker.CircuitBreaker;
import org.springframework.cloud.client.circuitbreaker.CircuitBreakerFactory;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.client.ResourceAccessException;
import org.springframework.web.client.RestTemplate;

import java.net.SocketException;
import java.net.SocketTimeoutException;
import java.net.URI;
import java.util.concurrent.TimeoutException;

@Service
public class ResilientRestClient {

    RestTemplate restTemplate;

    @Autowired
    CircuitBreakerFactory circuitBreakerFactory;

    public ResilientRestClient(@Value("${template.withSsl}") String restTemplateType,
                               @Autowired RestTemplate templateWithSSL,
                               @Autowired RestTemplate templateWithoutSSL) {
        if ("false".equalsIgnoreCase(restTemplateType)) {
            restTemplate = templateWithoutSSL;
        } else {
            restTemplate = templateWithSSL;
        }
    }

    public <T> ResponseEntity<T> exchange(String hystrixKey, URI url, HttpMethod method, HttpEntity<?> httpEntity, Class<T> type, T defaultResponse) {
        CircuitBreaker circuitBreaker = getCircuitBreaker(hystrixKey);
        return circuitBreaker.run(() -> {
            return restTemplate.exchange(url, method, httpEntity, type);
        }, throwable -> {
            return ResponseEntity.ok(defaultResponse);
        });
    }

    public <T> ResponseEntity<T> exchange(String hystrixKey, URI url, HttpMethod method, HttpEntity<?> httpEntity, Class<T> type) throws RestCallException {
        CircuitBreaker circuitBreaker = getCircuitBreaker(hystrixKey);
        try {
            return circuitBreaker.run(() -> {
                return restTemplate.exchange(url, method, httpEntity, type);
            }, throwable -> {
                throw new RuntimeException(throwable);
            });
        } catch (RuntimeException exception) {
            return buildException(exception);
        }
    }

    private <T> T buildException(RuntimeException exception) throws RestCallException {
        Throwable cause = exception.getCause();
        if (cause instanceof HttpClientErrorException) {
            HttpClientErrorException httpClientErrorException = (HttpClientErrorException) cause;
            throw new RestCallException("Unauthorized Access",
                    httpClientErrorException.getStatusCode(),
                    httpClientErrorException.getResponseBodyAsString());
        } else {
            throw new RestCallException("Internal server error",
                    null != cause ? cause.getMessage() : exception.getMessage());
        }
    }

    private CircuitBreaker getCircuitBreaker(String hystrixKey) {
        CircuitBreaker circuitBreaker = circuitBreakerFactory.create(hystrixKey);
        if (null == circuitBreaker) {
            circuitBreaker = circuitBreakerFactory.create("default");
        }
        return circuitBreaker;
    }
}
