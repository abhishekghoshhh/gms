package com.tw.gms.connector;

public class RestCallException extends Exception {
    private String message;

    public RestCallException(String message) {
        super(message);
        this.message = message;
    }
    public RestCallException(Throwable exception) {
        super(exception);
        this.message = exception.getMessage();
    }
}
