# 🔄 Kafka-Redis Webserver (Go)

A backend system built in Go that uses **Kafka** and **Redis** to process, route, and respond to API requests through asynchronous messaging pipelines.

> 🧠 Originally built as a take-home assignment for HashedTokens  
> ⚙️ Designed and implemented by [Dev Bhagavan](https://dev-bhagavan.vercel.app)

---

## 📦 Features

- 🧵 Asynchronous Kafka-based message routing
- 🧠 Non-blocking Go routines for high concurrency
- 💾 Redis-backed fast key-value storage
- 🚀 REST API with clear request/response flow
- 🐳 Runs on Docker (Kafka, Zookeeper, Redis)

---

## 🧱 Architecture

```
POST /push
   ↓
[Goroutine-1] → Kafka Producer → [Kafka Topic: messages]
                                       ↓
                              [Goroutine-2] → Kafka Consumer → Redis

GET /pull/:key
   ↓
[Goroutine-3] ← Redis ← Lookup by key → Response
                         ↑
                    Kafka Producer (audit/log)
```

---

## 📮 Endpoints

### `POST /push`

Send a key-value pair to Kafka

**Request:**

```json
{
  "key": "user123",
  "value": "Hello from Dev"
}
```

**Response:**
```json
{
  "status": "accepted"
}
```

---

### `GET /pull/:key`

Retrieve a value from Redis and push it to Kafka again (audit)

**Response:**
```json
{
  "key": "user123",
  "value": "Hello from Dev"
}
```

---

## ⚙️ Tech Stack

| Component      | Tool / Lib                     |
|----------------|--------------------------------|
| Language       | Go (Golang)                    |
| Kafka Client   | [`sarama`](https://github.com/Shopify/sarama) |
| Redis Client   | [`go-redis`](https://github.com/redis/go-redis) |
| Message Broker | Kafka + Zookeeper              |
| Cache          | Redis                          |
| Containerized  | Docker                         |

---

## 🚀 Running the Project

### 1. Clone and start Docker services

```bash
cd kafka-redis
docker-compose up -d
```

This will start Kafka, Zookeeper, and Redis locally.

### 2. Start the Go server

```bash
go mod tidy
go run main.go
```

### 3. Test with curl

```bash
curl -X POST http://localhost:8080/push \
  -H "Content-Type: application/json" \
  -d '{"key":"user123","value":"Hello"}'

curl http://localhost:8080/pull/user123
```

---

## ✍️ Author

- Dev Bhagavan  
- [Portfolio](https://dev-bhagavan.vercel.app) | [LinkedIn](https://www.linkedin.com/in/dev-bhagavan)  

---
