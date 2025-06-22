# go-fuzzing-example

A minimal Go repository demonstrating fuzz testing techniques and integration with `go-continuous-fuzzing`.

## Structure

- **parser**: Contains functions for parsing and evaluating expressions.
- **stringutils**: Provides string manipulation utilities.
- **tree**: Provides binary tree building utilities.

Each package includes:

- One "safe" function that is robust against crashes during fuzzing.
- One "unsafe" function that may crash or behave unexpectedly when fuzzed.
- Corresponding fuzz tests in their respective `*_test.go` files.

**NOTE**: The **tree** package contains only one function, which is an "unsafe" function that may crash or behave unexpectedly when fuzzed due to an out-of-memory (OOM) error.
