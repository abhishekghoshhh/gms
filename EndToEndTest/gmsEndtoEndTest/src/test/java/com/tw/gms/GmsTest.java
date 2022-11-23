package com.tw.gms;

import io.restassured.RestAssured;
import io.restassured.parsing.Parser;
import io.restassured.response.Response;
import org.testng.Assert;
import org.testng.annotations.BeforeTest;
import org.testng.annotations.Test;
import static io.restassured.RestAssured.given;


public class GmsTest {
    final static String url = "https://10.131.101.52:443/gmsService/search";

    @BeforeTest
    public static void configuration() {
        RestAssured.useRelaxedHTTPSValidation();
        RestAssured.registerParser("text/plain", Parser.TEXT);
        RestAssured.defaultParser = Parser.TEXT;

    }
    @Test
    public static void whenValidMultipleGroupsArePassedAsQueryParameter() {
        Response r = given().header("token", "eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI1MGY5M2RiNS0yZDg2LTQwN2QtYWExMS0yMTE0Mzg5ZWY2MTUiLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MCIsImV4cCI6MTY2OTAyNzgzNiwiaWF0IjoxNjY5MDI0MjM2LCJqdGkiOiJmMDhlMTUxNy1jOTVjLTRkZGMtODYxMy03ZmQwZjVhYTQxMGEiLCJjbGllbnRfaWQiOiJjbGllbnQifQ.JsF5A0Prkt8qTXKWNk5y2pOPOVxBOvpgr6ZEbQ3M7B1yCS5cNWUJzmzuR0iou_NtTiqfFrQ24LEcDq-EVS1aetsOIlkg_np9yWW3U9SHek9Rzq2sd6ea5SbfLGY9JuSGEmS7WVS4lYIK3VhJBoaW_FG8UQH2OwzN1SPCM9GePuLzu3TnklcP-_-ADYQNDOWVqvLbLxRbh2S3PFV8kf1W5XobPmCLOqxe-WvTZi7vrAHua0nmUvhmpQLuL8N4Uizr8JR3H-n2F-bPR99A27z2bZSgIQNWXOzhvsCvqE59HgJVi3zK5Nd4z6ODvMkMMKuDqwII7UdiZsdcMoY-st2-MQ")
                .queryParam("group", "Viewer")
                .queryParam("group", "Developer")
                .when()
                .get(url);
        String t = r.asString();
        r.then().statusCode(200);
        String[] group1 = t.split("\n");
        Assert.assertEquals(group1[0], "Viewer");
        Assert.assertEquals(group1[1], "Developer");

    }
    @Test
    public static void whenSingleGroupsArePassedAsQueryParameter() {
        Response response = given().header("token", "eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI1MGY5M2RiNS0yZDg2LTQwN2QtYWExMS0yMTE0Mzg5ZWY2MTUiLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MCIsImV4cCI6MTY2OTAyNzgzNiwiaWF0IjoxNjY5MDI0MjM2LCJqdGkiOiJmMDhlMTUxNy1jOTVjLTRkZGMtODYxMy03ZmQwZjVhYTQxMGEiLCJjbGllbnRfaWQiOiJjbGllbnQifQ.JsF5A0Prkt8qTXKWNk5y2pOPOVxBOvpgr6ZEbQ3M7B1yCS5cNWUJzmzuR0iou_NtTiqfFrQ24LEcDq-EVS1aetsOIlkg_np9yWW3U9SHek9Rzq2sd6ea5SbfLGY9JuSGEmS7WVS4lYIK3VhJBoaW_FG8UQH2OwzN1SPCM9GePuLzu3TnklcP-_-ADYQNDOWVqvLbLxRbh2S3PFV8kf1W5XobPmCLOqxe-WvTZi7vrAHua0nmUvhmpQLuL8N4Uizr8JR3H-n2F-bPR99A27z2bZSgIQNWXOzhvsCvqE59HgJVi3zK5Nd4z6ODvMkMMKuDqwII7UdiZsdcMoY-st2-MQ")
                .queryParam("group", "Developer")
                .when()
                .get(url);
        response.then().statusCode(200);
        String t = response.asString();
        String[] group1 = t.split("\n");
        Assert.assertEquals(group1[0], "Developer");

    }
    @Test
    public static void whenIncorrectTokensArePassed() {
        Response response = given().header("token", "incorrect token")
                .queryParam("group", "Developer")
                .when()
                .get(url);
        response.then().statusCode(401);
        String t = response.asString();
        String[] group1 = t.split("\n");
        Assert.assertEquals(group1[0], "");


    }
    @Test
    public static void whenNoGroupsArePassedInQueryParameter() {
        Response response = given().header("token", "eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI1MGY5M2RiNS0yZDg2LTQwN2QtYWExMS0yMTE0Mzg5ZWY2MTUiLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MCIsImV4cCI6MTY2OTAyNzgzNiwiaWF0IjoxNjY5MDI0MjM2LCJqdGkiOiJmMDhlMTUxNy1jOTVjLTRkZGMtODYxMy03ZmQwZjVhYTQxMGEiLCJjbGllbnRfaWQiOiJjbGllbnQifQ.JsF5A0Prkt8qTXKWNk5y2pOPOVxBOvpgr6ZEbQ3M7B1yCS5cNWUJzmzuR0iou_NtTiqfFrQ24LEcDq-EVS1aetsOIlkg_np9yWW3U9SHek9Rzq2sd6ea5SbfLGY9JuSGEmS7WVS4lYIK3VhJBoaW_FG8UQH2OwzN1SPCM9GePuLzu3TnklcP-_-ADYQNDOWVqvLbLxRbh2S3PFV8kf1W5XobPmCLOqxe-WvTZi7vrAHua0nmUvhmpQLuL8N4Uizr8JR3H-n2F-bPR99A27z2bZSgIQNWXOzhvsCvqE59HgJVi3zK5Nd4z6ODvMkMMKuDqwII7UdiZsdcMoY-st2-MQ")
                .queryParam("group", "Viewer")
                .queryParam("group", "Developer")
                .when()
                .get(url);
        String t = response.asString();
        String[] group1 = t.split("\n");
        response.then().statusCode(200);
        Assert.assertEquals(group1[0], "Viewer");
        Assert.assertEquals(group1[1], "Developer");
    }
    @Test
    public static void whenInvalidGroupsArePassedInQueryParameter() {
        Response response = given().header("token", "eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI1MGY5M2RiNS0yZDg2LTQwN2QtYWExMS0yMTE0Mzg5ZWY2MTUiLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MCIsImV4cCI6MTY2OTAyNzgzNiwiaWF0IjoxNjY5MDI0MjM2LCJqdGkiOiJmMDhlMTUxNy1jOTVjLTRkZGMtODYxMy03ZmQwZjVhYTQxMGEiLCJjbGllbnRfaWQiOiJjbGllbnQifQ.JsF5A0Prkt8qTXKWNk5y2pOPOVxBOvpgr6ZEbQ3M7B1yCS5cNWUJzmzuR0iou_NtTiqfFrQ24LEcDq-EVS1aetsOIlkg_np9yWW3U9SHek9Rzq2sd6ea5SbfLGY9JuSGEmS7WVS4lYIK3VhJBoaW_FG8UQH2OwzN1SPCM9GePuLzu3TnklcP-_-ADYQNDOWVqvLbLxRbh2S3PFV8kf1W5XobPmCLOqxe-WvTZi7vrAHua0nmUvhmpQLuL8N4Uizr8JR3H-n2F-bPR99A27z2bZSgIQNWXOzhvsCvqE59HgJVi3zK5Nd4z6ODvMkMMKuDqwII7UdiZsdcMoY-st2-MQ")
                .queryParam("group", "backend")
                .when()
                .get(url);
        response.then().statusCode(200);
        String t = response.asString();
        String[] group1 = t.split("\n");
        Assert.assertEquals(group1[0], "");
    }

}