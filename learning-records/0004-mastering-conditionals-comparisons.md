# Mastering Conditionals and Comparisons in Go

The student successfully solved the `vehicle-purchase` exercise, demonstrating strong competence in logical operations, conditional paths, and lexicographical string comparisons.

## Evidence
- Used `||` (logical OR) to check vehicle types in `NeedsLicense`.
- Used `<` for string comparison in `ChooseVehicle` to determine alphabetical precedence.
- Implemented multiple `if` branches in `CalculateResellPrice` using conditions like `age < 3`, `age < 10.0 && age >= 3`, and a default fallback.

## Lessons & Teachings
We discussed how boolean conditions can be returned directly instead of wrapped in an `if-return-bool` block, and how structure control flow can be simplified (removing redundant check `age >= 3` because the previous block already returns if `age < 3`).
