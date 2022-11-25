package com.tw.gms.connector;

import org.apache.logging.log4j.ThreadContext;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.slf4j.MDC;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.client.circuitbreaker.CircuitBreaker;
import org.springframework.cloud.client.circuitbreaker.CircuitBreakerFactory;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.client.RestTemplate;

import java.net.URI;
import java.util.HashMap;
import java.util.Map;

@Service
public class ResilientRestClient {

    @Autowired
    RestTemplate restTemplate;

    @Autowired
    CircuitBreakerFactory circuitBreakerFactory;

    Logger log = LoggerFactory.getLogger(ResilientRestClient.class);

    public <T> ResponseEntity<T> exchange(String hystrixKey, URI url, HttpMethod method, HttpEntity<?> httpEntity, Class<T> type, T defaultResponse) {
        log.debug("hystrixKey is {}", hystrixKey);
        CircuitBreaker circuitBreaker = getCircuitBreaker(hystrixKey);
        Map<String, String> contextMap = getContextMap();
        return circuitBreaker.run(() -> {
            MDC.setContextMap(contextMap);
            log.info("calling to {} ", url);
            log.info("headers are {}", httpEntity.getHeaders());
            ResponseEntity<T> responseEntity = restTemplate.exchange(url, method, httpEntity, type);
            log.info("response is {}", responseEntity.getBody());
            return responseEntity;
        }, throwable -> {
            MDC.setContextMap(contextMap);
            log.info("returning the default response {}", defaultResponse);
            return ResponseEntity.ok(defaultResponse);
        });
    }

    public <T> ResponseEntity<T> exchange(String hystrixKey, URI url, HttpMethod method, HttpEntity<?> httpEntity, Class<T> type) throws RestCallException {
        log.debug("hystrixKey is {}", hystrixKey);
        CircuitBreaker circuitBreaker = getCircuitBreaker(hystrixKey);
        Map<String, String> contextMap = getContextMap();
        try {
            return circuitBreaker.run(() -> {
                MDC.setContextMap(contextMap);
                log.info("calling to {} ", url);
                log.info("headers are {}", httpEntity.getHeaders());
                ResponseEntity<T> responseEntity = restTemplate.exchange(url, method, httpEntity, type);
                log.info("response is {}", responseEntity.getBody());
                return responseEntity;
            }, throwable -> {
                log.error("throwing exception from rest-template {}", throwable.getMessage());
                throwable.printStackTrace();
                throw new RuntimeException(throwable);
            });
        } catch (RuntimeException exception) {
            return buildException(exception);
        }
    }

    private Map<String, String> getContextMap() {
        if (null != MDC.getCopyOfContextMap()) return MDC.getCopyOfContextMap();
        if (null != ThreadContext.getContext()) return ThreadContext.getContext();
        return new HashMap<>();
    }

    private <T> T buildException(RuntimeException exception) throws RestCallException {
        Throwable cause = exception.getCause();
        if (cause instanceof HttpClientErrorException) {
            HttpClientErrorException httpClientErrorException = (HttpClientErrorException) cause;
            throw new RestCallException("Unauthorized Access",
                    httpClientErrorException.getStatusCode(),
                    httpClientErrorException.getResponseBodyAsString());
        }
        throw new RestCallException("Internal server error",
                null != cause ? cause.getMessage() : exception.getMessage());
    }

    private CircuitBreaker getCircuitBreaker(String hystrixKey) {
        CircuitBreaker circuitBreaker = circuitBreakerFactory.create(hystrixKey);
        if (null == circuitBreaker) {
            circuitBreaker = circuitBreakerFactory.create("default");
        }
        return circuitBreaker;
    }
}
