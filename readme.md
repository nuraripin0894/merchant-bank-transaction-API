# Merchant-Bank Transaction API

## Overview

This is a Go-based API for handling transactions between customers and merchants. The API supports user login-logout, payment processing, and logging of user actions to a history file.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Setup](#setup)
- [Environment Variables](#environment-variables)
- [Endpoints](#endpoints)

## Features

- **Login**: Authenticate customers and manage sessions.
- **Payment**: Allow authenticated customers to make payments to registered merchants.
- **Logout**: Terminate user sessions safely.
- **History Logging**: Record all user actions (login, payment, logout) to a JSON file for auditing and tracking.

## Technologies Used

- Go (Golang)
- PostgreSQL
- JSON for logging history
- [Postman](https://www.postman.com/) for API testing

## Setup

### Prerequisites

- Go installed on your machine (version 1.16 or later).
- PostgreSQL installed and running.
- A working PostgreSQL database.

### Usage

1. Clone this repository:

   ```bash
   git clone https://github.com/nuraripin0894/merchant-bank-transaction-API
   cd merchant-bank-transaction-API
   ```

2. Install dependencies:

   ```bash
   go get
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

4. Use Postman to interact with the API.

## Environment Variables

```
To run this project, you need to set the following environment variables in a .env file in the root directory of your project:
```

- **DB_USER =** `your_db_user`
- **DB_PASSWORD =** `your_db_password`
- **DB_NAME =** `your_db_name`

## Endpoints

### 1. Login

- **Endpoint:** `/login`
- **Method:** `POST`
- **Request Body:**

  ```
  x-www-form-urlencoded

  email: <your_email> (e.g., test@example.com)
  password: <your_password> (e.g., password123)
  ```

### 2. Payment

- **Endpoint:** `/payment`
- **Method:** `POST`
- **Request Body:**

  ```
  x-www-form-urlencoded

  merchant_id: 1 (for example)
  amount: 100.50
  ```

### 3. Logout

- **Endpoint:** `/logout`
- **Method:** `POST`

### 4. History

- **Endpoint:** `/history`
- **Method:** `GET`
