package com.tw.gms.service.impl;

import com.tw.gms.connector.ResilientRestClient;
import com.tw.gms.connector.RestCallException;
import com.tw.gms.model.ProfileResponse;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;

import java.net.URI;

import static com.tw.gms.utils.TestUtils.setFieldByReflection;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.ArgumentMatchers.*;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class ProfileFetcherTest {

    @InjectMocks
    ProfileFetcher profileFetcher;

    @Mock
    ResilientRestClient resilientRestClient;

    @Test
    public void fetch() throws NoSuchFieldException, RestCallException {
        ResponseEntity<ProfileResponse> response = ResponseEntity.ok(new ProfileResponse(null));
        setFieldByReflection(ProfileFetcher.class, profileFetcher, "iamHost", "http://127.0.0.1:8080");
        setFieldByReflection(ProfileFetcher.class, profileFetcher, "scimProfileApi", "/scim/me");
        when(resilientRestClient.exchange(anyString(), any(URI.class), any(HttpMethod.class), any(HttpEntity.class), eq(ProfileResponse.class))).thenReturn(response);
        assertEquals(response.getBody(), profileFetcher.fetch("token"));
    }

}