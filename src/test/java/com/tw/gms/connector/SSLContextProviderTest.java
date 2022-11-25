package com.tw.gms.connector;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.core.env.Environment;

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
    public void sslContextWithSSL() throws Exception {
        when(environment.getProperty(eq("rest-template.withSsl"), eq("true")))
                .thenReturn("true");
        when(environment.getProperty(eq("server.ssl.key-store")))
                .thenReturn("src/test/resources/springboot.p12");
        when(environment.getProperty(eq("server.ssl.key-store-password")))
                .thenReturn("password");
        when(environment.getProperty(eq("rest-template.validateCertificateChain"), eq("false")))
                .thenReturn("true");
        assertNotNull(sSLContextProvider.sslContext(environment));
    }

    @Test
    public void sslContextWithSSlWithoutLocationAndPassword() throws Exception {
        when(environment.getProperty(eq("rest-template.withSsl"), eq("true")))
                .thenReturn("true");

        assertThrows(RuntimeException.class, () -> sSLContextProvider.sslContext(environment));
    }

    @Test
    public void sslContextWithoutSSl() throws Exception {
        when(environment.getProperty(eq("rest-template.withSsl"), eq("true")))
                .thenReturn("false");
        assertNotNull(sSLContextProvider.sslContext(environment));
    }
}