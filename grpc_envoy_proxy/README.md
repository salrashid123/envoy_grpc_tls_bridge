## Envoy TLS proxy for gRPC


`client --> (no TLS) --> Envoy:8080 --> (TLS) --> Envoy:50051 --> (no TLS) --> server:50052`

get envoy binary

```
    docker cp `docker create envoyproxy/envoy-dev:latest`:/usr/local/bin/envoy .
```

1. Start gRPC Server

```
   go run src/grpc_server.go 
```

2. Start Envoy Server Proxy

```
    ./envoy -c envoy_server.yaml --base-id 0 -l debug
```

3. Start Envoy Client Proxy

```
   ./envoy  -c envoy_client.yaml --base-id 1 -l debug
```

3. Start gRPC Client

```
    export GRPC_GO_LOG_SEVERITY_LEVEL=info
    export GRPC_GO_LOG_VERBOSITY_LEVEL=99

    go run src/grpc_client.go
```

