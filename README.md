# My Go SQL

Go x DB simple project

## Simple run with go command
To run the application with console based mode
```go
go run myfirstgosql/main c
```
If you want to do database simple migration
```go
go run myfirstgosql/main d
```
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
docker run -it -e dbuser=root -e dbpassword=P@ssw0rd -e dbhost=host.docker.internal -e dbport=3306 -e dbschema=enigma -e dbengine=mysql --name mygosql --rm my-go-sql
```
or using docker compose
(the environment variables information is provided in .env files)
```
docker-compose up --build
docker-compose down
```
You should get the result on console
