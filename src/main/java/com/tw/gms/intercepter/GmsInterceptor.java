package com.tw.gms.intercepter;

import org.apache.http.protocol.HttpContext;
import org.apache.logging.log4j.ThreadContext;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.slf4j.MDC;
import org.springframework.lang.Nullable;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.UUID;

@Component
public class GmsInterceptor implements HandlerInterceptor {

    Logger log = LoggerFactory.getLogger(GmsInterceptor.class);

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler)
            throws Exception {
        log.info("log class name {}", log.getClass().getName());
        log.info("log class {}", log.getClass());
        log.info("uuid {}", UUID.randomUUID());
        log.info("time {}", new SimpleDateFormat("yyyy-MM-dd::HH:mm:ss.SSS").format(new Date()));
        log.info("uri {}", request.getRequestURI());
        log.info("method name {}", request.getMethod());
        log.info("thread context {}", ThreadContext.getContext());
        log.info("mdc context map {}", MDC.getCopyOfContextMap());
//        MDC.setContextMap(ThreadContext.getContext());
        return true;
    }

    @Override
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler,
                           @Nullable ModelAndView modelAndView) throws Exception {

        ThreadContext.clearAll();
    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler,
                                @Nullable Exception ex) throws Exception {
    }
}
