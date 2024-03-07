*** Go lang Microservice Project Test***

Migrations
https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md


```migrate create -ext sql -dir db/migrations -seq 1```

----------------

SQLC
https://docs.sqlc.dev/en/latest/

```sqlc generate```



-----------------
docker run --name accounts-go --network accounts_network -p 3000:3000 -e DB_SOURCE="postgres://accountUser:accountPassword@accounts-postgres:5432/accounts?sslmode=disable" accounts:latest