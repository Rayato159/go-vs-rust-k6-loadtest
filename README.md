<h1>Go vs. Rust Load Testing</h1>

<h2>Start Testing</h2>

<p>***You need to install Grafana K6 first (I installed on Docker)</p>

```bash
docker pull grafana/k6
```

```bash
docker run --rm -i grafana/k6 run --out json=rust_100.json - <script.js
```

<h2>Result</h2>

<h3>Rust</h3>
<p>100 Requests</p>

```bash
checks.........................: 100.00% ✓ 2900      ✗ 0
data_received..................: 380 kB  13 kB/s
data_sent......................: 244 kB  8.0 kB/s
http_req_blocked...............: avg=2.53ms  min=1.3µs   med=3.5µs   max=125.5ms  p(90)=5.8µs   p(95)=10.21µs
http_req_connecting............: avg=2.46ms  min=0s      med=0s      max=125.33ms p(90)=0s      p(95)=0s
http_req_duration..............: avg=41.18ms min=33.89ms med=38.19ms max=112.49ms p(90)=49.93ms p(95)=60.42ms
{ expected_response:true }...: avg=41.18ms min=33.89ms med=38.19ms max=112.49ms p(90)=49.93ms p(95)=60.42ms
http_req_failed................: 0.00%   ✓ 0         ✗ 2900
http_req_receiving.............: avg=71.13µs min=14µs    med=64.3µs  max=564.5µs  p(90)=118.8µs p(95)=141.9µs
http_req_sending...............: avg=24.17µs min=4.7µs   med=12.8µs  max=1.46ms   p(90)=29.2µs  p(95)=54.5µs
http_req_tls_handshaking.......: avg=0s      min=0s      med=0s      max=0s       p(90)=0s      p(95)=0s
http_req_waiting...............: avg=41.09ms min=33.85ms med=38.11ms max=112.23ms p(90)=49.85ms p(95)=60.35ms
http_reqs......................: 2900    95.355153/s
iteration_duration.............: avg=1.04s   min=1.03s   med=1.03s   max=1.19s    p(90)=1.05s   p(95)=1.06s
iterations.....................: 2900    95.355153/s
vus............................: 100     min=100     max=100
vus_max........................: 100     min=100     max=100
```

<p>1000 Requests</p>

```bash
checks.........................: 100.00% ✓ 23675      ✗ 0
data_received..................: 3.1 MB  100 kB/s
data_sent......................: 2.0 MB  64 kB/s
http_req_blocked...............: avg=99.87ms  min=1.6µs   med=4.3µs    max=3.3s     p(90)=7.2µs    p(95)=28.4µs
http_req_connecting............: avg=98.87ms  min=0s      med=0s       max=3.01s    p(90)=0s       p(95)=0s
http_req_duration..............: avg=173.25ms min=32.88ms med=120.88ms max=2.96s    p(90)=305.34ms p(95)=394.34ms
{ expected_response:true }...: avg=173.25ms min=32.88ms med=120.88ms max=2.96s    p(90)=305.34ms p(95)=394.34ms
http_req_failed................: 0.00%   ✓ 0          ✗ 23675
http_req_receiving.............: avg=86.25µs  min=16.8µs  med=76.39µs  max=4.37ms   p(90)=139.46µs p(95)=169.1µs
http_req_sending...............: avg=1.23ms   min=4.8µs   med=16.1µs   max=505.39ms p(90)=45µs     p(95)=74.82µs
http_req_tls_handshaking.......: avg=0s       min=0s      med=0s       max=0s       p(90)=0s       p(95)=0s
http_req_waiting...............: avg=171.93ms min=32.82ms med=120.78ms max=2.75s    p(90)=305.19ms p(95)=394.16ms
http_reqs......................: 23675   759.908034/s
iteration_duration.............: avg=1.27s    min=1.03s   med=1.12s    max=5.06s    p(90)=1.3s     p(95)=1.39s
iterations.....................: 23675   759.908034/s
vus............................: 582     min=582      max=1000
vus_max........................: 1000    min=1000     max=1000
```

<p>10000 Requests</p>

```bash
checks.........................: 61.45% ✓ 14099      ✗ 8842
data_received..................: 1.8 MB 51 kB/s
data_sent......................: 1.2 MB 33 kB/s
http_req_blocked...............: avg=238.78ms min=0s      med=3.3µs    max=23.74s  p(90)=6.5µs    p(95)=165.07ms
http_req_connecting............: avg=238.73ms min=0s      med=0s       max=23.74s  p(90)=0s       p(95)=164.91ms
http_req_duration..............: avg=857.69ms min=0s      med=155.01ms max=12.49s  p(90)=742.97ms p(95)=10.38s
{ expected_response:true }...: avg=1.39s    min=33.67ms med=387.17ms max=12.49s  p(90)=3.12s    p(95)=11.23s
http_req_failed................: 38.54% ✓ 8842       ✗ 14099
http_req_receiving.............: avg=53.76µs  min=0s      med=42.9µs   max=3.27ms  p(90)=121.8µs  p(95)=146.4µs
http_req_sending...............: avg=49.02µs  min=0s      med=12.4µs   max=15.53ms p(90)=37.19µs  p(95)=74.5µs
http_req_tls_handshaking.......: avg=0s       min=0s      med=0s       max=0s      p(90)=0s       p(95)=0s
http_req_waiting...............: avg=857.59ms min=0s      med=154.89ms max=12.49s  p(90)=742.92ms p(95)=10.38s
http_reqs......................: 22941  634.989965/s
iteration_duration.............: avg=13.68s   min=1.03s   med=1.72s    max=31.45s  p(90)=31.12s   p(95)=31.15s
iterations.....................: 22941  634.989965/s
vus............................: 5      min=0        max=10000
vus_max........................: 10000  min=2497     max=10000
```

<h3>Go</h3>
<p>100 Requests</p>

```bash
checks.........................: 100.00% ✓ 2900      ✗ 0
data_received..................: 421 kB  14 kB/s
data_sent......................: 241 kB  7.9 kB/s
http_req_blocked...............: avg=2.2ms   min=1.4µs   med=3.7µs   max=114.12ms p(90)=6.5µs   p(95)=14.5µs
http_req_connecting............: avg=2.19ms  min=0s      med=0s      max=113.92ms p(90)=0s      p(95)=0s
http_req_duration..............: avg=42.76ms min=33.66ms med=38.34ms max=187.1ms  p(90)=49.13ms p(95)=58.53ms
{ expected_response:true }...: avg=42.76ms min=33.66ms med=38.34ms max=187.1ms  p(90)=49.13ms p(95)=58.53ms
http_req_failed................: 0.00%   ✓ 0         ✗ 2900
http_req_receiving.............: avg=73.14µs min=12.7µs  med=62.7µs  max=769µs    p(90)=123.3µs p(95)=149.9µs
http_req_sending...............: avg=23.37µs min=5µs     med=13.5µs  max=1.43ms   p(90)=26.8µs  p(95)=48.42µs
http_req_tls_handshaking.......: avg=0s      min=0s      med=0s      max=0s       p(90)=0s      p(95)=0s
http_req_waiting...............: avg=42.66ms min=33.6ms  med=38.25ms max=186.99ms p(90)=49.03ms p(95)=58.47ms
http_reqs......................: 2900    94.962479/s
iteration_duration.............: avg=1.04s   min=1.03s   med=1.03s   max=1.18s    p(90)=1.05s   p(95)=1.1s
iterations.....................: 2900    94.962479/s
vus............................: 100     min=100     max=100
vus_max........................: 100     min=100     max=100
```

<p>1000 Requests</p>

```bash
checks.........................: 100.00% ✓ 23779      ✗ 0
data_received..................: 3.4 MB  111 kB/s
data_sent......................: 2.0 MB  63 kB/s
http_req_blocked...............: avg=28.41ms  min=1.4µs   med=4.7µs    max=1.3s    p(90)=7.8µs    p(95)=35.79µs
http_req_connecting............: avg=28.38ms  min=0s      med=0s       max=1.3s    p(90)=0s       p(95)=0s
http_req_duration..............: avg=250.03ms min=32.41ms med=191.58ms max=1.14s   p(90)=493.97ms p(95)=560.1ms
{ expected_response:true }...: avg=250.03ms min=32.41ms med=191.58ms max=1.14s   p(90)=493.97ms p(95)=560.1ms
http_req_failed................: 0.00%   ✓ 0          ✗ 23779
http_req_receiving.............: avg=100.5µs  min=16.1µs  med=86.3µs   max=20.14ms p(90)=154.9µs  p(95)=191.7µs
http_req_sending...............: avg=33.37µs  min=5.7µs   med=17.8µs   max=8.11ms  p(90)=50.1µs   p(95)=75.8µs
http_req_tls_handshaking.......: avg=0s       min=0s      med=0s       max=0s      p(90)=0s       p(95)=0s
http_req_waiting...............: avg=249.9ms  min=32.31ms med=191.45ms max=1.14s   p(90)=493.91ms p(95)=559.99ms
http_reqs......................: 23779   763.813911/s
iteration_duration.............: avg=1.27s    min=1.03s   med=1.19s    max=2.85s   p(90)=1.49s    p(95)=1.57s
iterations.....................: 23779   763.813911/s
vus............................: 569     min=569      max=1000
vus_max........................: 1000    min=1000     max=1000
```

<p>10000 Requests</p>

```bash
checks.........................: 64.89% ✓ 16026      ✗ 8670
data_received..................: 2.3 MB 64 kB/s
data_sent......................: 1.3 MB 37 kB/s
http_req_blocked...............: avg=216.84ms min=0s      med=3.7µs    max=17.41s  p(90)=7µs     p(95)=434.31ms
http_req_connecting............: avg=216.76ms min=0s      med=0s       max=17.41s  p(90)=0s      p(95)=434.2ms
http_req_duration..............: avg=952.34ms min=0s      med=321.63ms max=12.53s  p(90)=1.66s   p(95)=10.3s
{ expected_response:true }...: avg=1.46s    min=36.04ms med=412.26ms max=12.53s  p(90)=3.22s   p(95)=10.41s
http_req_failed................: 35.10% ✓ 8670       ✗ 16026
http_req_receiving.............: avg=61.16µs  min=0s      med=47.2µs   max=4.11ms  p(90)=127.4µs p(95)=156.82µs
http_req_sending...............: avg=181.55µs min=0s      med=13.2µs   max=34.19ms p(90)=45.8µs  p(95)=82.69µs
http_req_tls_handshaking.......: avg=0s       min=0s      med=0s       max=0s      p(90)=0s      p(95)=0s
http_req_waiting...............: avg=952.09ms min=0s      med=321.51ms max=12.53s  p(90)=1.66s   p(95)=10.3s
http_reqs......................: 24696  680.413062/s
iteration_duration.............: avg=12.73s   min=1.03s   med=1.63s    max=31.5s   p(90)=31.15s  p(95)=31.19s
iterations.....................: 24696  680.413062/s
vus............................: 92     min=0        max=10000
vus_max........................: 10000  min=2950     max=10000
```

<h2>Summary</h2>

<p>Rust faster than Go but uncertain</p>