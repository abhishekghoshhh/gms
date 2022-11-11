package com.tw.gms.connector;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.lang.Nullable;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import org.springframework.web.client.RestTemplate;

import java.net.URI;

@Service
public class ResilientRestClient {

    RestTemplate restTemplate;

    public ResilientRestClient(@Value("${template.withSsl}") String restTemplateType, @Autowired RestTemplate templateWithSSL, @Autowired RestTemplate templateWithoutSSL) {
        if ("false".equalsIgnoreCase(restTemplateType.trim())) {
            restTemplate = templateWithoutSSL;
        } else {
            restTemplate = templateWithSSL;
        }
    }

    public <T> ResponseEntity<T> exchange(String hystrixKey, URI url, HttpMethod method, HttpEntity<?> httpEntity, Class<T> type) {
        //implement hstrix and fallback and other type of httpException
        return restTemplate.exchange(url, method, httpEntity, type);
    }
}
