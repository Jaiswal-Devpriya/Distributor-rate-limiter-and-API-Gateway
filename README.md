# Distributed Rate Limiter and API Gateway

![Go](https://img.shields.io/badge/Go-1.26-blue)
![Redis](https://img.shields.io/badge/Redis-7-red)
![CI](https://github.com/Jaiswal-Devpriya/Distributor-rate-limiter-and-API-Gateway/actions/workflows/ci.yml/badge.svg)

A production-style API Gateway built in Go with a Redis-backed rate limiter. The gateway enforces per-client request quotas using shared Redis state, making rate limits consistent across gateway instances.

## Tech Stack

* Go
* Redis
* Docker Compose
* k6 Load Testing
* HTTP Middleware
* REST API

## Features

* API Gateway with request routing
* Redis-backed per-client rate limiting
* Client identification using `X-Client-ID` header
* Dockerized Go gateway and Redis services
* Health check endpoint
* k6 load testing for performance validation

## Architecture

<img width="458" height="633" alt="Architecture-diagram" src="https://github.com/user-attachments/assets/d2a1312e-5ed2-4975-a418-254805114c0f" />

### Request Flow

1. Client sends request to API Gateway
2. Rate limiting middleware checks request quota
3. Redis stores and tracks per-client request counts
4. Allowed requests are forwarded to route handlers
5. Exceeded requests return HTTP 429

## Docker Deployment

<img width="1169" height="154" alt="docker-running" src="https://github.com/user-attachments/assets/434504b8-f0ba-451c-8016-46661650c251" />
The API Gateway and Redis services are containerized using Docker Compose.

## API Endpoints

### Health Check

```bash
GET /health
```

Response:

```json
{
  "status": "ok"
}
```

### Gateway Data Endpoint

```bash
GET /api/data
```

Response when allowed:

```json
{
  "message": "request allowed through API gateway",
  "timestamp": "2026-06-13T16:11:29Z"
}
```

Response when rate limit is exceeded:

```json
{
  "error": "rate limit exceeded"
}
```

## Run Locally with Docker

```bash
docker compose up --build
```

Test the health endpoint:

```bash
curl http://localhost:8080/health
```

Test rate limiting:

```bash
for i in {1..7}; do curl -H "X-Client-ID: devpriya" http://localhost:8080/api/data; echo; done
```

Expected behavior: the first 5 requests are allowed, and later requests return `rate limit exceeded`.

## Rate Limiting Demonstration

<img width="967" height="498" alt="rate-limiter-working" src="https://github.com/user-attachments/assets/38ae03b0-6528-4301-987c-6975acf6f978" />

## Load Testing

Run the k6 load test:

```bash
k6 run k6-load-test.js
```

Latest test result:

```text
20 virtual users
30 seconds
600 total requests
100% expected responses
p95 latency: 8.84ms
```
## Continuous Integration


<img width="1508" height="638" alt="github-actions-success" src="https://github.com/user-attachments/assets/3db45835-cfe7-444a-bd12-a7e56142d164" />
GitHub Actions automatically builds the project on every push.

## Project Structure

```text
.
├── cmd/gateway/main.go
├── internal/limiter/token_bucket.go
├── internal/middleware/rate_limit.go
├── internal/proxy/router.go
├── docker-compose.yml
├── Dockerfile
├── k6-load-test.js
├── go.mod
└── README.md
```

## Resume Highlights

* Implemented a Redis-backed distributed rate limiter in Go, enforcing per-client request quotas through API Gateway middleware and shared rate-limit state.
* Containerized the API Gateway and Redis services with Docker Compose and validated performance using k6 load testing with 600 requests, 100% expected responses, and 8.84ms p95 latency.
