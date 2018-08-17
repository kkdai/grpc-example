# GRPC Examples and Survey

This repo include a GRPC example in Go, which target two major functionalities.

- Streaming request and streaming response refer `SayStreamHello`
- Large message size (original one is 4MB, enlarge to 8MB) `MaxMsgSize`

## How to build it

1. Install go and cmake
2. Install proto buffer
    - `brew install protobuf`
3. Make it
    - `make all`

## How to run it

1. Put some big image file (over 5MB) which name `big.jpg`
2. Run server-side `./server`
3. Start another terminal and run client `./client`

## Docker run

### Run GRPC Server in docker

`docker run -it --net=host kkdai/grpc-example /go/bin/server`

### Run GRPC Client in docker

`docker run -it -net=host kkdai/grpc-example /go/bin/client`

# Note:

If you want to build this docker your own, make sure you have a file which name `big.jpg`.

# Learn more detail:

- Refer blog post [here](http://www.evanlin.com/go-grpc-service/).  