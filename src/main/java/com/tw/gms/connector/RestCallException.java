package com.tw.gms.connector;

import org.springframework.http.HttpStatus;

public class RestCallException extends Exception {
    private String message;
    private HttpStatus httpStatus = HttpStatus.INTERNAL_SERVER_ERROR;
    private String description;

    public RestCallException(String message) {
        super(message);
        this.message = message;
    }

    public RestCallException(Throwable exception) {
        super(exception);
        this.message = exception.getMessage();
    }

    public RestCallException(String message, HttpStatus httpStatus, String description) {
        super(message);
        this.message = message;
        this.httpStatus = httpStatus;
        this.description = description;
    }
    public RestCallException(String message, String description) {
        super(message);
        this.message = message;
        this.description = description;
    }

    @Override
    public String getMessage() {
        return message;
    }

    public HttpStatus getHttpStatus() {
        return httpStatus;
    }

    public String getDescription() {
        return description;
    }
}
