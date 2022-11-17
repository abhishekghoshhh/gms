package com.tw.gms.connector;

import org.apache.http.conn.ssl.TrustStrategy;
import org.apache.http.ssl.SSLContextBuilder;
import org.apache.http.ssl.SSLContexts;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.env.Environment;

import javax.net.ssl.SSLContext;
import java.io.IOException;
import java.security.KeyManagementException;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;

import static org.springframework.util.ResourceUtils.getFile;

@Configuration
public class SSLContextProvider {

    public static final String TRUE = "true";
    Logger log = LoggerFactory.getLogger(SSLContextProvider.class);

    @Bean
    public SSLContext sslContext(@Autowired Environment environment)
            throws IOException, CertificateException, NoSuchAlgorithmException, KeyStoreException, KeyManagementException {
        String withSsl = environment.getProperty("rest-template.withSsl", TRUE);
        if (TRUE.equalsIgnoreCase(withSsl)) {
            //String location = "D:\\ssl_server.jks";
            String location = environment.getProperty("server.ssl.key-store");
            String pass = environment.getProperty("key-store-password");
            return SSLContextBuilder
                    .create()
                    .loadTrustMaterial(getFile(location), pass.toCharArray())
                    .build();

        } else {
            TrustStrategy acceptingTrustStrategy = (X509Certificate[] x509Certificates, String authType) -> {
                return true;
            };
            return SSLContexts.custom()
                    .loadTrustMaterial(null, acceptingTrustStrategy)
                    .build();
        }
    }
}
