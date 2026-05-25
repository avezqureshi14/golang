# Done vs Context

## Final Rule

Use `context` by default.
Use `done` channels only when things are simple and local.


## Use `done` channel when:

* You only need a **stop signal**
* Scope is **small and local**
* Goroutines are **internal (fully controlled by you)**
* No need for **timeout, deadline, or metadata**

**Mental model:** “just tell workers to stop”


## Avoid `done` when:

* There are **multiple layers** (A → B → C call chain)
* You need **timeouts or deadlines**
* You interact with **APIs, databases, or HTTP**
* You need to pass **request-scoped data**


# Context vs WaitGroup

## Core Difference

* `context` → **stops work**
* `WaitGroup` → **waits for work to finish**

They are not interchangeable.


## Context

**Purpose:** Signal cancellation

Tells goroutines:

> “Stop what you’re doing”

**Does NOT guarantee:**

* Goroutine has exited
* Cleanup is completed
* Resources are released


## WaitGroup

**Purpose:** Synchronization (wait for completion)

Guarantees:

* All goroutines have **finished execution**
* Creates a **blocking point** (`wg.Wait()`)
* Prevents **premature program exit**


## WaitGroup Does NOT:

* Stop or cancel goroutines
* Handle timeouts or deadlines
* Prevent goroutine leaks
* Guarantee correctness (misuse can cause deadlocks)


Who Gurantees what ?
- Goroutine has exited
> Guaranteed by using a WaitGroup (and ensuring the goroutine calls Done())
- Cleanup is completed
> Guaranteed by using defer inside the goroutine
- Resources are released
> Guaranteed by writing context-aware code + proper cleanup logic (usually via defer), not by context alone

### If all of this is true then no context is needed
- Goroutine has finite work
- No external dependency (network, DB, blocking channel, etc.)
- It cannot get stuck
- It will always complete


### Web Scraper 
-  create http client with timeout
-  send http get request
-  ensure response body is closed
-  parse html
-  extract title


### code must be in this below flow
- 1. Worker flow
reads jobs
exits on close
- 2. Producer flow
sends jobs
closes jobs
- 3. Collector flow
reads results
closes when workers done


1. timeout inside loop is not correct it leads to memory leak 


