#!/bin/bash


# Construct the wrk command with the Authorization header
wrk -t12 -c100 -d10s http://127.0.0.1:8081/gms/capabilities




: <<'END_COMMENT'
In go
╰─ wrk -t12 -c100 -d10s http://127.0.0.1:8081/gms/capabilities   ─╯
Running 10s test @ http://127.0.0.1:8081/gms/capabilities
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    23.16ms   59.31ms 667.23ms   96.55%
    Req/Sec   606.51    145.22     0.88k    79.52%
  71287 requests in 10.05s, 58.40MB read
Requests/sec:   7095.29
Transfer/sec:      5.81MB

╰─ wrk -t12 -c100 -d10s http://127.0.0.1:8081/gms/capabilities                                                                                                                                                   ─╯
Running 10s test @ http://127.0.0.1:8081/gms/capabilities
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    15.23ms   21.46ms 374.85ms   97.17%
    Req/Sec   629.07    136.41     0.88k    76.13%
  75069 requests in 10.04s, 61.50MB read
Requests/sec:   7474.63
Transfer/sec:      6.12MB

╰─ wrk -t12 -c100 -d10s http://127.0.0.1:8081/gms/capabilities                                                                                                                                                   ─╯
Running 10s test @ http://127.0.0.1:8081/gms/capabilities
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    15.63ms   22.91ms 381.55ms   97.29%
    Req/Sec   616.61    125.19   848.00     73.18%
  73582 requests in 10.05s, 60.28MB read
Requests/sec:   7320.14
Transfer/sec:      6.00MB

╰─ wrk -t12 -c100 -d10s http://127.0.0.1:8081/gms/capabilities                                                                                                                                                   ─╯
Running 10s test @ http://127.0.0.1:8081/gms/capabilities
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    19.53ms   39.46ms 544.25ms   96.36%
    Req/Sec   609.69    154.71   848.00     78.26%
  72134 requests in 10.04s, 59.09MB read
Requests/sec:   7181.45
Transfer/sec:      5.88MB


In java
╰─ wrk -t12 -c100 -d10s http://127.0.0.1:8081/gms/capabilities                                                                                                                                                   ─╯
Running 10s test @ http://127.0.0.1:8081/gms/capabilities
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    29.01ms   28.15ms 317.22ms   91.69%
    Req/Sec   326.82    120.04   616.00     68.13%
  38748 requests in 10.06s, 45.35MB read
Requests/sec:   3850.27
Transfer/sec:      4.51MB

╰─ wrk -t12 -c100 -d10s http://127.0.0.1:8081/gms/capabilities                                                                                                                                                   ─╯
Running 10s test @ http://127.0.0.1:8081/gms/capabilities
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    26.47ms   33.29ms 471.10ms   95.52%
    Req/Sec   366.33    113.39   626.00     71.27%
  43494 requests in 10.05s, 50.90MB read
Requests/sec:   4329.10
Transfer/sec:      5.07MB

╰─ wrk -t12 -c100 -d10s http://127.0.0.1:8081/gms/capabilities                                                                                                                                                   ─╯
Running 10s test @ http://127.0.0.1:8081/gms/capabilities
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    24.45ms   23.90ms 441.40ms   92.91%
    Req/Sec   373.24    124.70   626.00     63.65%
  44563 requests in 10.05s, 52.15MB read
Requests/sec:   4432.33
Transfer/sec:      5.19MB

╰─ wrk -t12 -c100 -d10s http://127.0.0.1:8081/gms/capabilities                                                                                                                                                   ─╯
Running 10s test @ http://127.0.0.1:8081/gms/capabilities
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    24.38ms   32.36ms 491.71ms   96.68%
    Req/Sec   399.54    103.44   620.00     68.19%
  47362 requests in 10.04s, 55.43MB read
Requests/sec:   4715.47
Transfer/sec:      5.52MB
END_COMMENT