package com.tw.gms.controller.advisor;

import com.tw.gms.exception.InvalidTokenException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import javax.servlet.http.HttpServletRequest;

@ControllerAdvice
public class GmsControllerAdvisor extends ResponseEntityExceptionHandler {

    @ExceptionHandler(InvalidTokenException.class)
    public ResponseEntity handleCityNotFoundException(
            InvalidTokenException ex, HttpServletRequest request) {
        return ResponseEntity.badRequest().build();
    }
}
