package com.tw.gms.service.impl;

import com.tw.gms.exception.InvalidTokenException;
import com.tw.gms.model.ProfileResponse;
import com.tw.gms.service.GmsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.util.CollectionUtils;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

@Service
public class GmsServiceImpl implements GmsService {
    public static final String EMPTY_STRING = "";
    @Autowired
    ProfileFetcher profileFetcher;

    @Override
    public String isAMember(String token, List<String> groups) throws InvalidTokenException {

        //call the introspection api
        ProfileResponse profileResponse = profileFetcher.fetch(token);

        if (CollectionUtils.isEmpty(profileResponse.getGroups())) {
            // case 1
            // user is not part of any group
            return EMPTY_STRING;
        } else if (CollectionUtils.isEmpty(groups)) {
            // case 2
            // there is no group list given in the request
            // return
            return String.join("\n", profileResponse.getGroups()) + "\n";
        } else {
            // case 2
            // there is some groups given as part of request
            // check there is any intersection between groups list and groups in which user is part of
            Set<String> inputGroups = new HashSet<>(groups);
            Set<String> profileGroups = new HashSet<>(profileResponse.getGroups());
            StringBuilder groupsResponseString = new StringBuilder();
            for(String group:inputGroups){
                if(profileGroups.contains(group)){
                    groupsResponseString.append(group).append("\n");
                }
            }
            return groupsResponseString.toString();
        }
    }

}
