#!/bin/bash

# Set the URL for the requests
url="http://127.0.0.1:57176/gms/capabilities"

# Number of requests to send
num_requests=100

# Initialize total time variable
total_time=0

# Loop to send GET requests and measure time
for ((i=1; i<=$num_requests; i++)); do
    start=$(date +%s.%N)
    curl -s "$url" > /dev/null
    end=$(date +%s.%N)
    duration=$(echo "$end - $start" | bc)
    total_time=$(echo "$total_time + $duration" | bc)
done

# Calculate average time
average_time=$(echo "scale=2; $total_time / $num_requests * 1000" | bc)

# Print the results
echo "Sent $num_requests GET requests to $url"
echo "Average time: $average_time mili seconds"