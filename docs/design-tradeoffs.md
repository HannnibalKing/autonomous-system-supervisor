# Design Tradeoffs

- Determinism vs adaptivity: deterministic seeds and stable ordering ease replay; ML hooks kept pluggable to avoid nondeterministic drift in core loop.
- Single binary vs sidecars: agents stay thin; control-plane may host plugins but avoids heavy sidecars to keep latency predictable.
- State store: pluggable (raft/etcd/embedded) so deployments choose durability level; demo uses in-memory.
- Fault classification: simple heuristics now, interface allows richer models later without changing agent wire format.
- Migration: interface-only in baseline to keep scaffold simple; real deployment can back with container runtime or process checkpoint.
