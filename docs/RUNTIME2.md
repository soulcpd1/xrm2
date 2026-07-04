# XRM2 Runtime Log (Part 2)

## Runtime status (July 2026)

Routing pipeline is now assembled:

Request
  ↓
Classifier
  ↓
Dispatcher
  ↓
Route
  ↓
Forwarder
  ↓
Outbound driver
  ↓
Network

### Completed

- Dispatcher integrated into runtime.
- Forwarder performs route lookup itself.
- StreamBridge no longer knows about drivers.
- Hysteria handler delegates routing to Core.
- Hycore outbound delegates routing to Forwarder.
- Agent outbound implemented.
- Direct outbound implemented.
- Nebula outbound implemented.
- Session routing preserved.
- Session cleanup fixed.
- Rule engine supports:
  - Domain
  - Suffix
  - Keyword
  - CIDR
  - Private ranges
- Default rule set extracted.
- Rule Loader API introduced.
- Runtime tests passing.

## Current architecture

Classifier
    ↓
Policy
    ↓
Dispatcher
    ↓
Route
    ↓
Forwarder
    ↓
Outbound Manager
    ├── direct
    ├── nebula
    └── agent

## Next milestone

Replace hardcoded rules with external configuration.

Planned:

- configs/rules.yaml
- rules.LoadFile()
- live rule reload
- hot runtime update
- per-country rule groups
- per-user rule groups
- geoip/geosite support
- balancing
- failover
- latency-based routing

