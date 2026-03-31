# GitHub Copilot Custom Instructions

## Function Documentation Rule

Every function must have a comment directly above its declaration that describes what the function does.

- The comment must start with the function name (following Go doc comment conventions).
- The comment must clearly describe the purpose, behavior, inputs, and outputs of the function.
- Single-line comments (`//`) are preferred for concise descriptions; multi-line block comments (`/* */`) may be used for complex functions requiring more detail.

### Example

```go
// Add returns the sum of two integers a and b.
func Add(a, b int) int {
    return a + b
}
```

Apply this rule to all new and existing functions, including unexported (private) functions.