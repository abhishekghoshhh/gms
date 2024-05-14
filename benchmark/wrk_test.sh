#!/bin/bash

# Set the token as an environment variable
TOKEN=""

# Construct the wrk command with the Authorization header
# wrk -t12 -c100 -d10s -H "Authorization: Bearer $TOKEN" http://127.0.0.1:8081/gms/search


wrk -t12 -c100 -d10s http://127.0.0.1:8081/gms/capabilities