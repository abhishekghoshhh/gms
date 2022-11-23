package com.tw.gms.config;

import java.util.*;
import java.util.stream.Collectors;

public class TestUtils {
    public static Set<String> toSet(String... values) {
        if (values.length == 0) return new HashSet<>();
        return Arrays.stream(values).collect(Collectors.toSet());
    }

    public static List<String> toList(String... values) {
        if (values.length == 0) return new ArrayList<>();
        return Arrays.stream(values).filter(str -> !str.isBlank()).collect(Collectors.toList());
    }
}
