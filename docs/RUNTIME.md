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


---

## Current Runtime Status

Completed:

✔ Hysteria integration

✔ Classifier -> Policy -> Dispatcher pipeline

✔ Driver abstraction

✔ StreamBridge

✔ XRM transport protocol v1

✔ XRM Agent

✔ Agent outbound driver

✔ Request header transmission

Current state:

Server and Agent can communicate using the XRM protocol.

Next milestone:

Run Agent as a standalone service on the VPS and redirect the "italy" route from the Nebula driver to the Agent driver.


---

## XRM Agent executable

Executable:

cmd/xrm-agent

Default listen:

0.0.0.0:7000

Deployment target:

VPS (Italy)

The Agent is completely stateless.


---

## Security Rules

The XRM Agent MUST NEVER expose a public listener.

Forbidden:

- 0.0.0.0:7000
- public Internet access

Required:

- listen only on Nebula interface
- default: 192.168.100.1:7000
- configurable address
- all Agent traffic goes through Nebula

Reason:

The Agent is an internal XRM component, not a public proxy.


---

## Current Development State

Infrastructure status:

✔ Runtime pipeline completed

✔ Driver abstraction completed

✔ Hysteria integration completed

✔ XRM protocol v1 completed

✔ Agent skeleton completed

✔ Agent transport completed

✔ Nebula transport selected

Architecture frozen.

No more infrastructure refactoring unless it fixes a bug.

Next phase:

Run the Agent on the VPS.

Verify:

Home XRM
    ↓
Nebula
    ↓
Agent
    ↓
Internet

After successful verification:

Focus moves to:

- GeoIP
- GeoSite
- DNS decisions
- application-aware routing
- censorship bypass


---

## Next Milestone

Goal:

Run the first XRM Agent on the VPS.

Verification sequence:

1. Agent starts on Nebula address.
2. Home XRM connects to Agent.
3. Agent receives proto.Request.
4. Agent performs net.Dial().
5. End-to-end TCP connection succeeds.

Only after this milestone is complete may the "italy" route be switched from the Nebula driver to the Agent driver.


---

## Agent Status

Implemented:

✔ TCP listener

✔ proto.ReadRequest()

✔ remote net.Dial()

✔ full duplex io.Copy()

The Agent implementation is considered feature-complete for TCP v1.

The next task is deployment and end-to-end verification.


---

## Agent Listen Policy

The default XRM Agent listen address is:

192.168.100.1:7000

The Agent MUST NOT bind to:

- 0.0.0.0
- :7000

Reason:

The Agent is reachable only through the Nebula overlay network and must never expose a public TCP listener.


---

## Route vs Driver

IMPORTANT

Route is a routing decision.

Examples:

- home
- italy
- backup

Driver is a transport implementation.

Examples:

- direct
- nebula
- agent

Routes MUST NOT depend directly on transport implementations.

Future execution chain:

Classifier
    ↓
Policy
    ↓
Route
    ↓
Route Executor
    ↓
Driver
    ↓
Network

This separation allows replacing transports without changing routing policy.


---

## Integration Phase

From this point the priority is runtime verification.

Rule:

One change → One test.

Avoid architectural refactoring until the first successful end-to-end connection is achieved.

Current objective:

Phone
    ↓
Hysteria
    ↓
Dispatcher
    ↓
Agent Driver
    ↓
Nebula
    ↓
XRM Agent
    ↓
Internet


---

## First Agent Routing Test

Temporary integration rule:

Route "italy" uses the Agent driver.

Reason:

This is an integration test only.

After end-to-end verification the routing layer may be redesigned without changing Policy.


---

## Integration Bug #1

Problem:

Dispatcher successfully resolves Route.

Handler ignores Route.Outbound.

StreamBridge therefore never receives the selected driver.

Current fix:

Pass Route.Outbound from Handler to StreamBridge.

No routing logic belongs inside StreamBridge.


---

## Runtime milestone

Current state:

- XRM runtime components exist.
- Agent exists.
- Dispatcher exists.
- Forwarder exists.
- StreamBridge exists.

BUT

No XRM server process exists yet.

Phones are still connected to the official Hysteria server started by systemd.

Therefore none of the XRM runtime code participates in real traffic.

Next milestone:

Implement "xrm serve".

Responsibilities:

- initialize Core
- initialize Hysteria server
- connect internal/hysteria.Handler
- expose the same listening interface as the official Hysteria server

Only after "xrm serve" is operational should routing, Agent and policy be tested.


---

## Runtime redesign

IMPORTANT

The current internal/hysteria package is NOT a real Hysteria server.

Current implementation contains only development stubs:

- TCP listener
- fake parser
- placeholder target resolver

It MUST NOT evolve into a custom VPN protocol.

The production runtime MUST embed the official Hysteria server library.

Target architecture:

Official Hysteria (QUIC)
        ↓
RequestHook
        ↓
common.Request
        ↓
Classifier
        ↓
Dispatcher
        ↓
Forwarder
        ↓
Driver
        ↓
Agent / Direct

The existing Core, Dispatcher, Forwarder and Agent remain unchanged.

Only the runtime entrypoint is replaced.


---

## Current milestone

Next implementation target:

Create the first runnable XRM server.

Acceptance criteria:

- binary: xrm serve
- starts official Hysteria server library
- listens on configured UDP port
- phones can connect
- RequestHook is invoked

Agent integration is NOT the next task.

The first goal is to replace the standalone hysteria process with the XRM runtime.


---

## Hysteria Core Integration Notes

Research completed on github.com/apernet/hysteria/core/v2.

Conclusions:

- core/v2 is a transport library only.
- It does NOT load YAML configuration.
- It does NOT provide a server application.
- It does NOT implement authentication.
- The embedding application is responsible for:
  - loading YAML;
  - creating server.Config;
  - implementing Authenticator;
  - opening UDP listener;
  - loading TLS certificates.

Therefore XRM owns runtime construction while embedding the official Hysteria transport.

Current integration chain:

internal/hysteria.Config
        ↓
internal/hycore.BuildConfig()
        ↓
hyserver.Config
        ↓
hyserver.NewServer()


---

## 2026-07-04 — Runtime refactoring (commit 9007f32)

### Completed

- Official Hysteria runtime is now the primary runtime used by XRM.
- Runtime successfully serves real clients through the integrated Hysteria core.
- Android client successfully connects through XRM.
- Traffic is forwarded through the remote VPS Agent.
- External IP is confirmed to be the Italian VPS.
- The official Hysteria daemon has been retired from the data path.
- xrm-agent is now deployed as a persistent systemd service on the VPS.

### Runtime architecture

Current runtime pipeline:

    Hysteria
        ↓
    StreamBridge
        ↓
    Forwarder
        ↓
    Outbound Manager
        ↓
    Driver (agent)

### Refactoring

Forwarder has been prepared for routing integration.

Injected dependencies:

- Classifier
- Dispatcher
- Outbound Manager

Behavior is intentionally unchanged.

No routing decisions are performed yet.

This commit only prepares the architecture for the next stage.

### Next milestone

Move route selection into Forwarder.

Target pipeline:

    Hysteria
        ↓
    StreamBridge
        ↓
    Forwarder
        ↓
    Classifier
        ↓
    Dispatcher
        ↓
    Route
        ↓
    Outbound
        ↓
    Driver

After this step Hysteria, Runtime and StreamBridge will no longer know anything about outbound drivers.

Commit:
9007f32

---

## 2026-07-04 — Routing enters the runtime (commit 6423b12)

Forwarder is no longer using the caller-selected outbound.

Connection flow is now:

    Hysteria
        ↓
    StreamBridge
        ↓
    Forwarder
        ↓
    Classifier
        ↓
    Dispatcher
        ↓
    Route
        ↓
    Outbound
        ↓
    Driver

Implemented:

- Forwarder now creates a Request.
- Request is classified.
- Dispatcher resolves the route.
- Route selects the outbound.
- Outbound resolves the driver.
- Driver performs the actual Dial().

The "driver" argument is now ignored and exists only for backward compatibility.

Next step:

Remove the obsolete driver parameter completely and migrate all callers to the new API.

Commit:
6423b12


## TEMP FILES (REMOVE AFTER STABILIZATION)

These files are temporary debugging tools and MUST be deleted later:

- cmd/xrm/test_main.go
- internal/core/test_forward.go

Purpose:
- validate forwarder → router flow
- isolate routing logic from outbound/inbound systems

