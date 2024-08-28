# Concurrency Problems in Golang

This repository contains implementations of classic concurrency problems using Golang. Each problem is solved using Go routines, channels, and other concurrency primitives provided by the Go programming language.

## Table of Contents

- [Dining Philosophers Problem](#dining-philosophers-problem)
- [Producer-Consumer Problem](#producer-consumer-problem)
- [Race Condition, Mutexes, Channels](#race-condition-mutexes-channels)
- [Sleeping Barber Problem](#sleeping-barber-problem)

## Problem Descriptions

### Dining Philosophers Problem

The Dining Philosophers problem is a classic synchronization problem involving philosophers sitting at a table with forks between them. Each philosopher needs both forks to eat and must alternate between thinking and eating. This problem explores deadlock, starvation, and resource allocation techniques.

- **Folder:** `DiningPhilosophers-problem`
- **Implementation:** Includes Go routines to manage the behavior of philosophers and forks.

### Producer-Consumer Problem

The Producer-Consumer problem is a standard concurrency problem where multiple producers generate data and place it into a buffer while multiple consumers take data from the buffer. This problem demonstrates the use of semaphores or mutexes to handle synchronization.

- **Folder:** `Producer-Consumer-Problem`
- **Implementation:** Go routines and channels are used to synchronize the producer and consumer processes.

### Race Condition, Mutexes, Channels

This section deals with race conditions that occur when multiple Go routines access shared data concurrently. The folder contains examples of how to use mutexes and channels to avoid race conditions and ensure data consistency.

- **Folder:** `Racecondition-Mutexes-Channels`
- **Implementation:** Demonstrates various ways to handle race conditions using Go's synchronization primitives.

### Sleeping Barber Problem

The Sleeping Barber problem is a classic synchronization problem that involves a barber, who sleeps when no customers are present and wakes up when a customer arrives. This problem is similar to the producer-consumer problem and focuses on managing limited resources.

- **Folder:** `Sleeping-Barber-Problem`
- **Implementation:** Simulates the behavior of a barber shop using Go routines and channels to manage customer arrivals and barber activity.

## Go Modules

- **File:** `go.mod`
- **Description:** The `go.mod` file specifies the dependencies required by the project.

## Main Files

- **File:** `main.go`
- **Description:** The main entry point for the Go application. This file typically contains the execution logic for the problems implemented in this repository.

- **File:** `main_test.go`
- **Description:** Contains test cases for the implementation in `main.go`. This helps to ensure the correctness and efficiency of the code.

## Getting Started

### Prerequisites

- Go 1.16+ installed on your system.

### Running the Code

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/your-repo-name.git
    cd your-repo-name
    ```

2. Navigate to the desired problem folder and run the Go file:

    ```bash
    cd DiningPhilosophers-problem
    go run main.go
    ```

3. To run tests:

    ```bash
    go test
    ```

