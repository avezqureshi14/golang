

# 1. Goroutine

### Rules:

* Start with `go`
* Never assume it finishes unless you synchronize it

```go
go func() {
    // work
}()
```

### Must remember:

* If main exits → goroutines die
* Always pair with `WaitGroup` or `channel` if you care about result

---

# 2. WaitGroup

### Rules:

* Create before starting goroutines
* `Add()` BEFORE starting goroutine
* `Done()` inside goroutine (always `defer`)
* `Wait()` in main/parent

```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()

wg.Wait()

```

Important :
WaitGroup ensures you don’t exit until all goroutines have finished, so you can verify they actually stop.

If any goroutine (like your generator) gets stuck, wg.Wait() blocks, exposing the bug instead of silently leaking.

### Must NOT:

* Call `Add()` inside goroutine (race risk)
* Forget `Done()` → deadlock

---

# 3. Channel

### Rules:

* Create before use

```go
ch := make(chan int)
```

* Send and receive must match

```go
ch <- 10
v := <-ch
```

* Close only by sender (important rule)

```go
close(ch)
```
* Reading whether the channel is closed or not
```go
	case v, ok := <-channel1:
    //if channel if closed this ok will be false and during that time we have to return 
    		if !ok {
				return
			}
```
* If the channel is being filled by many producers then in this case close the channel by creating a separate go routine after wg.Wait inside same go routine , because over here in this case we can't just close one channel after it finished putting value in channel other go routines are also there so we need to wait unitl everyone is done 

* Channels closing need to be done by someone who knows when the go routine is going to get completed , in case of single producer it is closed by producer itself because it knows it will be completed , now lets say if there are multiple producer now in this case no single producer can close the channel because other producer might be sending data on that channel so the person who knows over here that when go rotuine are going to get completed is wg.wait now we will run this wg.wait is separate go routine and then we close the channel

### Must remember:

* Never send on closed channel
* Never close from receiver side

---

# 4. Buffered Channel

### Rules:

```go
ch := make(chan int, 5)
```

* Sends block only when buffer full
* Receives block only when empty

### Use when:

* You want decoupling between producer & consumer speed

---

# 5. select

### Rules:

* Must always have at least one case
* Used to wait on multiple channels

* Check number of blocking operations and as per that use nested selects 
- Input boundary → waiting for data
- Output boundary → pushing data

```go
select {
case v := <-ch1:
case v := <-ch2:
default:
}
```

### Must remember:

* Without `default` → can block
* First ready case executes
default:
	    // nobody is reading → back off
		time.Sleep(10 * time.Millisecond)
---

# 6. Mutex

### Rules:

* Lock before shared state access
* Unlock always (use `defer`)

```go
mu.Lock()
defer mu.Unlock()

counter++
```

### Must NOT:

* Copy mutex
* Forget unlock → deadlock

---

# 7. RWMutex

### Rules:

* Multiple readers OR one writer

```go
mu.RLock()
// read
mu.RUnlock()

mu.Lock()
// write
mu.Unlock()
```

---

# 8. Atomic

### Rules:

* Only for simple numeric/state operations
* No locking required

```go
atomic.AddInt64(&counter, 1)
```

### Must remember:

* Not for complex structs
* Not a replacement for Mutex always

---

# 9. Context

### Rules:

* Always pass as first argument
* Create from parent context
* Always cancel when done

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

### In goroutines:

```go
select {
case <-ctx.Done():
    return
}
```

---

# 10. Once

### Rules:

* Runs exactly once
* Used for init logic

```go
once.Do(func() {
    // init
})
```

---

# 11. Channels Close Rule (VERY IMPORTANT)

### Rules:

* ONLY sender closes channel
* Closing signals “no more values”
* make sure all the channels created are closed using defer , but don't run defer inside a loop always called inside go routine

```go
close(ch)
```

### Receiver pattern:

```go
for v := range ch {
}
```

---

# 12. Worker Pool (pattern rule)

### Rules:

* Fixed number of workers
* One input channel
* Workers read until channel closes


```go
for job := range jobs {
}
```

---

# 13. Pipeline (pattern rule)

### Rules:

* Each stage = goroutine + channel
* Always pass output channel forward
* Close downstream properly

```
stage1 → stage2 → stage3
```

---

# 🧠 One-line memory cheat sheet

* Goroutine → "start async"
* WaitGroup → "wait for many"
* Channel → "communicate"
* select → "wait multiple"
* Mutex → "protect shared memory"
* RWMutex → "many readers, one writer"
* Atomic → "fast counters"
* Context → "cancel/timeouts"
* Once → "run once"
* Close channel → "no more data"

---

### General 
> If creating a channel inside a function run go routine over their itself and also close the channel 

```go
func generator(ctx context.Context) <-chan Pair {

	out := make(chan Pair)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				select {
				case <-ctx.Done():
					return
				case out <- Pair{i, j}:
				}
			}
		}
	}()
	return out
}
```