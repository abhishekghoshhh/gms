package com.tw.gms.connector;

import org.apache.http.client.config.RequestConfig;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClientBuilder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;
import org.springframework.http.client.ClientHttpRequestFactory;
import org.springframework.http.client.HttpComponentsClientHttpRequestFactory;
import org.springframework.web.client.RestTemplate;

import java.net.http.HttpClient;

@Configuration
public class RestTemplateProvider {

    @Bean
//    @Qualifier("templateWithSSL")
//    @Primary
    public RestTemplate templateWithSSL(@Autowired RestTemplateProperties restTemplateProperties) {
        //Implement settings for with ssl
        // add rest template properties
        RestTemplate restTemplate = new RestTemplate();
        restTemplate.setRequestFactory(buildRequestFactory(restTemplateProperties));
        return restTemplate;
    }

    @Bean
//    @Qualifier("templateWithoutSSL")
    public RestTemplate templateWithoutSSL(@Autowired RestTemplateProperties restTemplateProperties) {
        //Implement settings for without ssl
        // add rest template properties
        RestTemplate restTemplate = new RestTemplate();
        restTemplate.setRequestFactory(buildRequestFactory(restTemplateProperties));
        return restTemplate;
    }

    private ClientHttpRequestFactory buildRequestFactory(RestTemplateProperties restTemplateProperties) {
        CloseableHttpClient httpClient = HttpClientBuilder.create().build();
        HttpComponentsClientHttpRequestFactory requestFactory = new HttpComponentsClientHttpRequestFactory(httpClient);
        requestFactory.setConnectionRequestTimeout(restTemplateProperties.getConnectionRequestTimeout());
        requestFactory.setConnectTimeout(restTemplateProperties.getConnectTimeout());
        requestFactory.setReadTimeout(restTemplateProperties.getReadTimeout());
        return requestFactory;
    }
}
