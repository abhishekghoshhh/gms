package com.tw.gms.service.impl;

import com.tw.gms.connector.ResilientRestClient;
import com.tw.gms.model.ProfileResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.*;
import org.springframework.stereotype.Service;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;
import org.springframework.web.util.UriComponentsBuilder;

import java.net.URI;
import java.util.List;

@Service
public class ProfileFetcher {

    @Value("${iam.clientId}")
    private String clientId;

    @Value("${iam.clientSecret}")
    private String clientSecret;

    @Autowired
    ResilientRestClient resilientRestClient;

    public ProfileResponse fetch(String token) {
        URI uri = UriComponentsBuilder.fromHttpUrl("http://localhost:8080/introspect").build().toUri();

        HttpHeaders headers = new HttpHeaders();
        headers.setAccept(List.of(MediaType.APPLICATION_JSON));
        headers.setContentType(MediaType.APPLICATION_FORM_URLENCODED);

        headers.setBasicAuth(clientId,clientSecret);

        MultiValueMap<String, String> request = new LinkedMultiValueMap<String, String>();
        request.add("token", token);

        HttpEntity<?> httpEntity = new HttpEntity<>(request,headers);
        ResponseEntity<ProfileResponse> responseEntity = resilientRestClient.exchange("default", uri, HttpMethod.POST, httpEntity, ProfileResponse.class);
        return responseEntity.getBody();
    }
}
