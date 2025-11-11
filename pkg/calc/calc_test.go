package calc

import "testing"

func almostEqual(a, b float64) bool {
    const eps = 1e-9
    if a > b {
        return a-b < eps
    }
    return b-a < eps
}

func TestAdd(t *testing.T) {
    if got := Add(1, 2, 3.5); !almostEqual(got, 6.5) {
        t.Fatalf("Add: got %v", got)
    }
}

func TestMul(t *testing.T) {
    if got := Mul(2, 3, 4); !almostEqual(got, 24) {
        t.Fatalf("Mul: got %v", got)
    }
    if got := Mul(); !almostEqual(got, 1) {
        t.Fatalf("Mul empty: got %v", got)
    }
}

func TestSub(t *testing.T) {
    if got := Sub(10, 1, 2, 3); !almostEqual(got, 4) {
        t.Fatalf("Sub: got %v", got)
    }
    if got := Sub(); !almostEqual(got, 0) {
        t.Fatalf("Sub empty: got %v", got)
    }
}

func TestDiv(t *testing.T) {
    if got, err := Div(8, 2, 2); err != nil || !almostEqual(got, 2) {
        t.Fatalf("Div: got %v err %v", got, err)
    }
    if _, err := Div(1, 0); err == nil {
        t.Fatalf("Div by zero should error")
    }
    if _, err := Div(); err == nil {
        t.Fatalf("Div empty should error")
    }
}

func TestPow(t *testing.T) {
    if got := Pow(2, 8); !almostEqual(got, 256) {
        t.Fatalf("Pow: got %v", got)
    }
}

func TestSqrt(t *testing.T) {
    if got, err := Sqrt(9); err != nil || !almostEqual(got, 3) {
        t.Fatalf("Sqrt: got %v err %v", got, err)
    }
    if _, err := Sqrt(-1); err == nil {
        t.Fatalf("Sqrt negative should error")
    }
}

