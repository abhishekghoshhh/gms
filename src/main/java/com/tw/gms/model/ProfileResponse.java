package com.tw.gms.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.util.CollectionUtils;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@JsonIgnoreProperties(ignoreUnknown = true)
@Data
@NoArgsConstructor
@AllArgsConstructor
public class ProfileResponse {
    private List<Group> groups;

    public List<String> groupNamesList() {
        if (CollectionUtils.isEmpty(groups)) return new ArrayList<>();
        return groups.stream()
                .map(Group::getDisplay)
                .collect(Collectors.toList());
    }
}
