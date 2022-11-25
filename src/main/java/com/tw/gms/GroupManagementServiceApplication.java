package com.tw.gms;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class GroupManagementServiceApplication {
    private static Logger log = LoggerFactory.getLogger(GroupManagementServiceApplication.class);

    public static void main(String[] args) {
        log.info("arguments specified");
        for (String arg : args) {
            log.info("{}", args);
        }
        SpringApplication.run(GroupManagementServiceApplication.class, args);
    }

}
