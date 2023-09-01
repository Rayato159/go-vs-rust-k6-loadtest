<h1>Go vs. Rust Load Testing</h1>

<h2>Start Testing</h2>

<p>***You need to install Grafana K6 first (I installed on Docker)</p>

```bash
docker pull grafana/k6
```

```bash
docker run --rm -i grafana/k6 run - <script.js
```

<h2>Compute Engine Spec</h2>
CPU E2-micro 2 vCPU 1 core, RAM 1 GB, Balanced persistent disk

<h2>Result</h2>

<p>Rust 250, 500, 1000 requests</p>

```bash
#250 Requests
data_received..................: 923 kB  30 kB/s
data_sent......................: 585 kB  19 kB/s
http_req_blocked...............: avg=9.88ms  min=1µs     med=3.7µs   max=499.89ms p(90)=11µs     p(95)=22.17µs
http_req_duration..............: avg=71.55ms min=33.49ms med=44.51ms max=499.03ms p(90)=116.33ms p(95)=206.49ms
vus............................: min=64       max=250

#500 Requests
data_received..................: 1.8 MB  59 kB/s
data_sent......................: 1.2 MB  38 kB/s
http_req_blocked...............: avg=8.56ms  min=1.1µs   med=3.6µs   max=449.49ms p(90)=6.1µs    p(95)=15.09µs
http_req_duration..............: avg=74.73ms min=32.57ms med=41.99ms max=509.74ms p(90)=186.39ms p(95)=281.23ms
vus............................: min=66       max=500

#1000 Requests
data_received..................: 2.0 MB  65 kB/s
data_sent......................: 1.3 MB  41 kB/s
http_req_blocked...............: avg=89.68ms  min=1.2µs   med=3.7µs  max=15.35s p(90)=7µs     p(95)=476.18ms
http_req_duration..............: avg=883.84ms min=34.22ms med=1.07s  max=1.81s  p(90)=1.39s   p(95)=1.48s
vus............................: min=95       max=1000
```

<p>Go 250, 500, 1000 requests</p>

```bash
#250 Requests
data_received..................: 977 kB  32 kB/s
data_sent......................: 566 kB  18 kB/s
http_req_blocked...............: avg=5.49ms   min=1.4µs   med=4.2µs   max=259.29ms p(90)=7µs      p(95)=22.3µs
http_req_duration..............: avg=121.09ms min=33.04ms med=50.11ms max=689.9ms  p(90)=329.37ms p(95)=458.56ms
vus............................: min=250      max=250

#500 Requests
data_received..................: 1.7 MB  55 kB/s
data_sent......................: 1000 kB 32 kB/s
http_req_blocked...............: avg=20.6ms   min=1µs     med=3.8µs    max=964.68ms p(90)=6.4µs    p(95)=25.63µs
http_req_duration..............: avg=275.63ms min=33.71ms med=170.2ms  max=1.05s    p(90)=678.4ms  p(95)=830.93ms
vus............................: min=307      max=500

#1000 Requests
data_received..................: 2.3 MB  72 kB/s
data_sent......................: 1.3 MB  42 kB/s
http_req_blocked...............: avg=59.23ms  min=1.1µs   med=3.6µs  max=1.8s   p(90)=6.6µs   p(95)=326.87ms
http_req_duration..............: avg=935.49ms min=33.05ms med=1.04s  max=1.8s   p(90)=1.22s   p(95)=1.39s
vus............................: min=348      max=1000
```

<h2>Summary</h2>

<p>Rust faster than Go</p>

```text
rust_http_duration_avg = (71.55 + 74.73 + 883.84) ms / 3 = 343.37 ms
go_http_duration_avg   = (121.09 + 275.63 + 935.49) ms / 3 = 444.07 ms

rust is faster than go by -> 100 % - ((343.37/444.07) * 100 %) = 33 % for avg of 250, 500 and 100 requests
```