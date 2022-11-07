package com.tw.gms.service.impl;

import com.tw.gms.exception.InvalidTokenException;
import com.tw.gms.service.GmsService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class GmsServiceImpl implements GmsService {


    @Override
    public String isAMember(String authorization, List<String> groups) throws InvalidTokenException {
        String token = extractToken(authorization);
        // case 1
        // there is some groups given as part of request
        // check there is any intersection between groups list and groups in which user is part of

        // case 2
        // there is no group list given in the request
        // return
        return String.join("\n", groups) + "\n";

    }

    private String extractToken(String authorization) throws InvalidTokenException {
        if (authorization.contains("Bearer")) {
            return authorization.substring(7);
        } else {
            throw new InvalidTokenException("invalid authorization header " + authorization);
        }
    }
}
