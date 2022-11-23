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
import java.security.cert.X509Certificate;

@Configuration
public class SSLContextProvider {

    public static final String TRUE = "true";
    public static final String FALSE = "false";
    Logger log = LoggerFactory.getLogger(SSLContextProvider.class);

    @Bean
    public SSLContext sslContext(@Autowired Environment environment, @Autowired CertSignatureVerifier certSignatureVerifier)
            throws IOException, CertificateException, NoSuchAlgorithmException, KeyStoreException, KeyManagementException {
        String withSsl = environment.getProperty("rest-template.withSsl", TRUE);
        String validateCertificateChain = environment.getProperty("rest-template.validateCertificateChain", FALSE);
        TrustStrategy trustStrategy = getTrustStrategy(validateCertificateChain, certSignatureVerifier);
        if (TRUE.equalsIgnoreCase(withSsl)) {
            String location = environment.getProperty("server.ssl.key-store");
            String pass = environment.getProperty("server.ssl.key-store-password");
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

    private TrustStrategy getTrustStrategy(String validateCertificateChain, CertSignatureVerifier certSignatureVerifier) {
        if (TRUE.equalsIgnoreCase(validateCertificateChain)) {
            return (X509Certificate[] x509Certificates, String authType) -> {
                try {
                    log.debug("authType is {}", authType);
                    return certSignatureVerifier.verifyCertChainSignatures(x509Certificates);
                } catch (Exception e) {
                    log.error("error while validating certificate chain is {}", e.getMessage());
                    throw new RuntimeException(e);
                }
            };
        }
        return (x509Certificates, authType) -> true;
    }


}
