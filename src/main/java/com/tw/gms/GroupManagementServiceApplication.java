package com.tw.gms;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class GroupManagementServiceApplication {
    private static Logger log = LoggerFactory.getLogger(GroupManagementServiceApplication.class);

    public static void main(String[] args) {
        log.debug("arguments specified");
        for (String arg : args) {
            log.debug("{}", arg);
        }
        SpringApplication.run(GroupManagementServiceApplication.class, args);
    }
}
