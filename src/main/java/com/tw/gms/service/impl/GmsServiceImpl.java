package com.tw.gms.service.impl;

import com.tw.gms.exception.InvalidTokenException;
import com.tw.gms.model.ProfileResponse;
import com.tw.gms.service.GmsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.util.CollectionUtils;

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
        ProfileResponse profileResponse = profileFetcher.fetch(token);
        if (CollectionUtils.isEmpty(profileResponse.getGroups())) {
            return EMPTY_STRING;
        } else if (CollectionUtils.isEmpty(groups)) {
            return String.join("\n", profileResponse.getGroups()) + "\n";
        } else {
            Set<String> profileGroups = new HashSet<>(profileResponse.getGroups());
            StringBuilder groupsResponseString = new StringBuilder();
            for (String group : groups) {
                if (profileGroups.contains(group)) {
                    groupsResponseString.append(group).append("\n");
                }
            }
            return groupsResponseString.toString();
        }
    }

}
