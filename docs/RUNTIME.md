# XRM Runtime

Last updated after commit:
1b83876 (feat: implement Hysteria UDP adapter)

---

## Runtime Pipeline

Inbound
    ↓
common.Request
    ↓
Classifier
    ↓
Policy
    ↓
Dispatcher
    ↓
RouteManager
    ↓
Forwarder
    ↓
Outbound Driver
    ↓
net.Conn / net.PacketConn

StreamBridge is independent and is responsible ONLY for bidirectional stream forwarding.

---

## Components

### Classifier
Mutates common.Request.
Produces Category.

### Policy
Maps Category -> RouteName.

### Dispatcher
Uses Policy + RouteManager.
Mutates req.Route.

### RouteManager
Stores Route definitions.

### Forwarder
Uses Outbound.Manager.
Creates TCP and UDP connections.

### StreamBridge
Bridges TCP streams.
Contains NO routing logic.

### Outbound Drivers

Current:

- direct
- nebula

Driver API:

- Dial()
- ListenPacket()

---

## Hysteria

Implemented:

- TCP
- UDP adapter
- Runtime integration

Not implemented yet:

- Full UDP relay
- DNS
- ACL
- RuleSets
- SOCKS5
- HTTP CONNECT

---

## Current Architecture Rules

1. StreamBridge never resolves routes.
2. Dispatcher never opens sockets.
3. Forwarder always opens sockets.
4. Drivers never know about Policy.
5. common.Request is the single routing object.

---

## Next task

Implement full UDP relay for Hysteria using the new UDP adapter.

---

# Next milestone

DNS subsystem

Goal:

Request
    ↓
Classifier
    ↓
DNS Resolver
    ↓
Policy
    ↓
Dispatcher

Resolver responsibilities:

- DNS cache
- static hosts
- future DoH/DoT
- domain -> IP
- preserve original hostname for routing

Current status:

- Runtime: COMPLETE
- Hysteria TCP: COMPLETE
- Hysteria UDP: COMPLETE
- Next: DNS
