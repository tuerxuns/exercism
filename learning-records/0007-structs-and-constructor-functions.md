# Structs and Constructor Functions in Go

The student successfully solved the `need-for-speed` exercise, demonstrating clear understanding of struct definition, the `New*` constructor pattern, value semantics, and integer division for capacity calculations.

## Evidence
- Defined `Car` and `Track` structs with unexported fields (`battery`, `batteryDrain`, `speed`, `distance`).
- Implemented `NewCar` and `NewTrack` constructor functions that set default values (e.g. `battery: 100`).
- Implemented `Drive` using value semantics — receiving a `Car` by value and returning a modified copy rather than mutating via a pointer.
- Implemented `CanFinish` using integer division (`(car.battery / car.batteryDrain) * car.speed`) to calculate the maximum driveable distance.

## Lessons & Teachings
- Go structs are **value types** — passing a struct to a function gives it a copy. Returning a new struct is the idiomatic way to "update" state when not using pointer receivers.
- The `New*` constructor pattern is Go's substitute for class constructors. It enables setting defaults, validating inputs, and accessing unexported fields from within the same package.
- Named field initialisation (`Car{battery: 100, speed: speed}`) is preferred over positional initialisation — positional breaks silently if fields are reordered or a new field is added.
- Integer division truncates naturally, which is the correct behaviour when calculating "how many full drives are possible" from remaining battery.
