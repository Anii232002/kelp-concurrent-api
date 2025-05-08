# Go Project: Concurrent API Requests with Caching and Data Processing

## Description
This project simulates concurrent API requests, handles data processing, and demonstrates caching mechanisms using Go.

## Key Features:
- Concurrent API calls handling using Go routines
- Caching of initial and processed data
- Basic API server with multiple endpoints

## Flow diagram

![image](https://github.com/user-attachments/assets/91ebc996-5ddc-4e4d-b9cf-8543d0446ea0)

## Key Folders

### calculations 
has files for calculations like 
#### processor.go 
Main Calculation file for the endpoints
#### initial_data.go 
Calculates initial data for each company
#### Other helper functions
- cache_utils.go -> stores the cache for both initial data and processor (reusable, takes the cache map as the input for storing cache for both
inital data (company - key) and processor data (company:api will be the key)

- concurrency_util.go -> Used for waiting requests while same request is calculation the initial data or the processor data

- data-type.go -> contains the data strutures used across the calculations folder

### handlers
Consist of Handler functions for each endpoint

### main.go 
starting point of the application

### test.go 
Testing the main server with some api calls as per the given agenda of the challenge

## Prerequisites

Before running the server or test script, make sure you have the following installed:

- Go
- Git

## Getting Started

Follow the steps below to get the project running locally.

### Step 1: Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/Anii232002/kelp-concurrent-api.git
cd kelp-concurrent-api
```

### Step 2: Run the server

For Running the project

```bash
go run main.go
```

### Step 3: test against http calls

```bash
cd test
go run test.go
```




