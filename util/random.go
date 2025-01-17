package util

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomNumeric generates a random NUMERIC(10, 2) value
func RandomNumeric(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	random := min + rand.Float64()*(max-min)
	return float64(int(random*100)) / 100 // Round to 2 decimal places
}

// Float64ToBigInt converts a float64 to a big.Int

// func Float64ToBigInt(f float64) *big.Int {

// 	bigFloat := new(big.Float).SetFloat64(f)

// 	bigInt, _ := bigFloat.Int(nil)

// 	return bigInt

// }
func Float64ToBigInt(value float64) *big.Int {
	// Round to two decimal places before converting to BigInt
	roundedValue := math.Round(value * 100) // Scale by 100 for two decimal precision

	// Convert the rounded value to an integer and return as BigInt
	return big.NewInt(int64(roundedValue))
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// AddNumeric adds two pgtype.Numeric values and returns the result

func AddNumeric(a, b pgtype.Numeric) pgtype.Numeric {

	result := pgtype.Numeric{}

	// Assuming a and b are valid and have the same scale

	result.Int = new(big.Int).Add(a.Int, b.Int)

	result.Exp = a.Exp

	result.Valid = a.Valid && b.Valid

	return result

}
