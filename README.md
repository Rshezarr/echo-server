# Echo Server

This is a simple HTTP server that handles three routes:

-   `/events/`: This route handles Server-Sent Events (SSE) connections from clients. It listens for messages on a channel and sends them to the client via SSE.

-   `/echo`: This route listens for HTTP POST requests from clients and updates a global variable with the value of the w form field in the request. It also starts a goroutine that sends the value of the global variable to the /events/ route via a channel.

-   `/say`: This route listens for HTTP POST requests from clients and updates the global variable with the value of the w form field in the request.

## Usage

To start the server, run the following command:

```
$ go run main.go
```

The server will start listening on `localhost` on port `3000`. Follow the URL
To update the value of the global variable, send an HTTP POST request to the `/say` route with a `w` form field:

```sh
curl -X POST -d "w=Hello, World!" http://localhost:3000/say
```

OR type any message at `w` variable:

```
localhost:3000/echo?w=HelloWorld
```

This will update the global variable and send the new value to the `/events/` route via the channel.
