package com.tw.gms.service.impl;

import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

@Service
public class CustomRestTemplate {

    @Bean
    public RestTemplate templateWithSSL() {
        return new RestTemplate();
    }

    @Bean
    public RestTemplate templateWithoutSSL() {
        return new RestTemplate();
    }
}
