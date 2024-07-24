# Shopping-Cart

## Introduction

This project is a simple shopping-cart system written in **golang** developed as the Internet Engineering course midterm
project.

## Technologies used

- [Golang](https://golang.org/), Programming language.
- [Echo](https://echo.labstack.com/), HTTP web framework.
- [Gorm](https://gorm.io/), ORM library for Golang.
- [PostgreSQL](https://www.postgresql.org/), database management system.
- [Docker](https://www.docker.com/), Containerization platform.

Features
--------

The implemented service offers the following API endpoints:

- `GET /basket/`: Retrieves a list of baskets
- `POST /basket/`: Creates a new basket
- `PATCH /basket/<id>`: Updates the specified basket
- `GET /basket/<id>`: Retrieves a specific basket
- `DELETE /basket/<id>`: Deletes the specified basket

Implementation Details
----------------------

### Basket structure

- `id`
- `user_id`: Identifier for the user who owns the basket
- `created_at`: Time when the basket was created
- `updated_at`: Time when the basket was last updated
- `data`: Variable length data (maximum size: 2048 bytes)
- `state`: Status of the basket (COMPLETED or PENDING)

### User structure

- `id`
- `created_at`: Time when the user account was created
- `updated_at`: Time when the user account was last updated
- `username`: Unique identifier for the user
- `password`: User's password

### Authentication

Each user can perform CRUD operations only on their respective baskets after authentication.
Token-based(JWT) authentication is implemented to authenticate users.
Users receive a token upon successful login, which they must use in API requests to access their baskets.
CRUD operations on baskets are restricted to the authenticated user.

## Run project

In order to run the code you need to create your own .env file for your custom configurations and after that
you only need to enter the following command in the root directory of your code.

```shell
$ docker-compose up -d
```
