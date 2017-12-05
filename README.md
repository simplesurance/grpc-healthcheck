# grpc-healthcheck

grpc-healthcheck is Golang tool to check if one or more GRPC-Servers are
reachable.
Multiple addresses of GRPC-Servers can be passed as commandline parameter.
The tool tries to establish a connection to each of them.
If connection establishment fails it exits with code 1 immediately.
If a connection could be established to all passed addresses it exits with
code 0.

## Installation
`go get -u github.com/simplesurance/grpc-healthcheck`

## Usage Example
`grpc-healthcheck 192.168.0.1:5001 192.168.0.2:5003`
