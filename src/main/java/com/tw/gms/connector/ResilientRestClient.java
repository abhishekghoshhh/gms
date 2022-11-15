package com.tw.gms.connector;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.client.circuitbreaker.CircuitBreaker;
import org.springframework.cloud.client.circuitbreaker.CircuitBreakerFactory;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.net.URI;

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
        //implement hystrix and fallback and other type of httpException
        CircuitBreaker circuitBreaker = getCircuitBreaker(hystrixKey);
        return circuitBreaker.run(() -> {
            return restTemplate.exchange(url, method, httpEntity, type);
        }, throwable -> {
            return ResponseEntity.ok(defaultResponse);
        });
    }

    public <T> ResponseEntity<T> exchange(String hystrixKey, URI url, HttpMethod method, HttpEntity<?> httpEntity, Class<T> type) throws RestCallException {
        //implement hystrix and fallback and other type of httpException
        CircuitBreaker circuitBreaker = getCircuitBreaker(hystrixKey);
        try{
            return circuitBreaker.run(() -> {
                return restTemplate.exchange(url, method, httpEntity, type);
            }, throwable -> {
                throw new RuntimeException(throwable);
            });
        }catch (RuntimeException exception){
            throw new RestCallException(exception.getCause());
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
