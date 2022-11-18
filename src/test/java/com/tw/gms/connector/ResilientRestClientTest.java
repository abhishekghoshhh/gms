package com.tw.gms.connector;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInstance;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mockito;
import org.mockito.junit.jupiter.MockitoExtension;
import org.slf4j.MDC;
import org.springframework.cloud.client.circuitbreaker.CircuitBreaker;
import org.springframework.cloud.client.circuitbreaker.CircuitBreakerFactory;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.client.RestTemplate;

import java.net.URI;
import java.util.UUID;
import java.util.function.Function;
import java.util.function.Supplier;

import static com.tw.gms.utils.TestUtils.setFieldByReflection;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
@TestInstance(TestInstance.Lifecycle.PER_CLASS)
public class ResilientRestClientTest {

    public static final String FALL_BACK_RESPONSE = "new test response";
    public static final String DEFAULT = "default";
    ResilientRestClient resilientRestClient;

    RestTemplate restTemplate;

    CircuitBreakerFactory circuitBreakerFactory;

    CircuitBreaker circuitBreaker;

    HttpEntity<?> httpEntity;

    URI uri;

    HttpMethod httpMethod;

    @BeforeEach
    void setUp() throws NoSuchFieldException {
        resilientRestClient = new ResilientRestClient();
        restTemplate = mock(RestTemplate.class);
        circuitBreakerFactory = mock(CircuitBreakerFactory.class);
        circuitBreaker = new CircuitBreaker() {
            @Override
            public <T> T run(Supplier<T> toRun, Function<Throwable, T> fallback) {
                try {
                    return toRun.get();
                } catch (Exception ex) {
                    return fallback.apply(ex);
                }
            }
        };
        setFieldByReflection(ResilientRestClient.class, resilientRestClient, "restTemplate", restTemplate);
        setFieldByReflection(ResilientRestClient.class, resilientRestClient, "circuitBreakerFactory", circuitBreakerFactory);
        when(circuitBreakerFactory.create(Mockito.anyString())).thenReturn(circuitBreaker);
        MDC.put("uuid", UUID.randomUUID().toString());


        httpEntity = new HttpEntity<>(new HttpHeaders());
        uri = URI.create("http://localhost");
        httpMethod = HttpMethod.POST;
    }

    @Test
    public void exchange() throws Exception {
        ResponseEntity<String> response = ResponseEntity.ok("Hello world");
        when(restTemplate.exchange(uri, httpMethod, httpEntity, String.class))
                .thenReturn(response);
        assertEquals(response,
                resilientRestClient.exchange(DEFAULT, uri, httpMethod, httpEntity, String.class));
    }

    @Test
    public void exchangeWithException() throws Exception {
        when(restTemplate.exchange(uri, httpMethod, httpEntity, String.class))
                .thenThrow(new RuntimeException("Exception occurred"));
        assertThrows(RestCallException.class,
                () -> resilientRestClient.exchange(DEFAULT, uri, httpMethod, httpEntity, String.class));
    }

    @Test
    public void exchangeWithUnAuthorizedException() throws Exception {
        when(restTemplate.exchange(uri, httpMethod, httpEntity, String.class))
                .thenThrow(mock(HttpClientErrorException.class));
        assertThrows(RestCallException.class,
                () -> resilientRestClient.exchange(DEFAULT, uri, httpMethod, httpEntity, String.class));
    }

    @Test
    public void exchangeWithDefault() throws Exception {
        ResponseEntity<String> response = ResponseEntity.ok("Hello world");
        when(restTemplate.exchange(uri, httpMethod, httpEntity, String.class))
                .thenReturn(response);
        assertEquals(response,
                resilientRestClient.exchange(DEFAULT, uri, httpMethod, httpEntity, String.class, "new test response"));
    }

    @Test
    public void exchangeWithDefaultWithException() throws Exception {
        ResponseEntity<String> response = ResponseEntity.ok(FALL_BACK_RESPONSE);
        when(restTemplate.exchange(uri, httpMethod, httpEntity, String.class))
                .thenThrow(new RuntimeException("Exception occurred"));
        assertEquals(response,
                resilientRestClient.exchange(DEFAULT, uri, httpMethod, httpEntity, String.class, FALL_BACK_RESPONSE));
    }

    @Test
    public void exchangeWithDifferentHystrixKey() throws Exception {
        when(circuitBreakerFactory.create(Mockito.eq("iam"))).thenReturn(null);
        ResponseEntity<String> response = ResponseEntity.ok("Hello world");
        when(restTemplate.exchange(uri, httpMethod, httpEntity, String.class))
                .thenReturn(response);
        assertEquals(response,
                resilientRestClient.exchange("iam", uri, httpMethod, httpEntity, String.class));
    }
}