package com.tw.gms.connector;

import org.apache.http.conn.ssl.SSLConnectionSocketFactory;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.client.HttpComponentsClientHttpRequestFactory;
import org.springframework.web.client.RestTemplate;

import javax.net.ssl.SSLContext;

@Configuration
public class RestTemplateProvider {

    Logger log = LoggerFactory.getLogger(RestTemplateProvider.class);
    private RestTemplate restTemplate;

    @Bean
    public RestTemplate restTemplate(@Autowired RestTemplateProperties restTemplateProperties,
                                     @Autowired SSLContext sslContext,
                                     @Autowired HostNameVerificationProvider hostNameVerificationProvider) {
        SSLConnectionSocketFactory socketFactory = new SSLConnectionSocketFactory(sslContext, hostNameVerificationProvider);
        CloseableHttpClient httpClient = HttpClients.custom().setSSLSocketFactory(socketFactory).build();
        HttpComponentsClientHttpRequestFactory requestFactory = new HttpComponentsClientHttpRequestFactory(httpClient);
        addTimeoutSettings(restTemplateProperties, requestFactory);
        RestTemplate restTemplate = new RestTemplate();
        restTemplate.setRequestFactory(requestFactory);
        return restTemplate;
    }

    private void addTimeoutSettings(RestTemplateProperties restTemplateProperties, HttpComponentsClientHttpRequestFactory requestFactory) {
        requestFactory.setConnectionRequestTimeout(restTemplateProperties.getConnectionRequestTimeout());
        requestFactory.setConnectTimeout(restTemplateProperties.getConnectTimeout());
        requestFactory.setReadTimeout(restTemplateProperties.getReadTimeout());
    }

}
