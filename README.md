# Backend API

A gRPC based api.

For setup information see [SERVER_SETUP](docs/SERVER_SETUP.md)

## Building

```shell
GOBIN="$PWD" make
```

## Running

* Setup MySQL and import `db/schema.sql`.
* Set all environment variables in `.env.example`
* Run `server`

## Tests

Client tests are in `src/clienttests/test.go`

Based on the test, pass the appropriate subcommand and arguments.