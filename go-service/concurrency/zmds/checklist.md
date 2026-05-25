# Go Concurrency Design Checklist

Use this checklist **before writing goroutines, channels, worker pools, fan-out, or pipelines**.

```
Every goroutine must answer:

- When do I stop?
- Who stops me?
- What happens when something breaks?
```


## Quick Mnemonic: C-S-O-E-B

| Letter | Meaning | Ask Yourself |
|--------|---------|--------------|
| C | Create | Who creates goroutines? |
| S | Stop | Who stops them? |
| O | Ownership | Who owns send / close / consume? |
| E | Early Exit | What if flow stops unexpectedly? |
| B | Backpressure | What happens under load? |



# 1. Creation

Identify exactly **who starts goroutines**.

Possible sources:

- `main()`
- worker pool manager
- producer
- fan-out layer
- scheduler / orchestrator

### Rule

If you cannot clearly answer who creates the goroutines, stop writing code.



# 2. Stop Conditions

Every goroutine must have a defined exit path.

Possible stop conditions:

- channel closed
- `context.Context` canceled
- finite loop completed
- error condition handled
- parent shutdown signal

>  Channels used for synchronization don’t need to be closed. Closing is only required for signaling completion in data streams, typically with a single sender

### Rule

If a goroutine has no exit path, it is a leak.



# 3. Ownership of Data Flow

Define responsibility for each channel.

Ask:

- Who sends on the channel?
- Who closes the channel?
- Who consumes from the channel?

### Rule

Only one owner should close a channel.

Never close a channel from multiple goroutines.



# 4. Early Exit Handling

Design behavior when normal flow breaks.

Ask:

- What if consumer stops early?
- What if producer crashes?
- What if timeout happens?
- What if one worker fails?
- What if context is canceled mid-stream?

| Problem              | Tool                     | What it enforces     | Required Guarantee        |
| -------------------- | ------------------------ | -------------------- | ------------------------- |
| Consumer stops early | context + select         | stop producer safely | No worker should block    |
| Producer crashes     | errgroup                 | propagate failure    | Workers must not hang     |
| Timeout happens      | context.WithTimeout      | auto shutdown        | Everything must stop fast |
| Worker fails         | errgroup / recover       | fail-fast or isolate | Defined failure strategy  |
| Context canceled     | ctx.Done() in every loop | stop all work        | Clean shutdown            |

> ### sync.Once. <br>
> It’s great for the scenario: "What if 5 workers fail at once, but I only want to log the error or close the 'shutdown' channel once?"



### Rule

If early exits are not designed, deadlocks and stuck goroutines appear later.



# 5. Backpressure Strategy

Every concurrent system has pressure points.

Examples:

- buffered channel becomes full
- worker becomes slow
- downstream API slows down
- DB becomes saturated

Choose one strategy:

- block producer
- drop work
- retry with limits
- scale workers
- cancel request
- timeout and fail fast

### Rule

No backpressure strategy = unpredictable behavior.



# Final Pre-Code Review

Before coding, answer all 5:

- Who creates goroutines?
- Who stops them?
- Who owns each channel?
- What happens on early exit?
- What happens under pressure?

If any answer is unclear, redesign first.



# One-Line Reminder

Concurrency bugs usually come from unclear lifecycle, unclear ownership, or missing shutdown behavior.

