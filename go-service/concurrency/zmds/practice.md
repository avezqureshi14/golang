### Before coding:

Write 3 lines:

1. Who produces data?
2. Who consumes data?
3. Who stops everything?

If you can’t answer → don’t code yet.

---

# 🪜 Level 1 — Channel discipline (must master first)

## 1. Single producer → single consumer (safe close)

👉 Problem:
Generate numbers 1–100 in a goroutine and consume them.

### Constraints:

* producer must close channel
* consumer must stop safely
* no WaitGroup allowed

### What you learn:

* ownership of close
* basic lifecycle

---

## 2. Multiple producers → single consumer

👉 Problem:
2 goroutines generate numbers concurrently into same channel.

### Constraints:

* channel must close only when BOTH producers finish
* no data race on close

### Forces you to learn:

* “who owns close?” problem

---

# 🪜 Level 2 — Worker pool correctness

## 3. Worker pool (correct shutdown)

👉 Problem:
Build worker pool:

* 5 workers
* 20 jobs
* print results

### Constraints:

* no goroutine leak
* no deadlock
* clean shutdown

### Add twist:

One worker should sleep longer → test blocking

---

## 4. Worker pool with early exit

👉 Problem:
Stop all workers if value == 7 is found.

### You must use:

* context cancel
* proper worker exit

### You learn:

* cancellation propagation (REAL skill)

---

# 🪜 Level 3 — Fan-in / Fan-out correctness

## 5. Fan-in with slow source

👉 Problem:
Merge 3 channels:

* one fast
* one slow
* one medium

### Constraints:

* no blocking on slow producer
* no goroutine leak

---

## 6. Fan-out + aggregation

👉 Problem:
Fetch URLs in parallel and collect results.

### Constraints:

* limit concurrency (like 3 workers max)
* safe result collection
* handle failure of one worker

---

# 🪜 Level 4 — Real-world systems

## 7. Rate limiter (token bucket)

👉 Build:

* 3 req/sec limiter
* burst allowed

### Must include:

* refill goroutine
* cancel support

---

## 8. Scraper v2 (production-grade)

Upgrade your scraper:

### Must add:

* timeout per request
* retry (2 times)
* worker pool
* graceful shutdown (Ctrl+C context cancel)

---

# 🪜 Level 5 — Advanced correctness (where most fail)

## 9. Pipeline with 3 stages

👉 Build:

```
producer → filter → enrich → output
```

Each stage is a worker pool.

### Constraints:

* stage failure cancels entire pipeline
* no goroutine leaks
* each stage independent

---

## 10. Backpressure system

👉 Problem:
Slow consumer should automatically slow producers.

### You learn:

* buffering strategy
* blocking vs dropping design

---

# 🔥 How to actually practice (important)

Don’t just code and move on.

For every problem:

### Step 1: draw flow

```
producer → channel → workers → output
```

### Step 2: inject failure

Ask:

* what if worker is slow?
* what if producer stops?
* what if main exits?

### Step 3: break your code intentionally

If it doesn’t break → problem was too easy

---

# 🧠 Golden rule for your level

Right now your goal is NOT:

* writing more code

Your goal is:

> “predicting how code breaks under timing changes”

That’s concurrency maturity.

---

# 🚀 If you want real acceleration

I can guide you like this:

👉 give you 1 problem
👉 you solve it
👉 I brutally review lifecycle bugs
👉 then next harder problem

That’s how you actually reach interview-grade Go concurrency.

Just say: **“start ladder”**
