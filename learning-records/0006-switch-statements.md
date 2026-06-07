# Mastering Switch Statements

The student successfully solved the `blackjack` exercise, demonstrating strong competence with `switch` statements, card value mapping, and nested decision logic. Key lessons included: a switch without a `default` case does not guarantee a return, causing a compile error; and a value switch is more idiomatic than a boolean switch when matching one variable against many values.

## Evidence

- Implemented `ParseCard` using a `switch` statement to map card name strings to integer values.
- Implemented `FirstTurn` using a boolean `switch` with nested `if/else` for dealer-card checks.
- Refactored `ParseCard` from a boolean switch (`switch { case card == "ace": }`) to a cleaner value switch (`switch card { case "ace": }`).

## Implications

- Comfortable with both value and boolean switch patterns in Go.
- Ready for exercises involving structs, slices, and more complex data-driven logic.
