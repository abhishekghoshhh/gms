package com.tw.gms.service.impl;

import com.tw.gms.exception.InvalidTokenException;
import com.tw.gms.service.GmsService;
import org.springframework.stereotype.Service;
import org.springframework.util.CollectionUtils;

import java.util.List;

@Service
public class GmsServiceImpl implements GmsService {


    @Override
    public String isAMember(String authorization, String token, List<String> groups) throws InvalidTokenException {
        authorization = getAuthorization(authorization);
        token = getToken(token);


        // case 1
        // there is some groups given as part of request
        // check there is any intersection between groups list and groups in which user is part of

        // case 2
        // there is no group list given in the request
        // return
        return CollectionUtils.isEmpty(groups) ? "demo-group-1\ndemo-group-2\n" : String.join("\n", groups) + "\n";
    }


    private String getAuthorization(String authorization) throws InvalidTokenException {
//        if (null!=authorization && authorization.startsWith("Basic ")) {
//            return authorization.substring(6);
//        } else {
//            throw new InvalidTokenException("invalid authorization header " + authorization);
//        }
        return authorization;
    }

    private String getToken(String token) {
        return token;
    }
}
