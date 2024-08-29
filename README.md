repo for the go todo list app for the go-academy

## To start the CLI app: 
```
cd cli-app 

go run main.go
```

## To start the API:

```
cd web-app 

go run main.go
```

I made requests using postman. 

### A list of possible requests:

- GET localhost:8080/todo-list - gets a list of all todos
- POST localhost:8080/create - create a todo takes a JSON body in the format: {"description": "hi this is a to do"}
- PUT localhost:8080/update?id={id} - updates the status of the todo that matches the id
- DELETE localhost:8080/remove?id={id} - removes the todo that matches the ID
- GET localhost:8080/todo?id={id} - returns a single todo that matches the ID
