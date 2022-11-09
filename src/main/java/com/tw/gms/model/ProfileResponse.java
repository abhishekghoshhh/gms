package com.tw.gms.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import lombok.Data;

import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
@Data
public class ProfileResponse {
    private boolean active;
    private List<String> groups;
}
