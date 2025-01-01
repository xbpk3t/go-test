package decimals

import (
	"github.com/shopspring/decimal"
	"math/big"
	"math/rand"
	"testing"
)

const (
	numOperations = 1000000
)

// Benchmark for int operations
func BenchmarkIntOperations(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int64
		for j := 0; j < numOperations; j++ {
			a := int64(rand.Intn(1000))
			b := int64(rand.Intn(1000))
			sum += a + b
			sum += a - b
			sum += a * b
			if b != 0 {
				sum += a / b
			}
		}
	}
}

// Benchmark for decimal operations
func BenchmarkDecimalOperations(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum decimal.Decimal
		for j := 0; j < numOperations; j++ {
			a := decimal.NewFromInt(int64(rand.Intn(1000)))
			b := decimal.NewFromInt(int64(rand.Intn(1000)))
			sum = sum.Add(a.Add(b))
			sum = sum.Add(a.Sub(b))
			sum = sum.Add(a.Mul(b))
			if !b.IsZero() {
				sum = sum.Add(a.Div(b))
			}
		}
	}
}

func BenchmarkBigFloatOperations(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := new(big.Float)
		for j := 0; j < numOperations; j++ {
			a := new(big.Float).SetFloat64(rand.Float64() * 1000)
			b := new(big.Float).SetFloat64(rand.Float64() * 1000)
			sum.Add(sum, a.Add(a, b))
			sum.Add(sum, a.Sub(a, b))
			sum.Add(sum, a.Mul(a, b))
			if b.Cmp(big.NewFloat(0)) != 0 {
				sum.Add(sum, a.Quo(a, b))
			}
		}
	}
}
