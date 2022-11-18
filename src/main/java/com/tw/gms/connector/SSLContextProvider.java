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
    Logger log = LoggerFactory.getLogger(SSLContextProvider.class);

    @Bean
    public SSLContext sslContext(@Autowired Environment environment/*, @Autowired CertSignatureVerifier certSignatureVerifier*/)
            throws IOException, CertificateException, NoSuchAlgorithmException, KeyStoreException, KeyManagementException {
        String withSsl = environment.getProperty("rest-template.withSsl", TRUE);
        if (TRUE.equalsIgnoreCase(withSsl)) {
            TrustStrategy trustStrategy = (X509Certificate[] x509Certificates, String authType) -> {
//                try {
//                    return certSignatureVerifier.verifyCertChainSignatures(x509Certificates);
//                } catch (InvalidKeyException | NoSuchAlgorithmException | NoSuchProviderException | SignatureException |
//                         IllegalBlockSizeException | BadPaddingException | NoSuchPaddingException | IOException e) {
//                    log.error("error occurred during certificate verification {}", e.getMessage());
//                    return false;
//                }
                return true;
            };
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
            TrustStrategy trustStrategy = (X509Certificate[] x509Certificates, String authType) -> {
                return true;
            };
            return SSLContextBuilder
                    .create()
                    .loadTrustMaterial(trustStrategy)
                    .build();
        }
    }
}
