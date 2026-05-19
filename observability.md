## 1. High Latency

* Problem :
  Request taking too long

* Real-world effect

  * Application feels slow
  * Timeouts occur
  * Conversion drops

* Metrics impacted

  * `http_request_duration_seconds` increases
  * p95 / p99 latency increases
  * Throughput eventually decreases

* What to check

  * Grafana

    * p95 / p99 latency graphs
  * SigNoz

    * Slow traces and spans
  * pprof

    * CPU profile (`/debug/pprof/profile`)

* Fix

  * Optimize slow database queries (indexes, query rewrite)
  * Add timeout and retry for external APIs
  * Reduce CPU-heavy logic (optimize or cache)

---

## 2. High Error Rate

* Problem :
  Requests are failing

* Real-world effect

  * Users see errors (500s)
  * System reliability drops

* Metrics impacted

  * `http_requests_total{status=500}` increases
  * Error percentage increases

* What to check

  * Grafana

    * Error rate panels
  * SigNoz

    * Failing traces/spans
  * Logs

    * Exact error messages

* Fix

  * Fix panics / null pointer issues
  * Add retry or fallback for DB/service failures
  * Correct validation or business logic

---

## 3. CPU Spike

* Problem :
  CPU usage suddenly increases

* Real-world effect

  * System slows down
  * Latency increases

* Metrics impacted

  * CPU usage increases
  * Latency increases

* What to check

  * Grafana

    * CPU usage graphs
  * pprof

    * `top` functions consuming CPU

* Fix

  * Remove infinite loops
  * Optimize heavy computations
  * Add caching where needed

---

## 4. Memory Leak

* Problem : 
  Memory usage continuously increases

* Real-world effect

  * Out-of-memory crashes
  * Restart loops

* Metrics impacted

  * Heap allocation continuously increases
  * GC frequency increases

* What to check

  * Go memstats

    * `Alloc`, `HeapAlloc`
  * pprof

    * Heap profile (`/debug/pprof/heap`)
  * Grafana

    * Memory usage graphs

* Fix

  * Release unused objects
  * Fix goroutine leaks
  * Prevent uncontrolled growth of slices/maps

---

## 5. Goroutine Leak

* Problem :
  Too many goroutines accumulating

* Real-world effect

  * Memory and CPU degrade
  * System instability

* Metrics impacted

  * `runtime.NumGoroutine()` increases

* What to check

  * pprof

    * Goroutine dump (`/debug/pprof/goroutine`)
  * Runtime metrics

* Fix

  * Ensure context cancellation
  * Fix blocked channels
  * Stop uncontrolled background workers

---

## 6. Database Bottleneck

* Problem : 
  Database queries are slow or blocked

* Real-world effect

  * Entire system slows down
  * Latency propagates across services

* Metrics impacted

  * API latency increases
  * DB query duration increases

* What to check

  * SigNoz

    * DB spans in traces
  * Prometheus

    * DB latency metrics

* Fix

  * Add indexes
  * Optimize queries
  * Tune connection pool

---

## 7. Downstream Service Slowness

* Problem :
  One service slowing others in a chain

* Real-world effect

  * Full request chain becomes slow

* Metrics impacted

  * Overall trace duration increases
  * Specific span latency increases

* What to check

  * SigNoz

    * Trace waterfall view

* Fix

  * Add timeouts
  * Implement circuit breakers
  * Retry with backoff

---

# EXECUTION LOOP (MENTAL CHECKLIST)

* Detect issue
* Open Grafana

  * Identify spike (latency / errors / CPU / memory)
* Move to SigNoz

  * Find slow service or failing span
* Use pprof

  * Identify exact function/resource issue
* Fix root cause

---

# BIRD’S EYE VIEW

All problems reduce to four signals: **latency, errors, CPU, memory**.

* Metrics (Grafana/Prometheus) tell you *something is wrong*
* Traces (SigNoz) tell you *where it is wrong*
* Profiling (pprof) tells you *why it is wrong*

| Signal  | Step 1 (Detect) | Step 2 (Locate) | Step 3 (Deep Dive)  |
| ------- | --------------- | --------------- | ------------------- |
| Latency | Grafana         | SigNoz (traces) | pprof (CPU/profile) |
| Errors  | Grafana         | SigNoz (traces) | Logs                |
| CPU     | Grafana         | —               | pprof (CPU)         |
| Memory  | Grafana         | —               | pprof (heap)        |

