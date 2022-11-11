package com.tw.gms.intercepter;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class GmsWebConfig implements WebMvcConfigurer {
    @Autowired
    GmsIntercepter gmsIntercepter;

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(gmsIntercepter).addPathPatterns("/**");
    }
}
