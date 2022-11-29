**Description :**
Group Membership service is an IVOA-compatible wrapper service which is used to query the group membership questions to INDIGO-IAM to know if any user is part of any set of groups or not.

**Setup of Group Membership service :**
1. Check the iam-login-service from https://github.com/indigo-iam/iam/tree/master/iam-login-service. Run db, Iam-login-service in docker and in same network. 
   1. ``` docker network create iam ```
   2. ```
      docker run \
      --name db \
      --network iam \
      -e MYSQL_DATABASE=iam \
      -e MYSQL_USER=iam \
      -e MYSQL_ROOT_PASSWORD=pwd \
      -e MYSQL_PASSWORD=pwd \
      -d -p 3306:3306 mysql
      ```
   3. ```
      docker run -d \
      --name iam-login-service \
      --net=iam -p 8080:8080 \
      --env-file=path-to-iam-env-file \
      -v container-mount-path-keystore.jwks:local-path-to-keystore.jwks:ro \
      --restart unless-stopped \
      indigoiam/iam-login-service:v1.8.0
      ```
   For reference please check https://indigo-iam.github.io/v/v1.8.0/docs/getting-started/docker/
   <br/>
   <br/>
2. Check that IAM is successfully connected with DB or not and register some dummy user and group, and add users to the group.
3. Run Iam-test-client in either docker or as jar as it is an independent service which will just connect to iam-login-service. Make sure everything is working, and we are able to log in with the dummy users in Iam-test-client successfully. You can find the iam-test-client in https://github.com/indigo-iam/iam/tree/master/iam-test-client
4. Run GMS in either docker or as JAR file and specify the path the env file
   1. Using docker:
      1. build docker image of GMS : 
         ```docker build -t gms/gms:latest .```
      2. Run the gms container : 
      ```
         docker run -d \
         --name gms \
         --net=iam -p 443:443 \
         --env-file=path-to-gms-env-file \
         --restart unless-stopped \
         gms/gms:latest
      ```
   2. Using JAR:
      1. Build the jar file using maven : ``` mvn clean verify -DskipTests ```
      2. Change the gms keystore details in .env file 
      3. Export the env file : ```export $(xargs < .env) ```
      4. Run the JAR: ``` java -jar target/group-management-service-0.0.1.jar ```