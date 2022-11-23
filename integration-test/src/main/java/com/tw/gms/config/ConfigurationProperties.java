package com.tw.gms.config;

import java.io.FileInputStream;
import java.io.IOException;
import java.nio.file.Paths;
import java.util.*;
import java.util.stream.Collectors;

public class ConfigurationProperties {
    private static ConfigurationProperties configurationProperties = null;
    private final Properties properties = new Properties();

    private ConfigurationProperties() {
        String relativeLocation = "/src/test/resources/application.properties";
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

    public String get(String key, String defaultValue) {
        return properties.getProperty(key, defaultValue);
    }

}
