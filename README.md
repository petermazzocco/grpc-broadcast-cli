# gRPC Broadcast CLI

A simple gRPC-based client-server application that allows broadcasting messages from a client to a server. Built as part of the [roadmap.sh broadcast server project](https://roadmap.sh/projects/broadcast-server).

## Project Structure

- `main.go` - Entry point with CLI flags for running server or client mode
- `server.go` - gRPC server implementation that receives and logs broadcast messages
- `client.go` - gRPC client implementation that sends messages to the server
- `broadcast.proto` - Protocol buffer definition for the broadcast service
- `broadcast/` - Generated gRPC code from protobuf definitions

## Usage

### Start the Server

```bash
go run . -mode=server
```

The server will start listening on port 9000 and display "Broadcast gRPC started".

### Send a Message from Client

In a separate terminal, run:

```bash
go run . -mode=client -msg="Hello, World!"
```

You can customize the message:

```bash
go run . -mode=client -msg="Your custom broadcast message"
```

### Default Behavior

- Default mode is `server`
- Default message is `message`

## How It Works

1. The server (`server.go`) implements the `SendMessage` RPC method that:
   - Logs the received message
   - Returns a confirmation response

2. The client (`client.go`) connects to the server and:
   - Sends the specified message via gRPC
   - Displays the server's response

3. The main program (`main.go`) uses command-line flags to determine whether to run in server or client mode.

## Requirements

- Go 1.21+
- gRPC dependencies (automatically managed via `go.mod`)

## Building

```bash
go build .
```

Then run the binary:

```bash
# Start server
./grpc-broadcast-cli -mode=server

# Send message
./grpc-broadcast-cli -mode=client -msg="Hello from binary!"
```