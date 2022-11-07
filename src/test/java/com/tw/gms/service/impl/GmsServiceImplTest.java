package com.tw.gms.service.impl;

import static org.junit.jupiter.api.Assertions.*;

import com.tw.gms.exception.InvalidTokenException;
import org.junit.Before;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.MockitoAnnotations;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.http.ResponseEntity;

import java.util.List;

@ExtendWith(MockitoExtension.class)
public class GmsServiceImplTest {

    @InjectMocks
    GmsServiceImpl gmsServiceImpl;

//    @Before
//    public void setUp(){
//        MockitoAnnotations.initMocks(this);
//    }

    @Test
    public void shouldCheckIfUserBelongsToTheGivenGroups() throws InvalidTokenException {
        List<String> groups = List.of("group1", "group2");
        String expectedGroup = "group1\ngroup2\n";
        String response = gmsServiceImpl.isAMember("Bearer abcdefg", groups);
        assertEquals(expectedGroup, response);
    }
}