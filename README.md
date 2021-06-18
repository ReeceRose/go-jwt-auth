# JWT Authentication via Go REST API

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT/)
[![Build](https://github.com/ReeceRose/go-jwt-auth/actions/workflows/go.yml/badge.svg)](https://github.com/ReeceRose/go-jwt-auth/actions/workflows/go.yml/)
[![GoReportCard](https://goreportcard.com/badge/github.com/nanomsg/mangos)](https://goreportcard.com/report/github.com/ReeceRose/go-jwt-auth/)

Simple JWT authentication REST API built with Go & Fiber

# Overview

This project is designed to be used as a template to quickly bootstrap any application that required authentication. Simply add any existing routes/users/controllers and call the API from the endpoint of your choice. This project can be expanded upon to also include things like: email verification and password resets.

## API routes

Currently the project contains two routes:

- `/api/auth/login`
  - Used to login an existing user.
  - Returns a JWT if login is successful.
  - Can easily be modified to use/send cookies
- `/api/auth/register`
  - Used to register a new user.
  - Returns the user if registration is successful.

## Dependencies

This project has 6 direct dependencies:

- [Fiber](https://github.com/gofiber/fiber)
  - Express inspired web framework written in Go
- [Fiber JWT](https://github.com/gofiber/jwt)
  - JWT middleware for Fiber
- [Gorm](https://github.com/go-gorm/gorm)
  - ORM library for Golang
- [Gorm SQLite Driver](https://github.com/go-gorm/sqlite)
  - SQLite driver for Gorm
- [Crypto](https://github.com/golang/crypto)
  - Go supplementary cryptography libraries
- [JWT Go](https://github.com/form3tech-oss/jwt-go)
  - Golang implementation of JSON Web Tokens (JWT)

# Developer Setup

## Clone repository

In the desired folder, run:

```bash
git clone git@github.com:ReeceRose/go-jwt-auth.git .
```

## Vendor dependencies

In the root project folder, run:

```bash
go mod vendor
```

## Install dependencies

In the root project folder, run:

```bash
go get ./...
```

## Run API

```bash
go run .
```

# Production Use

## Build API

To get a production ready API, run:

```bash
go build .
```

_Depending on your system, this will either generate **go-jwt-auth** or **go-jwt-auth.exe**_

## Run production API

After copying the generated file from the step above to your desired server/location, run:

```bash
./go-jwt-auth # Note, this file name and exact syntax may differ depending on your system.
```
