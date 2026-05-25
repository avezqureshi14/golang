At scale we will distribute tasks between go api gateway and nginx , currently our go gatway is only doing everything 

            ┌──────────────┐
            │   Client     │
            └──────┬───────┘
                   ↓
        ┌────────────────────┐
        │       NGINX        │
        │--------------------│
        │ TLS termination    │
        │ Load balancing     │
        │ Rate limiting      │
        └──────┬─────────────┘
               ↓
        ┌────────────────────┐
        │    GO GATEWAY      │
        │--------------------│
        │ JWT auth           │
        │ Routing logic      │
        │ Header injection   │
        │ Aggregation        │
        └──────┬─────────────┘
               ↓
        ┌────────────────────┐
        │   Microservices    │
        └────────────────────┘