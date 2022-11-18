package com.tw.gms.controller;

import com.tw.gms.service.GmsService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.HttpHeaders;
import org.springframework.security.test.context.support.WithMockUser;
import org.springframework.test.web.servlet.MockMvc;

import java.util.List;

import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(GmsController.class)
@WithMockUser
public class GmsControllerTest {
    @Autowired
    private MockMvc mockMvc;


    @MockBean
    GmsService gmsService;

    @Test
    void isAMember() throws Exception {
        when(gmsService.isAMember( "token", List.of("group1", "group2"))).thenReturn("group1\ngroup2\n");
        HttpHeaders headers = new HttpHeaders();
        headers.set("token", "token");
        mockMvc.perform(get("/gmsService/search")
                        .param("group", "group1")
                        .param("group", "group2")
                        .headers(headers))
                .andExpect(status().isOk())
                .andExpect(content().string("group1\ngroup2\n"));
    }
}