# Go Project: Concurrent API Requests with Caching and Data Processing

## Description
This project simulates concurrent API requests, handles data processing, and demonstrates caching mechanisms using Go.

## Key Features:
- Concurrent API calls handling using Go routines
- Caching of initial and processed data
- Basic API server with multiple endpoints

## Key Folders

### calculations 
has files for calculations like processor.go (for calculation of data for company and api), initial_data.go (for initial data for each company) and other helper functions

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




