#!/bin/bash

# Set the URL for the GET request
url='http://localhost:57176/gms/search'

# Set the headers for the request
headers=(
  'Authorization: Bearer eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI3M2YxNmQ5My0yNDQxLTRhNTAtODhmZi04NTM2MGQ3OGM2YjUiLCJpc3MiOiJodHRwOlwvXC9pYW0tdmlldy5kZWZhdWx0LnN2Yy5jbHVzdGVyLmxvY2FsOjgwODAiLCJleHAiOjE3MTQ2ODgyMDgsImlhdCI6MTcxNDY1MjIwOCwianRpIjoiZWUzMGZjZDEtZjMxZi00OTU5LWFiYzItZWViZGNkZGQxYTBlIiwiY2xpZW50X2lkIjoiY2xpZW50In0.WlpQVp6iyErgXJ3Y5eKggPemUUBWPdxwvsVLbIkPyXYuKKC-v15Aihh8v_M7CfQwLtA3Sl2h4xAHKfd3ocv22Dk9HmeRKan84tKEDyX9XetcR6JbD2fwcjImETeStb4dTnqgwPdUzxXDlfBoon3-OtSGsSHZWWmK96kAun-laKxsQir11XBfFDSjI16-8ERXPwUjh7gyroqdDVSvvoRXzvXRmvlQ9EJwa6OnuuidatFogna8_S6iDLwHnxw_mXmrhYl54kvyfEnWV6K1Pq9xRG7g7PGseh-mG7kutISuOgdPKSFVg4QTGPBDzjynxPmflxcrDO3oOqUCYp1C9PaaMg'
  'User-Agent: insomnia/8.6.1'
)

# Number of requests to send
num_requests=100

# Initialize total time variable
total_time=0

# Send GET requests and measure the time for each request
for ((i=1; i<=$num_requests; i++)); do
  response_time=$(curl -s -w "%{time_total}\n" -o /dev/null -X GET "$url" -H "${headers[@]}")
  total_time=$(echo "$total_time + $response_time" | bc)
done

# Calculate the average time in milliseconds
average_time=$(echo "scale=3; ($total_time / $num_requests) * 1000" | bc)

# Print the results
echo "Sent $num_requests GET requests to $url"
echo "Average time per request: $average_time milliseconds"