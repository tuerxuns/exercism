# Go Programming Glossary

This glossary tracks core programming and Go terms as they are introduced and fully understood throughout your learning journey.

## Syntax & Basics

**Package**:
A collection of Go source files in the same directory that compile together. All files in the directory must declare the same package name.

**Function**:
A reusable block of code that takes zero or more inputs (arguments) and returns zero or more outputs.

**Variable**:
A named storage location in memory that holds a value of a specific type, which can be modified during program execution.

**Constant**:
A fixed value that cannot be changed during program execution, declared using the `const` keyword.

**Type**:
A classification of data that determines what values a variable can hold and what operations can be performed on it (e.g., `string`, `int`, `float64`, `bool`).

**Doc Comment**:
A code comment starting with `//` or `/*` placed immediately before a declaration (like a package, type, or function) without any intervening empty lines, which is used to generate official API documentation.

**Multiple Return Values**:
A Go function feature allowing more than one value to be returned from a single function. Declared as comma-separated types in parentheses in the signature (e.g. `func Quantities(layers []string) (int, float64)`), and returned as a comma-separated list after `return`.

## Data Types & Collections

**Float (Floating-point Number)**:
A data type representing real numbers that contain a decimal point (e.g., `float64` in Go).

**Error**:
A built-in interface type in Go used to represent an abnormal state or failure. Functions commonly return an `error` as their last return value, which is `nil` if the operation succeeded, or a non-nil value created using `errors.New` or `fmt.Errorf` if it failed.

**Struct**:
A user-defined type that represents a collection of fields grouped together under a single name, used to define custom data models.

**Slice**:
A dynamically-sized, flexible view into the elements of an underlying array, consisting of a pointer, a length, and a capacity.

**Slice Aliasing**:
When two slice variables share the same underlying array (e.g. `s2 := s1`), so mutations through one are visible through the other. Use `append([]T(nil), s...)` or `copy` to make a true independent copy.

**Blank Identifier (`_`)**:
A special write-only variable name in Go used to discard a value you don't need (e.g. the index in `for _, v := range slice`).

## Control Flow

**If/Else Statement**:
A control flow structure that runs a block of code if a condition is true, optionally running alternative blocks (`else if` or `else`) if the condition is false.

**Switch Statement**:
A control flow structure that matches a value (or boolean conditions) against a list of `case` branches, executing the first match. A `default` case runs if no other case matches. More readable than long `if/else if` chains when checking one value against many possibilities.

**For Loop**:
Go's only looping construct, used for all repetition. The standard form is `for init; condition; post { }`. It can also act as a while-style loop (`for condition { }`) or an infinite loop (`for { }`).

**Range**:
A keyword used with `for` to iterate over slices, arrays, maps, strings, channels, or (Go 1.22+) integers. For slices it yields the index and value on each iteration: `for i, v := range slice`. Either can be discarded with `_`.

## Logic & Operations

**Comparison Operator**:
Symbols used to compare two values (e.g., `==`, `!=`, `<`, `>`, `<=`, `>=`). The result is always a boolean value (`true` or `false`).

**Logical Operator**:
Operators used to combine multiple boolean conditions (e.g., `&&` for AND, `||` for OR, `!` for NOT).

**Spread Operator (`...`)**:
An ellipsis syntax used to either define a variadic parameter in a function signature or unpack a slice's elements to pass them individually to a variadic function.

## Strings & Formatting

**Strings Package**:
A standard library package in Go (`strings`) that contains utilities for searching, replacing, trimming, and manipulating UTF-8 strings.

**String Formatting**:
The process of constructing a new string by inserting formatted values (such as variables or expressions) into a template string, typically using functions like `fmt.Sprintf`.

**Format Verb**:
A placeholder sequence (like `%s`, `%d`, or `%f`) inside a formatting template string that dictates how a value of a corresponding type should be represented as text.

## Programming Paradigms & Patterns

**Constructor Function**:
An idiomatic Go pattern (typically named `New*`) used to initialize and return a new instance of a struct, often with pre-configured default values.

**Value Semantics**:
A programming paradigm where data is copied when passed to functions or returned, preventing unintended side effects from shared mutable state.

**Default Value Fallback**:
A common Go pattern where a function checks if a parameter is its zero value and substitutes a sensible default instead. Since Go has no built-in default parameter values, this pattern simulates them (e.g. `if time == 0 { timePerLayer = 2 }`).
