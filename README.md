# My Go SQL

Go x DB simple project

## Simple run with go command
To run the application with console based mode
```go
go run myfirstgosql/main --mode cli
```

To run the application with http mode
```go
go run myfirstgosql/main --mode http
```

If you want to do database simple migration. You can set the migration json file in environment variable
```go
go run myfirstgosql/main --mode migration-up
go run myfirstgosql/main --mode migration-down
```

If you want to select a specific environment
```go
go run myfirstgosql/main --env dev --mode cli
```
Please make sure you have dev.env file, if the file is not exist
it will try to read operating system environment variable,
and if the environment variable does not have the required info, then
the app default value will be taken.

The default is try to read .env file, if you do not use --env option.
If the .env file is not exist, the same behaviour like statement above.

In the production environment, usually we provide the config in operating system
environment variables.

To view Some help
```go
go run myfirstgosql/main --help
```
## Docker build
```
docker build . -t my-go-sql
```

## Docker images
Check the build result, you should have my-go-sql in the list
```
docker images 
```

#Test run
```
docker run -it -e DBUSER=root -e DBPASSWORD=P@ssw0rd -e DBHOST=host.docker.internal -e DBPORT=3306 -e DBSCHEMA=enigma -e DBENGINE=mysql -e APPMODE=cli --name mygosql --rm my-go-sql
```
or using docker compose
(the environment variables information is provided in .env files)
```
docker-compose up --build
docker-compose down
```
You should get the result on console
