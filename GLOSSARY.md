# Go Programming Glossary

This glossary tracks core programming and Go terms as they are introduced and fully understood throughout your learning journey.

## Terms

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

**String Formatting**:
The process of constructing a new string by inserting formatted values (such as variables or expressions) into a template string, typically using functions like `fmt.Sprintf`.

**Format Verb**:
A placeholder sequence (like `%s`, `%d`, or `%f`) inside a formatting template string that dictates how a value of a corresponding type should be represented as text.

**Float (Floating-point Number)**:
A data type representing real numbers that contain a decimal point (e.g., `float64` in Go).

**Comparison Operator**:
Symbols used to compare two values (e.g., `==`, `!=`, `<`, `>`, `<=`, `>=`). The result is always a boolean value (`true` or `false`).

**Logical Operator**:
Operators used to combine multiple boolean conditions (e.g., `&&` for AND, `||` for OR, `!` for NOT).

**If/Else Statement**:
A control flow structure that runs a block of code if a condition is true, optionally running alternative blocks (`else if` or `else`) if the condition is false.

**Switch Statement**:
A control flow structure that matches a value (or boolean conditions) against a list of `case` branches, executing the first match. A `default` case runs if no other case matches. More readable than long `if/else if` chains when checking one value against many possibilities.
