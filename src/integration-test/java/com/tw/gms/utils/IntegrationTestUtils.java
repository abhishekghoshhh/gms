package com.tw.gms.utils;

import java.util.*;
import java.util.stream.Collectors;

public class IntegrationTestUtils {
    public static Set<String> toSet(String... values) {
        if (values.length == 0) return new HashSet<>();
        return Arrays.stream(values).collect(Collectors.toSet());
    }

    public static List<String> toList(String... values) {
        if (values.length == 0) return new ArrayList<>();
        return Arrays.stream(values).filter(str -> !str.isBlank()).collect(Collectors.toList());
    }

    public static Set<String> toSetWithSeparator(String value, String separator) {
        return toSet(value.split(separator));
    }

    public static List<String> toListWithSeparator(String value, String separator) {
        return toList(value.split(separator));
    }
}
