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

---

DNS design

Resolver interface

LookupIP(ctx, host)

Implementations

- SystemResolver (current)
- CacheResolver (next)
- DoHResolver (future)
- DoTResolver (future)

Pipeline

Request
    ↓
Classifier
    ↓
Resolver
    ↓
Policy
    ↓
Dispatcher

Resolver never decides routing.

Resolver only resolves names.

---

# Development Rules

These rules are part of the project.

## Goal

Primary goal:

Get unrestricted Internet on phones as soon as possible.

Architecture exists to support this goal, not vice versa.

---

## Working Rules

1. Every 2–3 assistant responses update this document.

2. Before every major component (DNS, SOCKS, API, TUN, etc.) update this document.

3. After finishing a major component update this document.

4. Never redesign working code without a real reason.

5. Green "go build" means the current architecture is stable.

---

## Assistant Rules

Assistant MUST provide ready-to-run shell commands.

Preferred:

- cat > file <<'EOF'
- python scripts
- go fmt
- go build
- git commit

Avoid:

- nano
- vim
- manual editing
- "find this line"
- "replace this"
- partial snippets requiring manual merge

Every change should be executable by copy-paste.

---

## Current Priority

1. Working Internet on Android phones.
2. Stable Runtime.
3. DNS.
4. Rules.
5. ACL.
6. SOCKS/TUN/API.

Do not optimize prematurely.

---

# DNS Integration (current task)

Status:
- Resolver interface: DONE
- SystemResolver: DONE
- Core integration: TODO
- Dispatcher integration: TODO

Current policy:

DNS is NOT responsible for routing.

DNS only resolves names.

Request flow:

Client
    ↓
Classifier
    ↓
Resolver
    ↓
Policy
    ↓
Dispatcher
    ↓
Forwarder

Resolver may later support:

- cache
- hosts
- DoH
- DoT

Current implementation target:

SystemResolver only.

No cache.
No DoH.
No DoT.
