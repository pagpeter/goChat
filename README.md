# goChat
A simple chat between multiple clients over sockets in golang. It makes heavy use of structs and channels.

## How to use?

Start the server first. For developement, using `go run Server/*.go` is recommended.

After the server is up and running, start as many clients as you want. To do that, do `go run Client/*.go <name>`. Replace `<name>` with your nickname. 

Once you are connected to the server, you can just type in the terminal. When you press enter, the message gets send to the Server, and then to every other connected client.

## Why?

I wanted to learn more about networking in Go.
