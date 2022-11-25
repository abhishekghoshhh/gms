FROM openjdk:17-jdk-slim
COPY target/group-management-service-0.0.1.jar application.jar
EXPOSE 443
ENTRYPOINT ["java","-jar","application.jar","${GMS_JAVA_OPTS}"]