# wallet
An implementation of a wallet in Go and GraphQL.

Example operations: 
* add credit to the wallet
* withdrawal funds from the wallet
* check balance

## Commands

### Running the server

First start the db:
```shell
docker compose up -d
```

Then start the go server:
```shell
go run server.go
```

### Closing the server
```
control + c
```

```shell
docker compose down
```

### Database Migrations
```shell
sql-migrate new YOUR_MIGRATION_NAME
```

```shell
sqlc generate
```

Migrations are automatically applied on application startup.



### GraphQL

```shell
go run github.com/99designs/gqlgen generate
```

#### Sample Requests:
```graphql
mutation CreateWallet {
  create(startingBalance:50){
    id
    balance
  }
}

mutation Credit {
  credit(walletId:"c3d7ab57-4767-4a90-98dc-97c5b3fcf70c",amount:100) {
    id
    balance
  }
}

mutation Withdraw {
  withdraw(walletId: "c3d7ab57-4767-4a90-98dc-97c5b3fcf70c",amount:55) {
    id
    balance
  }
}

query GetWallet {
  wallet(id: "c3d7ab57-4767-4a90-98dc-97c5b3fcf70c") {
    id
    balance
  }
}
```

## Sources
* https://github.com/rubenv/sql-migrate
* https://gqlgen.com/getting-started/
* https://graphql.org/learn/
* https://go.dev/doc/tutorial/getting-started
