# Fault-Injection Coverage Metric

## Why not just pass/fail
Test suites answer "did the known cases we wrote still pass." They do not answer
"how much of the known fault space has a verified recovery classification."
Coverage regressions (a fault type silently returning `unknown`) can hide behind
a fully green test run if no test exercises that specific (type, severity) pair.

## Metric
`pkg/faults.CoverageTracker` records every `(type, severity)` fault class that
produces a non-`unknown` classification. `Report()` returns:
- `Covered` / `Total`: how many of the known classes classified successfully.
- `Percent`: coverage percentage.
- `Uncovered`: the exact classes that did not classify, for triage.

Known classes are the cross product of fault types (`network-partition`,
`crash`, `time-skew`, `corruption`) and severities (`low`, `medium`, `high`),
defined in `faults.KnownFaultClasses()`.

## Usage
```
go run ./tools/fault-injector -report
```
Exits non-zero if coverage is below 100%, so it can gate CI the same way a
test failure would, without requiring a hand-written test per class.

## Extending
Adding a new fault type or severity to the injector should also update
`KnownFaultClasses()` so the coverage metric tracks it automatically.
