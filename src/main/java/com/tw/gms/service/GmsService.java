package com.tw.gms.service;

import com.tw.gms.exception.InvalidTokenException;

import java.util.List;

public interface GmsService {
    String isAMember(String token, List<String> groups) throws InvalidTokenException;
}
