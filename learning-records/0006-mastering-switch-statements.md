# Mastering Switch Statements in Go

The student successfully solved the `blackjack` exercise, demonstrating strong competence with `switch` statements, card value mapping, and nested decision logic.

## Evidence
- Implemented `ParseCard` using a `switch` statement to map card name strings to integer values.
- Implemented `FirstTurn` using a boolean `switch` with nested `if/else` for dealer-card checks.
- Refactored `ParseCard` from a boolean switch (`switch { case card == "ace": }`) to a cleaner value switch (`switch card { case "ace": }`).

## Lessons & Teachings
- Go requires every code path in a function to return a value. A `switch` without a `default` case does not guarantee a return, causing a compile error. Adding `default` resolves it.
- A **value switch** (`switch card { ... }`) is more idiomatic than a **boolean switch** (`switch { case card == "ace": ... }`) when matching one variable against many values.
- Find & Replace is faster than multicursor when making the exact same change across many locations.
