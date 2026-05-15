## 1. Find the culprit

```bash
(pprof) top
```

👉 Ask:

* Is any function **>50%**?

  * YES → 🎯 that’s your problem
  * NO → look at top 3 combined

---

## 2. Go to exact code

```bash
(pprof) list <function_name>
```

👉 Ask:

* Which **line** has highest allocation?
* Is it inside a **loop**?

---

## 3. Identify pattern (THIS is your table applied)

| What you see in code     | Meaning 🚨            | Action ✅              |
| ------------------------ | --------------------- | --------------------- |
| `append()` in loop       | slice growing forever | limit / reset / reuse |
| `map[key] = value` loop  | map growing forever   | delete / cap size     |
| `make([]...)` frequently | too many allocations  | reuse buffer          |
| `string + string` loop   | new strings each time | use `strings.Builder` |
| global variable storing  | never freed memory    | avoid global growth   |

---

## 4. Sanity check

Ask yourself:

* ❓ Does this data **ever get cleared?**
* ❓ Is this inside an **infinite loop / long-running process?**
* ❓ Am I **holding references** accidentally?

If answer = NO → 💀 leak confirmed

---

## 5. Validate fix

After change:

```bash
go tool pprof ...
(pprof) top
```

👉 Expect:

* Problem function % ↓↓↓
* Memory stable

---

# ⚡ Ultra-short version (remember this)

```text
top → find big %
list → find line
loop? → growing?
not freed? → leak
```

---




