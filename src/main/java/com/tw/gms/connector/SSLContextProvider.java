package com.tw.gms.connector;

import org.apache.http.conn.ssl.TrustStrategy;
import org.apache.http.ssl.SSLContextBuilder;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.env.Environment;
import org.springframework.util.ResourceUtils;

import javax.net.ssl.SSLContext;
import java.io.IOException;
import java.security.KeyManagementException;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.cert.CertificateException;

@Configuration
public class SSLContextProvider {

    public static final String TRUE = "true";
    public static final String FALSE = "false";
    Logger log = LoggerFactory.getLogger(SSLContextProvider.class);

    @Bean
    public SSLContext sslContext(@Autowired Environment environment)
            throws IOException, CertificateException, NoSuchAlgorithmException, KeyStoreException, KeyManagementException {
        String withSsl = environment.getProperty("rest-template.withSsl", TRUE);
        log.info("withSsl flag value {}", withSsl);
        String validateCertificateChain = environment.getProperty("rest-template.validateCertificateChain", FALSE);
        log.info("validateCertificateChain flag value {}", validateCertificateChain);
        TrustStrategy trustStrategy = (x509Certificates, authType) -> true;
        if (TRUE.equalsIgnoreCase(withSsl)) {
            String location = environment.getProperty("server.ssl.key-store");
            String pass = environment.getProperty("server.ssl.key-store-password");
            log.info("location of the certificate {}", location);
            if (null == location || location.isBlank() || null == pass || pass.isBlank()) {
                throw new RuntimeException("keystore/password should not be empty");
            }
            return SSLContextBuilder
                    .create()
                    .loadTrustMaterial(ResourceUtils.getFile(location), pass.toCharArray(), trustStrategy)
                    .build();
        } else {
            return SSLContextBuilder
                    .create()
                    .loadTrustMaterial(trustStrategy)
                    .build();
        }
    }
}
