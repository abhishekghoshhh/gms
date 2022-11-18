package com.tw.gms.connector;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.core.env.Environment;

import java.io.IOException;
import java.security.KeyManagementException;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.cert.CertificateException;

import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class SSLContextProviderTest {

    @InjectMocks
    SSLContextProvider sSLContextProvider;

    @Mock
    Environment environment;

    @Test
    public void sslContextWithSSL()
            throws CertificateException, IOException, NoSuchAlgorithmException, KeyStoreException, KeyManagementException {
        when(environment.getProperty(eq("rest-template.withSsl"), eq("true")))
                .thenReturn("true");
        when(environment.getProperty(eq("server.ssl.key-store")))
                .thenReturn("src/test/resources/springboot.p12");
        when(environment.getProperty(eq("server.ssl.key-store-password")))
                .thenReturn("password");
        assertNotNull(sSLContextProvider.sslContext(environment));
    }

    @Test
    public void sslContextWithSSlWithoutLocationAndPassword()
            throws CertificateException, IOException, NoSuchAlgorithmException, KeyStoreException, KeyManagementException {
        when(environment.getProperty(eq("rest-template.withSsl"), eq("true")))
                .thenReturn("true");
        assertThrows(RuntimeException.class,()->sSLContextProvider.sslContext(environment));
    }
    @Test
    public void sslContextWithoutSSl()
            throws CertificateException, IOException, NoSuchAlgorithmException, KeyStoreException, KeyManagementException {
        when(environment.getProperty(eq("rest-template.withSsl"), eq("true")))
                .thenReturn("false");
        assertNotNull(sSLContextProvider.sslContext(environment));
    }
}