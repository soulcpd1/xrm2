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

---

# Current Project Status

Last verified:

- Runtime: COMPLETE
- Hysteria TCP: COMPLETE
- Hysteria UDP: COMPLETE
- DNS abstraction: COMPLETE
- Runtime architecture: STABLE

Mobile clients:

- iPhone: CONNECTS
- Android: CONNECTS

Traffic:

- Internet access: WORKING
- Bypass blocked resources: NOT WORKING

This means:

Transport layer is considered COMPLETE.

The remaining issue is traffic routing / bypass logic.

---

# Current Goal

Primary objective:

Reliable unrestricted Internet on iPhone and Android.

Architecture is finished enough.
Do not redesign Runtime unless a real bug is found.

---

# Next Investigation

Determine why blocked traffic is not bypassed.

Verify, in order:

1. Classifier result.
2. Policy decision.
3. Dispatcher selected route.
4. Forwarder selected driver.
5. Actual outbound used.
6. Whether Nebula is ever selected.
7. Whether Direct is always selected.

Do NOT add new subsystems until this is understood.

---

# Working Agreement

Continue using these rules:

- Update this document every 2–3 assistant responses.
- Prefer working software over architectural refactoring.
- Every code change must be copy-paste executable.
- No manual editing instructions.
- One component at a time.
- Green "go build" means do not redesign.


---

# Investigation Strategy

Current hypothesis:

The bypass failure is most likely inside the runtime decision chain.

Investigation order:

Request
  ↓
Classifier
  ↓
Policy
  ↓
Dispatcher
  ↓
Forwarder
  ↓
Outbound Driver

At each stage we will verify actual runtime values instead of assuming they are correct.

No new features until the decision chain is verified.

---

# New Priority

Transport layer is COMPLETE.

Next implementation target:

Real Nebula outbound.

Current situation:

direct:
    net.Dialer

nebula:
    net.Dialer   <-- temporary placeholder

This means both drivers currently behave almost identically.

Next objective:

Implement a real Nebula transport so that routing decisions have an actual effect.

After that:

Classifier
    ↓
Policy
    ↓
Dispatcher
    ↓
Forwarder
    ↓
direct | nebula

will produce different network paths.


---

# Nebula Integration Rules

XRM MUST coexist with the existing Xray + Nebula deployment.

Requirements:

- Never modify Nebula configuration.
- Never modify Linux routing tables.
- Never restart nebula.service.
- Never assume exclusive ownership of nebula1.

Nebula is an existing transport.

XRM only chooses whether a connection should use:

- Direct
- Nebula

The Nebula driver is only responsible for creating sockets that use the existing nebula1 interface.


---

# Nebula Driver Design

The Nebula driver is independent of Xray.

Xray remains a separate application using the same Nebula network.

XRM will not reuse Xray routing logic.

Instead, the Nebula outbound driver will create connections through the existing `nebula1` interface while leaving Xray and Nebula configuration untouched.


---

# Architecture Decision (2026-07-04)

The project goal has changed from "replace Xray" to building an independent transport.

Current situation:

- VLESS + Reality is already being blocked by mobile operators.
- Amnezia works but requires initial connectivity.
- Nebula between Home and VPS is stable and permanent.

Therefore XRM will use Nebula as the transport between nodes.

Future architecture:

Phone
    ↓
Hysteria
    ↓
XRM Server
    ↓
Nebula
    ↓
XRM Agent (VPS)
    ↓
Internet

The VPS agent is responsible only for:
- accepting streams;
- creating outbound TCP/UDP connections;
- relaying packets.

No VLESS/Reality dependency is planned.

Next milestone:

Implement XRM Agent running on the VPS and define the internal transport protocol between Server and Agent.


---

## Internal XRM Protocol

Current protocol (v1):

Client ----> Agent

uint16 network_length
network bytes

uint16 address_length
address bytes

After successful request exchange, the connection becomes a raw bidirectional stream.

Future versions may extend the header, but v1 must remain stable.


---

## XRM Agent

Status:

✔ accepts TCP connections

✔ reads internal XRM protocol

✔ opens outbound connection

✔ bridges both streams

Next:

Server must replace direct Dial() with a connection to the Agent over Nebula.


---

## Routing Policy (mandatory)

Traffic MUST NOT always pass through Italy.

Routing policy:

- RU domains -> Direct
- Private networks -> Direct
- LAN -> Direct
- Everything else -> Nebula Agent -> VPS -> Internet

Dispatcher is the single place that decides this.

The Agent MUST NEVER decide routing policy.
The Agent only forwards traffic to the Internet.


---

## Agent Driver

Status:

✔ New outbound driver: agent

Current behavior:
- opens TCP connection to XRM Agent over Nebula

Next commit:
- send proto.Request
- Agent performs net.Dial()
- stream becomes transparent


---

## Agent Integration Status

Completed:

- XRM Agent skeleton
- Internal transport protocol (v1)
- Agent outbound driver
- TCP stream forwarding inside Agent

Current limitation:

The Agent driver still connects only to the Agent itself.
The destination request (proto.Request) is not yet sent.

The routing policy remains:

RU / Private / LAN -> Direct

Everything else -> Agent (after transport integration).

