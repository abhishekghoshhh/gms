package com.tw.gms.service.impl;

import com.tw.gms.exception.InvalidTokenException;
import com.tw.gms.model.ProfileResponse;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.ArrayList;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

@ExtendWith(MockitoExtension.class)
public class GmsServiceImplTest {

    @InjectMocks
    GmsServiceImpl gmsServiceImpl;

    @Mock
    ProfileFetcher profileFetcher;


    @Test
    public void shouldCheckIfUserBelongsToNoGroup() throws InvalidTokenException {
        List<String> groups = List.of("group1", "group2");
        ProfileResponse profileResponse = new ProfileResponse(true, null);
        String token = "token";
        String expectedGroup = "";
        Mockito.when(profileFetcher.fetch(token)).thenReturn(profileResponse);
        assertEquals(expectedGroup, gmsServiceImpl.isAMember(token, groups));
    }

    @Test
    public void shouldCheckIfNoGroupsGivenInRequest() throws InvalidTokenException {
        List<String> groups = new ArrayList<>();
        ProfileResponse profileResponse = new ProfileResponse(true, List.of("group1", "group2"));
        String token = "token";
        Mockito.when(profileFetcher.fetch(token)).thenReturn(profileResponse);
        String expectedGroup = "group1\ngroup2\n";
        assertEquals(expectedGroup, gmsServiceImpl.isAMember(token, groups));
    }

    @Test
    public void shouldCheckIfUserBelongsToTheGivenGroups() throws InvalidTokenException {
        List<String> groups = List.of("group1", "group3","group5");
        ProfileResponse profileResponse = new ProfileResponse(true, List.of("group1", "group2","group3","group4"));
        String token = "token";
        Mockito.when(profileFetcher.fetch(token)).thenReturn(profileResponse);
        String expectedGroup = "group1\ngroup3\n";
        assertEquals(expectedGroup, gmsServiceImpl.isAMember(token, groups));
    }
}