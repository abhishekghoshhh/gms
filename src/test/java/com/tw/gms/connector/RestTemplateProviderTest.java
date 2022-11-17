package com.tw.gms.connector;

import org.apache.http.ssl.SSLContexts;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.junit.jupiter.MockitoExtension;

import javax.net.ssl.SSLContext;
import java.security.KeyManagementException;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;

import static org.junit.jupiter.api.Assertions.assertNotNull;

@ExtendWith(MockitoExtension.class)
public class RestTemplateProviderTest {
    @InjectMocks
    RestTemplateProvider restTemplateProvider;


    @Test
    public void restTemplate() throws NoSuchAlgorithmException, KeyStoreException, KeyManagementException {
        SSLContext sslContext = SSLContexts.custom()
                .loadTrustMaterial(null, (x509Certificates, authType) -> true)
                .build();
        HostNameVerficationProvider hostNameVerficationProvider = new HostNameVerficationProvider("false");
        RestTemplateProperties restTemplateProperties = new RestTemplateProperties();
        assertNotNull(restTemplateProvider.restTemplate(restTemplateProperties, sslContext, hostNameVerficationProvider));
    }
}