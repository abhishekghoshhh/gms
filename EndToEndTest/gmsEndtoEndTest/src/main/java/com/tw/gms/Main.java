package com.tw.gms;

import io.restassured.RestAssured;
import io.restassured.parsing.Parser;
import io.restassured.response.Response;
import org.apache.commons.lang3.builder.ToStringExclude;
import org.testng.Assert;
import org.testng.annotations.BeforeTest;
import org.testng.annotations.Test;

import javax.net.ssl.SSLHandshakeException;
import java.util.Locale;

import static io.restassured.RestAssured.given;


public class Main  {
    final static String url = "https://10.131.101.52:473/gmsService/search";

    @BeforeTest
    public static void main()  {
        RestAssured.useRelaxedHTTPSValidation();
        RestAssured.registerParser("text/plain", Parser.TEXT);
        RestAssured.defaultParser = Parser.TEXT;

    }




//When valid multiple groups are passed as query parameter
    @Test
    public static void getResponse2() {
        Response r=given().header("token", "eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI1MGY5M2RiNS0yZDg2LTQwN2QtYWExMS0yMTE0Mzg5ZWY2MTUiLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MCIsImV4cCI6MTY2OTAyNzgzNiwiaWF0IjoxNjY5MDI0MjM2LCJqdGkiOiJmMDhlMTUxNy1jOTVjLTRkZGMtODYxMy03ZmQwZjVhYTQxMGEiLCJjbGllbnRfaWQiOiJjbGllbnQifQ.JsF5A0Prkt8qTXKWNk5y2pOPOVxBOvpgr6ZEbQ3M7B1yCS5cNWUJzmzuR0iou_NtTiqfFrQ24LEcDq-EVS1aetsOIlkg_np9yWW3U9SHek9Rzq2sd6ea5SbfLGY9JuSGEmS7WVS4lYIK3VhJBoaW_FG8UQH2OwzN1SPCM9GePuLzu3TnklcP-_-ADYQNDOWVqvLbLxRbh2S3PFV8kf1W5XobPmCLOqxe-WvTZi7vrAHua0nmUvhmpQLuL8N4Uizr8JR3H-n2F-bPR99A27z2bZSgIQNWXOzhvsCvqE59HgJVi3zK5Nd4z6ODvMkMMKuDqwII7UdiZsdcMoY-st2-MQ")
                .queryParam("group","Viewer")
                .queryParam("group","Developer")
                .when()
                .get(url);
        String t =r.asString();
        r.then().statusCode(200);
        System.out.print(t+"hari"+"gcd");
        String[] group1=t.split("\n");
        Assert.assertEquals(group1[0],"Viewer");
        Assert.assertEquals(group1[1],"Developer");

    }

    //When single groups are passed as query parameter
    @Test
    public static void getResponse3() {
        Response r=  given().header("token", "eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI1MGY5M2RiNS0yZDg2LTQwN2QtYWExMS0yMTE0Mzg5ZWY2MTUiLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MCIsImV4cCI6MTY2ODc3Mjg2MSwiaWF0IjoxNjY4NzY5MjYxLCJqdGkiOiI2YzFjYmI3NC1mNmVlLTQ0YWMtODUwNC04N2U0OTI3NWZmNjMiLCJjbGllbnRfaWQiOiJjbGllbnQifQ.ZRug-Mz41uiJXqppbC3cLuACL1qkqAB13uA_rKoN_wLjCAdLj8UlZy09H1RWdddz6-fx0DN7jxo4an6l-DTlYdUzsNcWmUdFb6cdaTp7WOtlYDu8-f-ua0KsQV93UYzksiQXkbNBfc1JN7PgzM1Xc7ss4CZGsMpHKfHz3j7yX2CGZwjNittZlSU_Pf4IqDF0rT7TPuh6xoJcopg4gcZDh-Lb_iIJzwDUB7RRqRV4S0evYbV83q7xqIDpysIpUidk5LhwIOAFZiq14s8WBdyX4pbs9SLJyKHSv_2RVwfA2n_f_UXRXIg0s-LLxmck04n9ZPwVne66Lhyc8TrUBDEXsQ")
                .queryParam("group","Developer")
                .when()
                .get(url);
        String t =r.asString();
        String[] group1=t.split("\n");
        Assert.assertEquals(group1[0],"Developer");

    }

    //When incorrect tokens are passed

    @Test
    public static void getResponse4() {
        Response r= given().header("token", "incorrect token")
                .queryParam("group","Developer")
                .when()
                .get(url);
        r.then().statusCode(401);
        String t =r.asString();
        String[] group1=t.split("\n");
        Assert.assertEquals(group1[0],"");


    }

    //when no groups are passed in query parameter
    @Test
    public static void getResponse5() {
        Response r=given().header("token", "eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI1MGY5M2RiNS0yZDg2LTQwN2QtYWExMS0yMTE0Mzg5ZWY2MTUiLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MCIsImV4cCI6MTY2OTAyNzgzNiwiaWF0IjoxNjY5MDI0MjM2LCJqdGkiOiJmMDhlMTUxNy1jOTVjLTRkZGMtODYxMy03ZmQwZjVhYTQxMGEiLCJjbGllbnRfaWQiOiJjbGllbnQifQ.JsF5A0Prkt8qTXKWNk5y2pOPOVxBOvpgr6ZEbQ3M7B1yCS5cNWUJzmzuR0iou_NtTiqfFrQ24LEcDq-EVS1aetsOIlkg_np9yWW3U9SHek9Rzq2sd6ea5SbfLGY9JuSGEmS7WVS4lYIK3VhJBoaW_FG8UQH2OwzN1SPCM9GePuLzu3TnklcP-_-ADYQNDOWVqvLbLxRbh2S3PFV8kf1W5XobPmCLOqxe-WvTZi7vrAHua0nmUvhmpQLuL8N4Uizr8JR3H-n2F-bPR99A27z2bZSgIQNWXOzhvsCvqE59HgJVi3zK5Nd4z6ODvMkMMKuDqwII7UdiZsdcMoY-st2-MQ")
                .queryParam("group","Viewer")
                .queryParam("group","Developer")
                .when()
                .get(url);
        String t =r.asString();
        String[] group1=t.split("\n");
        r.then().statusCode(200);
        Assert.assertEquals(group1[0],"Viewer");
        Assert.assertEquals(group1[1],"Developer");
    }

    //when invalid groups are passed in query parameter
    @Test
    public static void getResponse6() {
        Response r= given().header("token", "eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI1MGY5M2RiNS0yZDg2LTQwN2QtYWExMS0yMTE0Mzg5ZWY2MTUiLCJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3Q6ODA4MCIsImV4cCI6MTY2ODc3Mjg2MSwiaWF0IjoxNjY4NzY5MjYxLCJqdGkiOiI2YzFjYmI3NC1mNmVlLTQ0YWMtODUwNC04N2U0OTI3NWZmNjMiLCJjbGllbnRfaWQiOiJjbGllbnQifQ.ZRug-Mz41uiJXqppbC3cLuACL1qkqAB13uA_rKoN_wLjCAdLj8UlZy09H1RWdddz6-fx0DN7jxo4an6l-DTlYdUzsNcWmUdFb6cdaTp7WOtlYDu8-f-ua0KsQV93UYzksiQXkbNBfc1JN7PgzM1Xc7ss4CZGsMpHKfHz3j7yX2CGZwjNittZlSU_Pf4IqDF0rT7TPuh6xoJcopg4gcZDh-Lb_iIJzwDUB7RRqRV4S0evYbV83q7xqIDpysIpUidk5LhwIOAFZiq14s8WBdyX4pbs9SLJyKHSv_2RVwfA2n_f_UXRXIg0s-LLxmck04n9ZPwVne66Lhyc8TrUBDEXsQ")
                .queryParam("group","backend")
                .when()
                .get(url);
        r.then().statusCode(200);
        String t =r.asString();
        String[] group1=t.split("\n");
        Assert.assertEquals(group1[0],"");
    }

}