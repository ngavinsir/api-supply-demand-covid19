## API for Supply/demand COVID 19

You will need a postgres database to run the application. If you have docker installed, then you can create a database by running

```
docker-compose up db
```

Then, if you have go installed, you can build and run the server by running

```
go run main.go
```

Or if you don't have go installed, you can launch both the database and the server with docker compose

```
docker-compose up --build
```

### Testing

You will need a postgres database to run some of the test. If you have docker installed, then you can create the test database by running

```
docker-compose up test-db
```

Then, run the test command, for example

```
go test ./model -v
```
