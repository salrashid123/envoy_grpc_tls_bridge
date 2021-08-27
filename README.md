
## gRPC TLS Tunnel proxy using CONNECT

Simple golang gRPC client/server where the connectivity uses TLS provided by envoy:

There are two modes described in this repo:

1. `client --> (no TLS) --> Envoy --> (TLS) --> Envoy --> (no TLS) --> server`

2. `client --> CONNECT (no TLS) --> Envoy --> (TLS) --> Envoy:051 --> (no TLS) --> server`


In the first, its your regular TLS offloading to envoy where just the bit between envoy->envoy is using TLS.  THis is essentially the sidecare model you see with Istio and many other systems.

The second one is much more rare and is essentially a demonstration of envoys' [CONNECT proxy support](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/upgrades#connect-support).

In thi mode, the envoy client basically acts as a TLS proxy but uses the CONNECT settings to proxy thorugh.