# Error Handling and Conditional Loops

The student successfully solved the `collatz-conjecture` exercise, demonstrating understanding of returning standard errors using Go's `errors` package, handling nil errors on success, and using conditional "while-style" loops.

## Evidence

- Correctly returned an error value (`errors.New`) for invalid inputs (`n <= 0`) and returned `nil` on success.
- Identified and resolved an infinite cycle issue (`1 -> 4 -> 2 -> 1`) by changing the loop condition from `n > 0` to `n > 1`.
- Implemented and verified the Collatz algorithm.

## Implications

- The student is now comfortable with basic error returning in Go.
- Ready for exercises requiring error propagation, error wrapping, or more advanced loop controls.
