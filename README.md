# Basic User API

## How to get started:
To start the application, copy and paste this into your terminal and hit `return`:

    go run main.go router.go routes.go logger.go handlers.go user.go data.go

Once the application has been started, port `:8080` will be open for your enjoyment. You can hit additional endpoints as well.

### Goals:
We hope to keep the `main.go` file to be as lightweight as possible. Right now we have an instance of our router, which handles all of the requests on the server. We also have the `ListenAndServe` function from the `net/http` package that creates the local server connection.   

### Add additional handlers:
Each endpoint requires a handler. The handler will take care of any logic when a particular endpoint has been hit. If it's a simple `GET` request, the handler will manage what gets returned when the endpoint has been hit. The handler can also manage the logic of a `POST` request as well. You can find the other handlers in the `handlers.go` file.

```go
func MyNewHandler(w http.ResponseWriter, r *http.Request) {
    // Handler code goes in here
}
```
Each handler will require that the same method signature be included. These handle the requests and setting the response headers.

### Add additional routes:
If you'd like to add additional routes to this project, it's fairly simple.

First, you'll need to add the handler in the `handlers.go` file. From there, you can follow the same pattern as you may see with the other handlers in the `routes.go` file. You'll need to give it a name, method (such as `GET` or `POST`), the pattern of the endpoint, and the handler name.
```go
var routes = Routes{
    Route{
    	Name:        "MyNewHandler",
    	Method:      "GET",
    	Pattern:     "/",
    	HandlerFunc: MyNewHandler,
}
```

### Setting up resources
To find out how a resource has been set up, you can see an example within the `user.go` file. This file creates a User struct, which will act as our model for when we get user information from the database. When the user object comes in from the database, we'll need to encode it into one of our User objects.   

### Creating additional resources
The `data.go` file is currently the place where additional logic is kept for creating or updating resources. We only have creation logic in there. We can add additional logic to this file in the future for other use cases.

### Logging:
We have set up a very simple logger. It will display basic information when each endpoint has been hit. You can read the logs in the terminal window where you started the application. You can see it's setup in the `logger.go` file.
