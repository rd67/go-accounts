
***Note: Not able to migrate the sql using cli tool - https://github.com/golang-migrate/migrate Because using docker-compose instead of simple sql container. Can use the migrations in the golang code.
***

```migrate create -ext mysql -dir db/migrations -seq initial```


# source .env && clear && migrate -path db/migrations -database "mysql://$MYSQL_USER:$MYSQL_PASSWORD@$MYSQL_HOST:$MYSQL_PORT/$MYSQL_DATABASE" up

# source .env && clear && migrate -path db/migrations -database "mysql://accountUser:accountPassword@127.0.0.1:3307/accounts" up
# mysql://accountUser:accountPassword@127.0.0.1:3307/accounts?query



# docker run -v db/migrations --network accounts-network migrate/migrate -path=db/migrations/ -database "mysql://$MYSQL_USER:$MYSQL_PASSWORD@$MYSQL_HOST/$MYSQL_DATABASE" up 2


docker run -d \
  --name mysql_db \
  -e MYSQL_ROOT_PASSWORD=rootPassword \
  -e MYSQL_DATABASE=accounts \
  -e MYSQL_USER=accountUser \
  -e MYSQL_PASSWORD=accountPassword \
  -p 3307:3306 \
  mysql:8.2
