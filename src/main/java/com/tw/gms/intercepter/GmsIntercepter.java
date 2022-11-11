package com.tw.gms.intercepter;

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
import java.util.UUID;

@Component
public class GmsIntercepter implements HandlerInterceptor {

    Logger log = LoggerFactory.getLogger(GmsIntercepter.class);

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler)
            throws Exception {
        ThreadContext.put("uuid", UUID.randomUUID().toString()); // Add the fishtag;
        ThreadContext.put("ipAddress", request.getRemoteAddr());
        ThreadContext.put("loginId", (String) request.getSession().getAttribute("loginId"));
        ThreadContext.put("hostName", request.getServerName());
        MDC.setContextMap(ThreadContext.getContext());
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
