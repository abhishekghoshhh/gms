package com.tw.gms.connector;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.junit.jupiter.MockitoExtension;

import static org.junit.jupiter.api.Assertions.assertNotNull;

@ExtendWith(MockitoExtension.class)
public class RestTemplateProviderTest {
    @InjectMocks
    RestTemplateProvider restTemplateProvider;

    @Test
    public void templateWithSSL() {
        assertNotNull(restTemplateProvider.templateWithSSL());
    }

    @Test
    public void templateWithoutSSL() {
        assertNotNull(restTemplateProvider.templateWithoutSSL());
    }
}