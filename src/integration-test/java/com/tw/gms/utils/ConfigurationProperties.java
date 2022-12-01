package com.tw.gms.utils;

import java.io.FileInputStream;
import java.io.IOException;
import java.nio.file.Paths;
import java.util.*;
import java.util.stream.Collectors;

public class ConfigurationProperties {
    private static ConfigurationProperties configurationProperties = null;
    private final Properties properties = new Properties();

    private ConfigurationProperties() {
        String relativeLocation = "src/integration-test/resources/application.properties";
        loadApplicationProperties(relativeLocation);
    }

    public static ConfigurationProperties getInstance() {
        if (null == configurationProperties)
            configurationProperties = new ConfigurationProperties();
        return configurationProperties;
    }

    private void loadApplicationProperties(String relativeLocation) {
        try {
            properties.load(
                    new FileInputStream(
                            Paths.get(System.getProperty("user.dir"),
                                    relativeLocation).toString()
                    ));
        } catch (IOException e) {
            System.out.println(e.getMessage());
            throw new RuntimeException(e);
        }
    }


    public String get(String key) {
        return properties.getProperty(key);
    }

    public Set<String> getAsSet(String key) {
        String value = get(key);
        if (null == value || "".equalsIgnoreCase(value.trim())) return new HashSet<>();
        String[] values = value.split(",");
        return Arrays
                .stream(values).parallel()
                .filter(str -> null != str && !"".equalsIgnoreCase(str.trim()))
                .map(String::trim)
                .collect(Collectors.toSet());
    }

    public List<String> getAsList(String key) {
        String value = get(key);
        if (null == value || "".equalsIgnoreCase(value.trim())) return new ArrayList<>();
        String[] values = value.split(",");
        return Arrays
                .stream(values).parallel()
                .filter(str -> null != str && !"".equalsIgnoreCase(str.trim()))
                .map(String::trim)
                .collect(Collectors.toList());
    }

}
