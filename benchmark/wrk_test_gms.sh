#!/bin/bash

# Set the token as an environment variable
TOKEN="eyJraWQiOiJyc2ExIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI3M2YxNmQ5My0yNDQxLTRhNTAtODhmZi04NTM2MGQ3OGM2YjUiLCJpc3MiOiJodHRwOlwvXC9pYW0tdmlldy5kZWZhdWx0LnN2Yy5jbHVzdGVyLmxvY2FsOjgwODAiLCJleHAiOjE3MTU2OTYzMzQsImlhdCI6MTcxNTY5MjczNCwianRpIjoiOTUzYTg1YTQtYzI5OS00NzI5LTg3MGEtNmI3Nzc1YzAxY2JkIiwiY2xpZW50X2lkIjoiY2xpZW50In0.Gklfy5PI2KxtxWANbq0XFdcMBTneHa4CqjDKMMaR-EuXgkIPC5bhn0lYh2CeTC_xcjed3N8BKycy01C9YFjZmg3TTg19JH9DEAHvXL8sBZEaE5vTUyb-c5D_WMJBeIYbOFQ4s-boldMW2LKT_UIdjHANpeTHvjeRQywPiFkr_bsItGhfVkLo-ohnlHbyPzbUslI9pkS_wJ4mp8OBExmo3Mc2pdJ7DY8_m6lEHGwj0EVYxiRPs4SRdT9L_0pXLQl2KVx8neVo3ZDd6USrIFIvYso03i-nBTESYcfFL4LeRRtIniueuUItdFt_02m2QVNICGzTr4K5lXTkPfR8edvmGw"

# Construct the wrk command with the Authorization header
wrk -t5 -c5 -d15s -H "Authorization: Bearer $TOKEN" http://127.0.0.1:8081/gms/search



: <<'END_COMMENT'
In Go

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   497.16ms  312.51ms   1.86s    89.02%
    Req/Sec     2.06      1.06     5.00     62.26%
  159 requests in 15.03s, 15.68KB read
Requests/sec:     10.58
Transfer/sec:      1.04KB

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   443.59ms  281.88ms   1.91s    93.55%
    Req/Sec     2.34      0.98     5.00     72.16%
  176 requests in 15.02s, 17.36KB read
  Socket errors: connect 0, read 0, write 0, timeout 1
Requests/sec:     11.71
Transfer/sec:      1.16KB

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   351.10ms   87.01ms 676.46ms   74.63%
    Req/Sec     2.59      0.98     5.00     74.15%
  205 requests in 15.02s, 20.22KB read
Requests/sec:     13.64
Transfer/sec:      1.35KB

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   330.04ms   73.73ms 595.88ms   69.78%
    Req/Sec     2.75      0.85     5.00     77.78%
  225 requests in 15.04s, 22.19KB read
Requests/sec:     14.96
Transfer/sec:      1.48KB

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   429.03ms  300.80ms   1.94s    91.96%
    Req/Sec     2.54      1.00     5.00     72.58%
  186 requests in 15.02s, 18.35KB read
Requests/sec:     12.38
Transfer/sec:      1.22KB

In java

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.01s   321.81ms   1.95s    77.94%
    Req/Sec     0.60      0.49     1.00     60.00%
  70 requests in 15.04s, 20.17KB read
  Socket errors: connect 0, read 0, write 0, timeout 2
Requests/sec:      4.65
Transfer/sec:      1.34KB

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   826.67ms  261.55ms   2.00s    86.42%
    Req/Sec     0.72      0.45     1.00     72.29%
  83 requests in 15.02s, 23.91KB read
  Socket errors: connect 0, read 0, write 0, timeout 3
Requests/sec:      5.53
Transfer/sec:      1.59KB

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   779.57ms  212.61ms   1.50s    85.26%
    Req/Sec     0.83      0.38     1.00     83.16%
  95 requests in 15.03s, 27.37KB read
Requests/sec:      6.32
Transfer/sec:      1.82KB

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   699.80ms  154.05ms   1.23s    90.48%
    Req/Sec     0.90      0.29     1.00     90.48%
  105 requests in 15.02s, 30.25KB read
Requests/sec:      6.99
Transfer/sec:      2.01KB

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   762.29ms  188.78ms   1.75s    88.64%
    Req/Sec     0.87      0.34     1.00     86.67%
  90 requests in 15.02s, 25.93KB read
  Socket errors: connect 0, read 0, write 0, timeout 3
Requests/sec:      5.99
Transfer/sec:      1.73KB

Running 15s test @ http://127.0.0.1:8081/gms/search
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   709.50ms  131.76ms   1.22s    87.62%
    Req/Sec     0.93      0.25     1.00     93.33%
  105 requests in 15.02s, 30.25KB read
Requests/sec:      6.99
Transfer/sec:      2.01KB
END_COMMENT