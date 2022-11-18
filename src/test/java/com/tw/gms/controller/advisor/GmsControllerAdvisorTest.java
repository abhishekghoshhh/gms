package com.tw.gms.controller.advisor;

import com.tw.gms.connector.RestCallException;
import com.tw.gms.controller.GmsController;
import com.tw.gms.service.GmsService;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.security.test.context.support.WithMockUser;
import org.springframework.test.web.servlet.MockMvc;

import javax.servlet.http.HttpServletRequest;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(GmsController.class)
@WithMockUser
public class GmsControllerAdvisorTest {
    @Autowired
    private MockMvc mockMvc;


    @MockBean
    GmsService gmsService;

    @Test
    public void shouldReturnUnauthorizedStatus() throws Exception {
        RestCallException restCallException = new RestCallException("error", HttpStatus.UNAUTHORIZED, "description");
        when(gmsService.isAMember("token", List.of("group1", "group2"))).thenThrow(restCallException);
        HttpHeaders headers = new HttpHeaders();
        headers.set("token", "token");
        mockMvc.perform(get("/gmsService/search")
                        .param("group", "group1")
                        .param("group", "group2")
                        .headers(headers))
                .andExpect(status().isUnauthorized())
                .andExpect(content().string(""));
    }

    @Test
    public void handleRestCallException() throws Exception {
        GmsControllerAdvisor gmsControllerAdvisor = new GmsControllerAdvisor();
        RestCallException restCallException = new RestCallException("error", HttpStatus.INTERNAL_SERVER_ERROR, "description");
        HttpServletRequest request = Mockito.mock(HttpServletRequest.class);
        assertEquals(restCallException.getHttpStatus(),
                gmsControllerAdvisor.handleRestCallException(restCallException, request).getStatusCode()
        );
    }
}