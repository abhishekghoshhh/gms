package com.tw.gms.utils;

import org.springframework.util.ReflectionUtils;

import java.lang.reflect.Field;

public class TestUtils {
    public static void setFieldByReflection(Class<?> classType, Object object, String fieldName, Object fieldValue) throws NoSuchFieldException {
        Field field = classType.getDeclaredField(fieldName);
        field.setAccessible(true);
        ReflectionUtils.setField(field, object, fieldValue);
    }
}
