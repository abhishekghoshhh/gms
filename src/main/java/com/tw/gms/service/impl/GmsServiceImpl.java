package com.tw.gms.service.impl;

import com.tw.gms.connector.RestCallException;
import com.tw.gms.model.Group;
import com.tw.gms.model.ProfileResponse;
import com.tw.gms.service.GmsService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
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
    Logger log = LoggerFactory.getLogger(GmsServiceImpl.class);

    @Override
    public String isAMember(String token, List<String> groups) throws RestCallException {
        log.info("Incoming group list is {}", groups);

        ProfileResponse profileResponse = profileFetcher.fetch(token);
        Set<String> profileGroups = groupNames(profileResponse);
        log.info("User group list {}", profileGroups);

        if (CollectionUtils.isEmpty(profileGroups)) {
            log.debug("User belongs to no group");
            return EMPTY_STRING;
        } else if (CollectionUtils.isEmpty(groups)) {
            log.debug("Incoming groups list is empty");
            return String.join("\n", profileGroups)
                    .concat("\n");
        } else {
            List<String> filteredGroups = groups.stream()
                    .filter(profileGroups::contains)
                    .distinct()
                    .collect(Collectors.toList());
            return filteredGroups.isEmpty() ? EMPTY_STRING : String.join("\n", filteredGroups).concat("\n");
        }
    }

    private Set<String> groupNames(ProfileResponse profileResponse) {
        if (null != profileResponse && !CollectionUtils.isEmpty(profileResponse.getGroups()))
            return profileResponse.getGroups().stream()
                    .map(Group::getDisplay)
                    .collect(Collectors.toSet());
        return new HashSet<>();
    }

}
