package com.tw.gms;

import static io.restassured.RestAssured.given;

public class Main {
    final static String url = "http://localhost:8081/gmsService/search";

    public static void main(String[] args) {
        getResposne1();
    }

    public static void getResponse() {
        given().header("Authorization", "Bearer abcdef")
                .when()
                .get(url)
                .then()
                .log()
                .all();

    }

    public static void getResposne1() {
        given().header("Authorization", "Bearer abcdef").
                queryParam("group", "group1", "group2")
                .when()
                .get(url)
                .then()
                .log()
                .all();

    }
}