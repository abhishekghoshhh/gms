package com.tw.gms.connector;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.net.ssl.HostnameVerifier;
import javax.net.ssl.SSLSession;

@Component
public class HostNameVerificationProvider implements HostnameVerifier {
    private final String verifyHostName;

    public HostNameVerificationProvider(@Value("${rest-template.verifyHostName:false}") String verifyHostName) {
        this.verifyHostName = verifyHostName;
    }

    @Override
    public boolean verify(String hostName, SSLSession sslSession) {
        if ("true".equalsIgnoreCase(verifyHostName)) {
            //TODO implement this
            //return "localhost".equalsIgnoreCase(hostName) || "127.0.0.1".equals(hostName);
        }
        return true;
    }
}
