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
        RestTemplateProperties restTemplateProperties = new RestTemplateProperties();
        assertNotNull(restTemplateProvider.templateWithSSL(restTemplateProperties));
    }

    @Test
    public void templateWithoutSSL() {
        RestTemplateProperties restTemplateProperties = new RestTemplateProperties();
        assertNotNull(restTemplateProvider.templateWithoutSSL(restTemplateProperties));
    }
}