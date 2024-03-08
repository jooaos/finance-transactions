# About

The project consists of an API that simulates some financial transactions

# How run it?

### Dependecies

For run this project you just need:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/compose-file/)

And if you want more facility, you can install [Make](https://www.geeksforgeeks.org/how-to-install-make-on-ubuntu/).

### Step by step

> All commands here are using make, but if you don't want use it, you can access Makefile, copy command and paste in your terminal

If it's the first time that you'll run this project, you first need
```
git clone git@github.com:jooaos/pismo.git && cd pismo
./scripts/first-install.sh
```

After it, you can up the application 

Run the follow command and check if your app is up
```
make up
docker container ps | grep "pismo"
```

It's necessary to appear two containes: `pismo_app` and `pismo_db`, if it's showing, you can go to next steps

Now it's time to migrate the database, and for it use:
```
make migration-up
```
Now you can use the application using the endpoint `http://localhost:8080` as wanted. Here in project we make available two docs about the endpoints, that are:
- [Postman](https://github.com/jooaos/pismo/tree/main/docs/postman)
- [Swagger](https://github.com/jooaos/pismo/tree/main/docs/swagger)

> About **Swagger**, you can up a swagger container using the command `make up-swagger` and access `http://localhost` to check the docs, if you want down you can use `make down-swagger`

If you want down the application, you can use the command:
```
make down-all
```

# Tests
In this application we have two kind of tests:
- Integration tests and
- Unit tests

### Integration tests
Before create/test the integrations tests, it's necessary set up the environment, and for it you just need to use the command:
```
make integration-test-migrate-up
```
After it, you can run tests using the command:
```
make integration-test-run
```

### Unit tests
For unti tests is more simple, if you want test you can use the command:
```
make unit-test
```

# Migrations
List of commands about migrations
| Command                   | Action                                            |
| -------                   | ------                                            |
| `make migration-up`       | Run the migrations that database needs            |
| `make migration-down`     | Rollback all migrations                           |
| `make migration-create`   | Create migration, and you can pass the `NAME=`    |


