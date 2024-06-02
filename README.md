# Disbursement App

# Overview

Topic: The user has a balance in the application wallet and the balance wants to be disbursed.  
● Write code in Golang  
● Write API (only 1 endpoint) for disbursement case only  
● User data and balances can be stored as hard coded or database

## Features

## endpoint list:

- disburse (user_id [int], amount [float])
    - case:
        - disbursement success [200]
        - insufficient balance [400]
        - invalid request [400]
        - any internal error [500]

for detailed request & response please check `Sample Request & Response` section.

## database:

- sqlite
    - data model & initiated data:
        - `users` table:

| id | balance  |
|:--:|:---------|
| 1  | 10000000 |
| 2  | 20000000 |
| 3  | 25000000 |

for table schema check here: https://github.com/rizanw/disburse-app/blob/main/internal/repo/db/module/schema.go

# Sample Request & Response

## success sample

### db condition

| id | balance  |
|:--:|:---------|
| 1  | 10000000 |
| 2  | 20000000 |
| 3  | 25000000 |

### request

```shell
curl --location '0000:3000/disburse' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 2,
    "amount": 5000000
}'
```

### response

```json
{
  "user_id": 2,
  "request_amount": 5000000,
  "balance": 15000000,
  "status": "success"
}
```

## Insufficient Balance sample

### db condition:

| id | balance  |
|:--:|:---------|
| 1  | 10000000 |
| 2  | 20000000 |
| 3  | 25000000 |

### request

```shell
curl --location '0000:3000/disburse' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 1,
    "amount": 20000000
}'
```

### response

```json
{
  "user_id": 1,
  "request_amount": 20000000,
  "balance": 10000000,
  "status": "insufficient balance"
}
```

# Local Development

## Prerequisites

Make sure you have installed all the following prerequisites on your development machine:

* go version : [1.19](https://golang.org/dl/)

## Local Run Guides:

To clone this repo:

```bash
git clone https://github.com/rizanw/disburse-app.git 
```

To build and start the apps:

- build the binaries:

```bash 
make build
```

- start the app:

```bash 
make run
```

## Unit Test

To run unit test

```bash
make test
```

# Project Structure

- `bin/` is directory for compiled binary
- `cmd/` is the main program directory
- `files/` contains app files (including db & config)
    - `file/db` contains sqlite db directory
    - `file/etc/disburse-app` contains app config files
- `internal/` contains the whole logic of the app
    - `internal/config` is the config of the app, has relation to files directory
    - `internal/handler` is application logic interface between this app with client
    - `internal/model` is model business design
    - `internal/repo` is the repositories to fetch/store data of this app
    - `internal/usecase` is main business logic
- `go.mod` the golang dependencies list