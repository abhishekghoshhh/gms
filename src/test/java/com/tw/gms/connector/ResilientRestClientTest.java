package com.tw.gms.connector;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mockito;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.cloud.client.circuitbreaker.CircuitBreaker;
import org.springframework.cloud.client.circuitbreaker.CircuitBreakerFactory;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.ReflectionUtils;
import org.springframework.web.client.RestTemplate;

import java.lang.reflect.Field;
import java.net.URI;

import static org.junit.jupiter.api.Assertions.assertEquals;

@ExtendWith(MockitoExtension.class)
public class ResilientRestClientTest {

    @InjectMocks
    ResilientRestClient resilientRestClient;

    @Test
    public void exchange() throws Exception {
        RestTemplate restTemplate = Mockito.mock(RestTemplate.class);
        CircuitBreakerFactory circuitBreakerFactory = Mockito.mock(CircuitBreakerFactory.class);
        CircuitBreaker circuitBreaker = Mockito.mock(CircuitBreaker.class);

        setFieldByReflection(ResilientRestClient.class, resilientRestClient, "restTemplate", restTemplate);
        setFieldByReflection(ResilientRestClient.class, resilientRestClient, "circuitBreakerFactory", circuitBreakerFactory);
        URI uri = URI.create("http://localhost");
        HttpMethod method = HttpMethod.POST;
        HttpEntity<?> entity = new HttpEntity<>(new LinkedMultiValueMap<>());
        ResponseEntity<String> response = ResponseEntity.ok("Hello world");

        Mockito.when(circuitBreakerFactory.create(Mockito.anyString())).thenReturn(circuitBreaker);
        Mockito.when(circuitBreaker.run(Mockito.any(),Mockito.any())).thenReturn(response);
//        Mockito.when(restTemplate.exchange(uri, method, entity, String.class)).thenReturn(response);
        assertEquals(response, resilientRestClient.exchange("default", uri, method, entity, String.class));
    }

    private void setFieldByReflection(Class<?> classType, Object object, String fieldName, Object fieldValue) throws NoSuchFieldException {
        Field field = classType.getDeclaredField(fieldName);
        field.setAccessible(true);
        ReflectionUtils.setField(field, object, fieldValue);
    }
}