package com.tw.gms;

import com.tw.gms.config.ConfigurationProperties;
import io.restassured.RestAssured;
import io.restassured.parsing.Parser;
import io.restassured.response.Response;
import org.testng.Assert;
import org.testng.annotations.BeforeTest;
import org.testng.annotations.Test;

import java.util.List;

import static com.tw.gms.config.TestUtils.toList;
import static com.tw.gms.config.TestUtils.toSet;
import static io.restassured.RestAssured.given;


public class GmsTest {
    private static String url;
    private static ConfigurationProperties configurationProperties;

    static {
        configurationProperties = ConfigurationProperties.getInstance();
    }

    @BeforeTest
    public void configuration() {
        url = configurationProperties.get("gmsUrl");
        RestAssured.useRelaxedHTTPSValidation();
        RestAssured.registerParser("text/plain", Parser.TEXT);
        RestAssured.defaultParser = Parser.TEXT;

    }

    @Test
    public static void whenNoGroupsArePassedInQueryParameter() {
        Response response = given().header("token", configurationProperties.get("demo2Token"))
                .when()
                .get(url);
        response.then().statusCode(200);
        Assert.assertEquals(
                toSet(response.asString().split("\n")),
                toSet("Developer", "Admin", "Viewer"));
    }

    @Test
    public static void whenSingleGroupsArePassedAsQueryParameter() {
        Response response = given().header("token", configurationProperties.get("demo3Token"))
                .queryParam("group", "Developer")
                .when()
                .get(url);
        response.then().statusCode(200);
        List<String> groupsAsList = toList(response.asString().split("\n"));
        Assert.assertEquals(groupsAsList.size(), 1);
        Assert.assertEquals(groupsAsList, List.of("Developer"));
    }

    @Test
    public static void whenValidMultipleGroupsArePassedAsQueryParameter() {
        Response response = given().header("token", configurationProperties.get("demo3Token"))
                .queryParam("group", "Viewer")
                .queryParam("group", "Developer")
                .when()
                .get(url);
        response.then().statusCode(200);
        Assert.assertEquals(
                toSet(response.asString().split("\n")),
                toSet("Developer", "Viewer"));

    }

    @Test
    public static void whenUnknownGroupsArePassedInQueryParameter() {
        Response response = given().header("token", configurationProperties.get("demo4Token"))
                .queryParam("group", "backend")
                .when()
                .get(url);
        response.then().statusCode(200);
        Assert.assertEquals(
                toList(response.asString().split("\n")).size(),
                0);
    }

    @Test
    public static void whenUserIsNotPartOfAnyGroups() {
        Response response = given().header("token", configurationProperties.get("demo1Token"))
                .queryParam("group", "Viewer")
                .queryParam("group", "Developer")
                .when()
                .get(url);
        response.then().statusCode(200);
        Assert.assertEquals(
                toList(response.asString().split("\n")).size(),
                0);
    }

    @Test
    public static void whenIncorrectTokenIsPassed() {
        Response response = given().header("token", configurationProperties.get("dummyToken"))
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

}