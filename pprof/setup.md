| Tool       | Question it answers       |
| ---------- | --------------------------|
| Prometheus | Says something is wrong   |
| Grafana    | What is the trend?        |
| Logs       | What exactly is wrong     |
| Tracing    | Says where u are wrong    |
| pprof      | Why it is wrong           |

```
                 ┌──────────────┐
                 │  Grafana     │
                 │ dashboards   │
                 └─────┬────────┘
                       │
        ┌──────────────┴──────────────┐
        │                             │
 Prometheus                     Logs (Loki/ELK)
 (metrics + alerts)            (structured logs)
        │                             │
        └──────────────┬──────────────┘
                       │
           OpenTelemetry Collector
                       │
        ┌──────────────┴──────────────┐
        │                             │
     Traces (Jaeger/Tempo)      Profiling (pprof/Pyroscope)

```