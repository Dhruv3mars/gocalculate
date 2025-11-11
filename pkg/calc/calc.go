package calc

import (
    "errors"
    "math"
)

// Add returns the sum of all provided numbers.
func Add(nums ...float64) float64 {
    var sum float64
    for _, n := range nums {
        sum += n
    }
    return sum
}

// Mul returns the product of all provided numbers. Empty input yields 1.
func Mul(nums ...float64) float64 {
    if len(nums) == 0 {
        return 1
    }
    p := 1.0
    for _, n := range nums {
        p *= n
    }
    return p
}

// Sub subtracts all subsequent numbers from the first.
// If only one number is provided, it returns that number.
func Sub(nums ...float64) float64 {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    res := nums[0]
    for _, n := range nums[1:] {
        res -= n
    }
    return res
}

// ErrDivByZero is returned when attempting to divide by zero.
var ErrDivByZero = errors.New("division by zero")

// Div divides the first number by each subsequent number in order.
// Returns an error if any divisor is zero or if no numbers provided.
func Div(nums ...float64) (float64, error) {
    if len(nums) == 0 {
        return 0, errors.New("no numbers provided")
    }
    res := nums[0]
    for _, n := range nums[1:] {
        if n == 0 {
            return 0, ErrDivByZero
        }
        res /= n
    }
    return res, nil
}

// Pow raises base to the power exp.
func Pow(base, exp float64) float64 {
    return math.Pow(base, exp)
}

// ErrSqrtOfNegative is returned when attempting sqrt of a negative number.
var ErrSqrtOfNegative = errors.New("square root of negative number")

// Sqrt returns the square root of x, or an error if x < 0.
func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrSqrtOfNegative
    }
    return math.Sqrt(x), nil
}

