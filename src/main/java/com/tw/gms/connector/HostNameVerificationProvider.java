package com.tw.gms.connector;

import org.apache.http.conn.ssl.DefaultHostnameVerifier;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.net.ssl.HostnameVerifier;
import javax.net.ssl.SSLSession;

@Configuration
public class HostNameVerificationProvider {
    public static final String TRUE = "true";
    Logger log = LoggerFactory.getLogger(HostNameVerificationProvider.class);

    @Bean
    public HostnameVerifier hostnameVerifier(@Value("${rest-template.verifyHostName:false}") String verifyHostName) {
        log.debug("verifyHostName flag is {}", verifyHostName);
        if (TRUE.equalsIgnoreCase(verifyHostName)) {
            return new DefaultHostnameVerifier();
        } else {
            return (String hostName, SSLSession sslSession) -> true;
        }
    }
}
