package com.tw.gms.config;

import java.io.FileInputStream;
import java.io.IOException;
import java.util.Properties;

public class ConfigurationProperties {
    private static ConfigurationProperties configurationProperties = null;
    private Properties properties = new Properties();

    private ConfigurationProperties() {
        loadApplicationProperties();
    }

    private void loadApplicationProperties() {
        String relativeLocation = "/src/main/resources/application.properties";
        try {
            properties.load(new FileInputStream(relativeLocation));
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static ConfigurationProperties getInstance() {
        if (null == configurationProperties)
            configurationProperties = new ConfigurationProperties();
        return configurationProperties;
    }
    public String get(String key){
        return null;
    }
}
