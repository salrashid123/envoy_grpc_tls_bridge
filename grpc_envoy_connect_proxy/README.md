## Envoy CONNECT proxy for gRPC


`client --> CONNECT (no TLS) --> Envoy:3128 --> (TLS) --> Envoy:50051 --> (no TLS) --> server:50052`

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

```bash
    export GRPC_GO_LOG_SEVERITY_LEVEL=info
    export GRPC_GO_LOG_VERBOSITY_LEVEL=99

    export https_proxy="localhost:3128" 

    go run src/grpc_client.go
```

envoy_server log

```log
[2021-08-27 15:03:42.630][1975228][debug][http] [source/common/http/conn_manager_impl.cc:896] [C2][S7664381431679857344] request headers complete (end_stream=false):
':authority', 'server.domain.com:50051'
':method', 'CONNECT'
'user-agent', 'grpc-go/1.33.2'

[2021-08-27 15:03:42.630][1975228][debug][router] [source/common/router/router.cc:425] [C2][S7664381431679857344] cluster 'service_bbc' match for URL ''
[2021-08-27 15:03:42.630][1975228][debug][router] [source/common/router/router.cc:582] [C2][S7664381431679857344] router decoding headers:
':authority', 'server.domain.com:50051'
':method', 'CONNECT'
':scheme', 'https'
'user-agent', 'grpc-go/1.33.2'
'x-forwarded-proto', 'http'
'x-request-id', '9930231e-26fd-4634-8fbd-9e6c0b4dc30f'
'x-envoy-expected-rq-timeout-ms', '15000'

[2021-08-27 15:03:42.630][1975228][debug][pool] [source/common/tcp/original_conn_pool.cc:98] creating a new connection
[2021-08-27 15:03:42.630][1975228][debug][pool] [source/common/tcp/original_conn_pool.cc:383] [C3] connecting
[2021-08-27 15:03:42.630][1975228][debug][connection] [source/common/network/connection_impl.cc:836] [C3] connecting to 127.0.0.1:50051
[2021-08-27 15:03:42.630][1975228][debug][connection] [source/common/network/connection_impl.cc:852] [C3] connection in progress
[2021-08-27 15:03:42.630][1975228][debug][pool] [source/common/tcp/original_conn_pool.cc:125] queueing request due to no available connections
[2021-08-27 15:03:42.630][1975228][debug][connection] [source/common/network/connection_impl.cc:651] [C3] connected
[2021-08-27 15:03:42.631][1975228][debug][pool] [source/common/tcp/original_conn_pool.cc:303] [C3] assigning connection
[2021-08-27 15:03:42.631][1975228][debug][router] [source/common/router/upstream_request.cc:354] [C2][S7664381431679857344] pool ready
[2021-08-27 15:03:42.631][1975228][debug][router] [source/common/router/router.cc:1174] [C2][S7664381431679857344] upstream headers complete: end_stream=false
[2021-08-27 15:03:42.631][1975228][debug][http] [source/common/http/conn_manager_impl.cc:1495] [C2][S7664381431679857344] encoding headers via codec (end_stream=false):
':status', '200'
'date', 'Fri, 27 Aug 2021 19:03:42 GMT'
'server', 'envoy'

[2021-08-27 15:03:42.633][1975228][debug][misc] [source/common/network/io_socket_error_impl.cc:30] Unknown error code 104 details Connection reset by peer
```