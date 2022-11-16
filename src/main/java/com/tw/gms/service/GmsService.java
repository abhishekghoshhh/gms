package com.tw.gms.service;

import com.tw.gms.connector.RestCallException;

import java.util.List;

public interface GmsService {
    String isAMember(String token, List<String> groups) throws RestCallException;
}
