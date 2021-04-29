# Djuno
[![Codecov](https://img.shields.io/codecov/c/github/desmos-labs/djuno)](https://codecov.io/gh/desmos-labs/djuno)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/desmos-labs/djuno/Tests)](https://github.com/desmos-labs/djuno/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/desmos-labs/djuno)](https://goreportcard.com/report/github.com/desmos-labs/djuno)

This project represents the [Juno](https://github.com/desmos-labs/juno) implementation for
the [Desmos blockchain](https://github.com/desmos-labs/desmos).

It extends the custom Juno behavior with custom message handlers for all the Desmos messagejuno. This allows to store
the needed data inside a [PostgreSQL](https://www.postgresql.org/) database on top of
which [GraphQL](https://graphql.org/) APIs can then be created using [Hasura](https://hasura.io/)

## Installation
To install the binary simply run `make install`.

**Note**: Requires [Go 1.13+](https://golang.org/dl/)

## Database
Before running the parser, you need to:

1. Create a [PostgreSQL](https://www.postgresql.org/) database.
2. Run the SQL queries you find inside the `*.sql` files in the [schema folder](schema) inside such database to create
   all the necessary tables.

## Initialize the config
First of all, initialize DJuno configuration file by running

```
djuno init
```

This will create a folder inside your home folder. In Linux this will be located at `~/.djuno`. Inside such folder you
will find a `config.toml` file. Make sure you edit that file to have all the values that it needs.

### Configuration
The configuration must be a TOML file containing the following fields:

```toml
[cosmos]
prefix = "desmos"
modules = [
    "bank",
    "notifications",
    "posts",
    "profiles",
    "relationships",
    "reports"
]

[rpc]
address = "<rpc-ip/host>:<rpc-port>"

[grpc]
address = "<client-ip/host>:<client-port>"

[api]
address = "<client-ip/host>:<client-port>"

[database]
type = "postgresql"

[database.config]
host = "<db-host>"
port = 5432
name = "<db-name>"
user = "<db-user>"
password = "<db-password>"
```

<details> 

<summary>Example of a configuration to parse the chain state from a local full-node:</summary>

```toml
[cosmos]
prefix = "desmos"
modules = [
    "bank",
    "notifications",
    "posts",
    "profiles",
    "relationships",
    "reports"
]

[rpc]
address = "http://localhost:26657"

[grpc]
address = "http://localhost:9090"

[api]
address = "http://localhost:1317"

[database]
type = "postgresql"

[database.config]
host = "localhost"
port = 5432
user = "user"
password = "password"
name = "djuno"
```

</details>

## Running the parser
To parse the chain state, you need to use the following command:

```shell
djuno parse 
```