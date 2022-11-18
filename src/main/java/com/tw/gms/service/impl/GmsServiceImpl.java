package com.tw.gms.service.impl;

import com.tw.gms.connector.RestCallException;
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
        log.debug("Incoming groups are {}",groups);
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
