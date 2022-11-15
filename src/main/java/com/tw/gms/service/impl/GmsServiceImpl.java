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
import java.util.stream.Collectors;

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
            return String.join("\n", profileResponse.groupNamesList())
                    .concat("\n");
        } else {
            Set<String> profileGroups = new HashSet<>(profileResponse.groupNamesList());
            List<String> filteredGroups = groups.stream()
                    .filter(profileGroups::contains)
                    .collect(Collectors.toList());
            return filteredGroups.isEmpty() ? EMPTY_STRING : String.join("\n", filteredGroups).concat("\n");
        }
    }

}
