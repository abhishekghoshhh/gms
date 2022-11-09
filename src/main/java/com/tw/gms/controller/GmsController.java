package com.tw.gms.controller;

import com.tw.gms.exception.InvalidTokenException;
import com.tw.gms.service.GmsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class GmsController {
    @Autowired
    GmsService gmsService;

    @GetMapping(value = "/gmsService/search", produces = MediaType.TEXT_PLAIN_VALUE)
    public ResponseEntity<String> isAMember
            (@RequestHeader(value = "Authorization", required = false) String authorization,
             @RequestHeader(value = "token") String token,
             @RequestParam(name = "group", required = false) List<String> groups) throws InvalidTokenException {
        return ResponseEntity.ok(gmsService.isAMember(authorization,token, groups));
    }
}