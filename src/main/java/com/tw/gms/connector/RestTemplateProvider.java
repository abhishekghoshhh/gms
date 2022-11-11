package com.tw.gms.connector;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;

@Configuration
public class RestTemplateProvider {

    @Bean
    public RestTemplate templateWithSSL() {
        //Implement settings for with ssl
        // add rest template properties
        return new RestTemplate();
    }

    @Bean
    public RestTemplate templateWithoutSSL() {
        //Implement settings for without ssl
        // add rest template properties
        return new RestTemplate();
    }
}
