package com.tw.gms.connector;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.junit.jupiter.MockitoExtension;

import static org.junit.jupiter.api.Assertions.assertNotNull;

@ExtendWith(MockitoExtension.class)
public class HostNameVerificationProviderTest {

    @InjectMocks
    HostNameVerificationProvider hostNameVerificationProvider;

    @Test
    public void verifyWithTrue() {
        assertNotNull(hostNameVerificationProvider.hostnameVerifier("true"));
    }

    @Test
    public void verifyWithFalse() {
        assertNotNull(hostNameVerificationProvider.hostnameVerifier("false"));
    }
}