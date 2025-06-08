# go-fuzzing-example

A minimal Go repository demonstrating fuzz testing techniques and integration with `go-continuous-fuzzing`.

## Structure

- **parser**: Contains functions for parsing and evaluating expressions.
- **stringutils**: Provides string manipulation utilities.

Each package includes:

- One "safe" function that is robust against crashes during fuzzing.
- One "unsafe" function that may crash or behave unexpectedly when fuzzed.
- Corresponding fuzz tests in their respective `*_test.go` files.
