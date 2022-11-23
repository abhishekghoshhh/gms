package com.tw.gms.service.impl;

import com.tw.gms.connector.ResilientRestClient;
import com.tw.gms.connector.RestCallException;
import com.tw.gms.model.ProfileResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.*;
import org.springframework.stereotype.Service;
import org.springframework.web.util.UriComponentsBuilder;

import java.net.URI;
import java.util.List;

@Service
public class ProfileFetcher {

    Logger log = LoggerFactory.getLogger(ProfileFetcher.class);
    @Value("${iam.scim.host}")
    String iamHost;

    @Value("${iam.scim.path:/scim/me}")
    String scimProfileApi;

    @Autowired
    ResilientRestClient resilientRestClient;

    public ProfileResponse fetch(String token) throws RestCallException {
        URI uri = UriComponentsBuilder.fromHttpUrl(iamHost + scimProfileApi).build().toUri();
        HttpHeaders headers = new HttpHeaders();
        headers.setAccept(List.of(MediaType.ALL));
        headers.setBearerAuth(token);
        HttpEntity<?> httpEntity = new HttpEntity<>(headers);
        ResponseEntity<ProfileResponse> responseEntity = resilientRestClient
                .exchange("iam-profile", uri, HttpMethod.GET, httpEntity, ProfileResponse.class);
        log.info("user profile response is {}", responseEntity.getBody());
        return responseEntity.getBody();
    }
}
