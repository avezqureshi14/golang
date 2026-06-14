# PART 1 — Memory Debugging Mental Model

Forget commands for a second.

### Core question:

> **Why is memory growing?**

There are ONLY 3 real causes:

---

## 1. You are CREATING too much

Pattern:

* `make()`, `append()`, `string concat`

Signal in pprof:

* high allocation %
* many small allocs

Fix:

* reuse
* pooling
* builders

---

## 2. You are NOT RELEASING

Pattern:

* maps growing forever
* slices never reset
* globals holding data

Signal:

* memory keeps increasing over time

Fix:

* delete / reset
* limit size
* drop references

---

## 3. You are HOLDING REFERENCES

This is where most people fail.

Example:

```go
bigSlice := make([]byte, 1e6)
small := bigSlice[:10]
```

You think:

> “I’m using only 10 bytes”

Reality:

> entire 1MB is still in memory

---

##  So your new debugging flow becomes:

```text
Memory high?

→ Are we allocating too much?
→ Or not freeing?
→ Or holding references?
```

pprof is just a **tool to confirm**, not the thinking itself

---

# PART 2 — Your pprof flow (refined)

Your version is already good. I’ll sharpen it:

---

##  Step 1: Find hotspot

```bash
(pprof) top
```

Ask:

> “Who is allocating the most?”

---

##  Step 2: Zoom into code

```bash
(pprof) list <func>
```

Ask:

> “Which line creates memory?”

---

##  Step 3: Classify the issue

Ask:

* Is this inside loop? → growth
* Is it accumulating? → leak
* Is it copying? → waste

---

##  Step 4: Ask the killer question

> **“When does this memory die?”**

If answer is:

* “never” → leak 💀
* “later” → maybe ok
* “quickly” → fine

---

##  Step 5: Fix + verify

---

# 💥 Ultra mental model (BETTER than your version)

```text
top → who allocates
list → where
loop? → grows?
dies? → or stuck?
```

---