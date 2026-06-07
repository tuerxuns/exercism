# Structs and Constructor Functions

The student successfully solved the `need-for-speed` exercise, demonstrating clear understanding of struct definition, the `New*` constructor pattern, value semantics, and integer division for capacity calculations. Key lessons included: Go structs are value types passed by copy — returning a new struct is the idiomatic way to "update" state; the `New*` constructor pattern is Go's substitute for class constructors; and named field initialisation is preferred over positional to avoid silent breakage on reordering.

## Evidence

- Defined `Car` and `Track` structs with unexported fields (`battery`, `batteryDrain`, `speed`, `distance`).
- Implemented `NewCar` and `NewTrack` constructor functions that set default values (e.g. `battery: 100`).
- Implemented `Drive` using value semantics — receiving a `Car` by value and returning a modified copy rather than mutating via a pointer.
- Implemented `CanFinish` using integer division to calculate the maximum driveable distance.

## Implications

- Understands struct definition, constructor patterns, and value semantics in Go.
- Ready for exercises involving slices, variadic functions, and iteration patterns.
