gocalculate
============

Simple, dependency‑free calculator in Go with a CLI and reusable library package.

Features
- Subcommands: add, sub, mul, div, pow, sqrt
- Accepts multiple operands for add/mul/sub/div (sequential semantics)
- Friendly errors (division by zero, invalid/insufficient args)
- Tiny REPL: `gocalculate repl` for quick interactive use
- Reusable library in `pkg/calc`

Install
- With Go installed (1.20+ recommended):
  - From source in this repo: `go build -o gocalculate ./cmd/gocalculate`
  - Show usage without building: `go run ./cmd/gocalculate`

Usage
- General: `gocalculate <command> [numbers...]`
- Commands:
  - `add a b [c ...]`        Sum numbers
  - `sub a b [c ...]`        Subtract subsequent from first
  - `mul a b [c ...]`        Multiply numbers
  - `div a b [c ...]`        Divide first by each subsequent (no zero)
  - `pow a b`                a^b
  - `sqrt a`                 Square root of a (a >= 0)
  - `repl`                   Interactive shell (type commands)
  - Tip: run with no arguments to see usage

Examples
```
$ gocalculate add 1 2 3
6

$ gocalculate sub 10 1 2
7

$ gocalculate mul 2 5
10

$ gocalculate div 8 2 2
2

$ gocalculate pow 2 8
256

$ gocalculate sqrt 9
3

$ gocalculate repl
gocalculate REPL. Type 'exit' to quit.
> add 1 2
3
> sqrt 16
4
> exit
```

Library Usage
```
import "gocalculate/pkg/calc"

sum := calc.Add(1, 2, 3)
prod := calc.Mul(2, 3, 4)
diff := calc.Sub(10, 1, 2)
q, err := calc.Div(8, 2, 2) // q=2, err=nil
p := calc.Pow(2, 8)         // 256
r, err := calc.Sqrt(9)      // r=3, err=nil
```

Development
- Run tests: `go test ./...`
- Build CLI: `go build -o gocalculate ./cmd/gocalculate`

Exit Codes
- `0` on success
- `2` on user errors (bad args, div by zero, etc.)

Notes
- All math uses `float64`.
- Division and subtraction operate left‑to‑right over the provided operands.
