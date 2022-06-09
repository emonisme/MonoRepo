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
## Testing

After you run rest server, you can test it localy by hitting `localhost:8080` or you can use ngrok to test it use public internet. Run this command if you want to test it using ngrok
```
ngrok http 8080
```
After you run it, you will see this output
```
ngrok   (Ctrl+C to quit)

Hello World! https://ngrok.com/next-generation
Session Status                online
Session Expires               1 hour, 59 minutes
Update                        update available (version 3.0.4, Ctrl-U to update)
Terms of Service              https://ngrok.com/tos
Version                       3.0.4
Region                        Asia Pacific (ap)
Latency                       18ms
Web Interface                 http://127.0.0.1:4040
Forwarding                    https://8b1a-182-253-124-82.ap.ngrok.io -> http://localhost:8080

Connections                   ttl     opn     rt1     rt5     p50     p90
                              0       0       0.00    0.00    0.00    0.00
```
now you can test using public internet by accessing url in forwarding. in this case https://8b1a-182-253-124-82.ap.ngrok.io
