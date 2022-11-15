package com.tw.gms.service.impl;

import com.tw.gms.connector.ResilientRestClient;
import com.tw.gms.model.ProfileResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.*;
import org.springframework.stereotype.Service;
import org.springframework.web.util.UriComponentsBuilder;

import java.net.URI;
import java.util.List;

@Service
public class ProfileFetcher {

    @Value("${iam.scim.host}")
    private String iamHost;

    @Value("${iam.scim.path:/scim/me}")
    private String scimProfileApi;

    @Autowired
    ResilientRestClient resilientRestClient;

    public ProfileResponse fetch(String token) {
        URI uri = UriComponentsBuilder.fromHttpUrl(iamHost + scimProfileApi).build().toUri();
        HttpHeaders headers = new HttpHeaders();
        headers.setAccept(List.of(MediaType.ALL));
        headers.setBearerAuth(token);
        HttpEntity<?> httpEntity = new HttpEntity<>(headers);
        ResponseEntity<ProfileResponse> responseEntity = resilientRestClient
                .exchange("default", uri, HttpMethod.GET, httpEntity, ProfileResponse.class);
        return responseEntity.getBody();
    }
}
