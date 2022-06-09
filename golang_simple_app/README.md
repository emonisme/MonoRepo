# Golang Simple App

## Summary
golang_simple_app is a book management backend microservice. It handles CRUD book object and save it to database

## Setup
- Run depedency
```
make docker up
```
- Download database migration tools
```
make tool-migrate
```
- Run migration. See makefile for custom database env config. Also see more argument on [golang-migrate](https://github.com/golang-migrate/migrate)
```
MIGRATE_ARGS=up make migrate
```
- Run golang main on cmd
```
go run cmd/rest/main.go
```
