package com.tw.gms.integration;

import com.tw.gms.integration.utils.ConfigurationProperties;
import io.restassured.RestAssured;
import io.restassured.parsing.Parser;
import io.restassured.response.Response;

import org.testng.Assert;
import org.testng.annotations.BeforeTest;
import org.testng.annotations.Test;

import java.util.List;

import static com.tw.gms.integration.utils.IntegrationTestUtils.*;
import static io.restassured.RestAssured.given;


public class GmsTest {
    private static String url;
    private static final ConfigurationProperties properties;

    static {
        properties = ConfigurationProperties.getInstance();
    }

    @Test(enabled = false)
    public static void whenNoGroupsArePassedInQueryParameter() {
        Response response = given()
                .header("token", properties.get("tokenOfUserWithGroups"))
                .when()
                .get(url);
        response.then().statusCode(200);
        Assert.assertEquals(
                toSetWithSeparator(response.asString(), "\n"),
                properties.getAsSet("groups"));
    }

    @Test(enabled = false)
    public static void whenSingleGroupIsPassedAsQueryParameter() {
        Response response = given()
                .header("token", properties.get("tokenOfUserWithGroups"))
                .queryParam("group", properties.getAsList("groups").get(0))
                .when()
                .get(url);
        response.then().statusCode(200);
        List<String> groupsAsList = toListWithSeparator(response.asString(), "\n");
        Assert.assertEquals(groupsAsList.size(), 1);
        Assert.assertEquals(groupsAsList, properties.getAsList("groups").subList(0, 1));
    }

    @Test(enabled = false)
    public static void whenValidMultipleGroupsArePassedAsQueryParameter() {
        Response response = given()
                .header("token", properties.get("tokenOfUserWithGroups"))
                .queryParam("group", properties.getAsList("groups").get(0))
                .queryParam("group", properties.getAsList("groups").get(1))
                .queryParam("group", "Maintainer")
                .when()
                .get(url);
        response.then().statusCode(200);
        Assert.assertEquals(
                toSetWithSeparator(response.asString(), "\n"),
                toSet(properties.getAsList("groups").get(0), properties.getAsList("groups").get(1)));

    }

    @Test(enabled = false)
    public static void whenUnknownGroupsArePassedInQueryParameter() {
        Response response = given()
                .header("token", properties.get("tokenOfUserWithGroups"))
                .queryParam("group", "Unknown")
                .when()
                .get(url);
        response.then().statusCode(200);
        Assert.assertEquals(
                toListWithSeparator(response.asString(), "\n").size(),
                0);
    }

    @Test(enabled = false)
    public static void whenUserIsNotPartOfAnyGroups() {
        Response response = given()
                .header("token", properties.get("tokenOfUserWithNoGroups"))
                .queryParam("group", "Viewer")
                .queryParam("group", "Developer")
                .when()
                .get(url);
        response.then().statusCode(200);
        Assert.assertEquals(
                toListWithSeparator(response.asString(), "\n").size(),
                0);
    }

    @Test(enabled = false)
    public static void whenIncorrectTokenIsPassed() {
        Response response = given().
                header("token", "incorrectToken")
                .when()
                .get(url);
        response.then().statusCode(401);
    }

    @Test
    public static void whenNoTokenIsPassed() {
        Response response = given()
                .when()
                .get(url);
        response.then().statusCode(400);
    }

    @BeforeTest
    public void configuration() {
        url = properties.get("gmsUrl");
        RestAssured.useRelaxedHTTPSValidation();
        RestAssured.registerParser("text/plain", Parser.TEXT);
        RestAssured.defaultParser = Parser.TEXT;
    }

}