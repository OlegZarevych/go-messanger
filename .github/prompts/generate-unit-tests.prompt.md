---
mode: 'agent'
description: 'Generate a unit test for each new public Go function'
---

Your goal is to generate a Go unit test for every new public (exported) function added to this repository.

## Rules

- Create or update the `_test.go` file in the **same package** as the function under test.
- Use only the standard library: `testing`, `net/http/httptest`, etc. Do **not** add third-party test dependencies.
- Name each test `Test<FunctionName>` (e.g. `NewHelloHandler` → `TestNewHelloHandler`).
- Follow the **Arrange / Act / Assert** pattern with inline comments (`// Arrange`, `// Act`, `// Assert`).
- Cover at least:
  - The **happy path** (expected inputs produce expected outputs / behavior).
  - **Edge cases** relevant to the function (empty input, boundary values, error returns, etc.).
- For HTTP handlers, use `net/http/httptest.NewRequest` and `httptest.NewRecorder`; assert status code, headers, and body.
- Use `t.Fatalf` / `t.Errorf` for assertion failures (no third-party assertion libraries).
- Keep tests deterministic and side-effect free (no real network calls, no real file I/O unless the function explicitly requires it).

## Output format

Return **only** the complete content of the `_test.go` file (or the relevant additions to it). Do not include explanations outside of code comments.

## Example

Given this public function in `handler/hello.go`:

```go
func NewHelloHandler() http.Handler {
    return http.HandlerFunc(helloHandler)
}
```

Generate:

```go
package handler

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestNewHelloHandler(t *testing.T) {
    // Arrange
    req := httptest.NewRequest(http.MethodGet, "/hello", nil)
    rec := httptest.NewRecorder()

    // Act
    handler := NewHelloHandler()
    handler.ServeHTTP(rec, req)

    // Assert
    res := rec.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Fatalf("expected status 200, got %d", res.StatusCode)
    }

    body := rec.Body.String()
    expected := `{"message":"Hello, world!"}`
    if body != expected {
        t.Fatalf("expected body %q, got %q", expected, body)
    }
}
```
