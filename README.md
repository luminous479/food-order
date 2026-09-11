# Food Order Worker Pool

A concurrent food-order processing system built with **Go**, demonstrating the Worker Pool pattern and core concurrency primitives.

## Overview

This project simulates a restaurant order-processing system where multiple workers process incoming orders concurrently.

Orders are distributed through a Go channel to a fixed number of workers. A `sync.WaitGroup` coordinates worker completion, while `sync.Mutex` protects shared processing statistics from concurrent access.

## Architecture

```text
                    Order Producer
                         │
                         ▼
                  ┌──────────────┐
                  │ Orders       │
                  │ Channel      │
                  └──────┬───────┘
                         │
             ┌───────────┼───────────┐
             ▼           ▼           ▼
         Worker 1    Worker 2    Worker 3
             │           │           │
             └───────────┼───────────┘
                         ▼
                Shared Processing
                    Statistics
                         │
                    Mutex Lock
```

## Key Concepts

* **Goroutines** — execute workers concurrently.
* **Channels** — safely distribute orders between goroutines.
* **Worker Pool** — maintains a fixed number of workers for controlled concurrency.
* **WaitGroup** — waits for all workers to complete processing.
* **Mutex** — protects shared state from concurrent access.
* **Package Separation** — separates order and worker responsibilities.

## Project Structure

```text
food-order-worker/
├── internal/
│   ├── order/
│   │   └── order.go
│   └── worker/
│       └── worker.go
├── main.go
└── go.mod
```


## Learning Objectives

This project was built to strengthen practical understanding of **Go concurrency** and prepare for building concurrent backend services using Go.

### Technologies

* Go
* Goroutines
* Channels
* `sync.WaitGroup`
* `sync.Mutex`
