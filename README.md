# wallet
An implementation of a wallet in Go and GraphQL.

Example operations: 
* add credit to the wallet
* withdrawal funds from the wallet
* check balance

## Commands

### Database
`sql-migrate new YOUR_MIGRATION_NAME`

Migrations are automatically applied on application startup.

### GraphQL

`go run github.com/99designs/gqlgen generate`
