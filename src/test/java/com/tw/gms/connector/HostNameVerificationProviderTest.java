package com.tw.gms.connector;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.junit.jupiter.MockitoExtension;

import javax.net.ssl.SSLSession;

import static com.tw.gms.utils.TestUtils.setFieldByReflection;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.Mockito.mock;

@ExtendWith(MockitoExtension.class)
public class HostNameVerificationProviderTest {

    @InjectMocks
    HostNameVerificationProvider hostNameVerificationProvider;

    @Test
    public void verifyWithTrue() throws NoSuchFieldException {
        setFieldByReflection(HostNameVerificationProvider.class, hostNameVerificationProvider, "verifyHostName", "true");
        assertTrue(hostNameVerificationProvider.verify("localhost", mock(SSLSession.class)));
    }

    @Test
    public void verifyWithFalse()  throws NoSuchFieldException {
        setFieldByReflection(HostNameVerificationProvider.class, hostNameVerificationProvider, "verifyHostName", "false");
        assertTrue(hostNameVerificationProvider.verify("localhost", mock(SSLSession.class)));
    }
}