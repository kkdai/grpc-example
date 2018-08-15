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

# Learning detail:

- Refer blog post here.  (WIP)