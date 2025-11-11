// Command gocalculate provides a simple CLI calculator.
package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"

    "gocalculate/pkg/calc"
)

func usage() {
    // Print help/usage information for the CLI.
    fmt.Println("gocalculate - simple calculator")
    fmt.Println("Usage:")
    fmt.Println("  gocalculate <command> [numbers...]")
    fmt.Println("Commands:")
    fmt.Println("  add a b [c ...]     Sum numbers")
    fmt.Println("  sub a b [c ...]     Subtract subsequent from first")
    fmt.Println("  mul a b [c ...]     Multiply numbers")
    fmt.Println("  div a b [c ...]     Divide first by each subsequent (no zero)")
    fmt.Println("  pow a b             a^b")
    fmt.Println("  sqrt a              Square root of a (a>=0)")
    fmt.Println("  repl                Interactive mode (type commands)")
}

// parseFloats converts string args to float64 values.
func parseFloats(args []string) ([]float64, error) {
    vals := make([]float64, 0, len(args))
    for _, s := range args {
        v, err := strconv.ParseFloat(s, 64)
        if err != nil {
            return nil, fmt.Errorf("invalid number '%s'", s)
        }
        vals = append(vals, v)
    }
    return vals, nil
}

// runCommand executes a calculator command and returns output, exit code, and error.
func runCommand(cmd string, args []string) (string, int, error) {
    switch cmd {
    case "add":
        // add requires at least one operand.
        if len(args) == 0 {
            return "", 2, errors.New("add requires at least one number")
        }
        nums, err := parseFloats(args)
        if err != nil { return "", 2, err }
        return fmt.Sprintf("%g", calc.Add(nums...)), 0, nil
    case "sub":
        // sub with one operand returns that operand.
        if len(args) == 0 {
            return "", 2, errors.New("sub requires at least one number")
        }
        nums, err := parseFloats(args)
        if err != nil { return "", 2, err }
        return fmt.Sprintf("%g", calc.Sub(nums...)), 0, nil
    case "mul":
        // mul over zero operands is invalid at CLI level.
        if len(args) == 0 {
            return "", 2, errors.New("mul requires at least one number")
        }
        nums, err := parseFloats(args)
        if err != nil { return "", 2, err }
        return fmt.Sprintf("%g", calc.Mul(nums...)), 0, nil
    case "div":
        // div requires at least two operands; prevents divide-by-zero.
        if len(args) < 2 {
            return "", 2, errors.New("div requires at least two numbers")
        }
        nums, err := parseFloats(args)
        if err != nil { return "", 2, err }
        res, derr := calc.Div(nums...)
        if derr != nil { return "", 2, derr }
        return fmt.Sprintf("%g", res), 0, nil
    case "pow":
        // pow expects exactly two operands: base and exponent.
        if len(args) != 2 {
            return "", 2, errors.New("pow requires exactly two numbers")
        }
        nums, err := parseFloats(args)
        if err != nil { return "", 2, err }
        return fmt.Sprintf("%g", calc.Pow(nums[0], nums[1])), 0, nil
    case "sqrt":
        // sqrt expects exactly one non-negative operand.
        if len(args) != 1 {
            return "", 2, errors.New("sqrt requires exactly one number")
        }
        nums, err := parseFloats(args)
        if err != nil { return "", 2, err }
        res, serr := calc.Sqrt(nums[0])
        if serr != nil { return "", 2, serr }
        return fmt.Sprintf("%g", res), 0, nil
    default:
        return "", 2, fmt.Errorf("unknown command: %s", cmd)
    }
}

// repl starts a simple interactive prompt that accepts the same commands.
func repl() int {
    fmt.Println("gocalculate REPL. Type 'exit' to quit.")
    in := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("> ")
        if !in.Scan() {
            break
        }
        line := strings.TrimSpace(in.Text())
        if line == "" {
            continue
        }
        if line == "exit" || line == "quit" {
            return 0
        }
        parts := strings.Fields(line)
        cmd := parts[0]
        args := []string{}
        if len(parts) > 1 {
            args = parts[1:]
        }
        out, code, err := runCommand(cmd, args)
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        if out != "" {
            fmt.Println(out)
        }
        if code != 0 {
            return code
        }
    }
    if err := in.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "input error:", err)
        return 1
    }
    return 0
}

func main() {
    // If no subcommand is provided, print usage and exit with code 2.
    if len(os.Args) < 2 {
        usage()
        os.Exit(2)
    }
    cmd := os.Args[1]
    // Launch interactive mode when requested.
    if cmd == "repl" {
        os.Exit(repl())
    }
    out, code, err := runCommand(cmd, os.Args[2:])
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(code)
    }
    if out != "" {
        fmt.Println(out)
    }
    os.Exit(code)
}
