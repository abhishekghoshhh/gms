package com.tw.gms.controller.advisor;

import com.tw.gms.connector.RestCallException;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import javax.servlet.http.HttpServletRequest;

@ControllerAdvice
public class GmsControllerAdvisor extends ResponseEntityExceptionHandler {
    @ExceptionHandler(RestCallException.class)
    public ResponseEntity<ErrorResponse> handleRestCallException(
            RestCallException ex, HttpServletRequest request) {
        return ResponseEntity
                .status(ex.getHttpStatus())
                .build();
//                .body(new ErrorResponse(
//                        ex.getMessage(),
//                        ex.getHttpStatus().value(),
//                        ex.getDescription())
//                );
    }
}
