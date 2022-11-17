package com.tw.gms.connector;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@ConfigurationProperties(prefix = "rest-template")
@Configuration
@Data
@NoArgsConstructor
@AllArgsConstructor
@ToString
public class RestTemplateProperties {
    private int connectionRequestTimeout = 1000;
    private int connectTimeout = 1000;
    private int readTimeout = 1000;
}
