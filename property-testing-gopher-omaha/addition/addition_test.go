package addition

import (
	"math"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func Add(x, y int) int {
	return x + y
}

func Test_Add(t *testing.T) {
	properties := gopter.NewProperties(
		gopter.DefaultTestParametersWithSeed(180925), // Date of talk
	)

	properties.Property("Given two positives, output is greater than either", prop.ForAll(
		func(nums []interface{}) bool {
			num1 := int(nums[0].(uint))
			num2 := int(nums[1].(uint))
			result := Add(num1, num2)
			return num1 < result && num2 < result
		},
		gopter.CombineGens(gen.UInt(), gen.UInt()),
	))
	properties.Property("Given two negatives, output is less than either", prop.ForAll(
		func(nums []interface{}) bool {
			num1 := nums[0].(int)
			num2 := nums[1].(int)
			result := Add(num1, num2)
			return result < num1 && result < num2
		},
		gopter.CombineGens(
			gen.IntRange(math.MinInt32/2, 0),
			gen.IntRange(math.MinInt32/2, 0),
		),
	))
	properties.Property("Given one negative and one positive, output should be between them", prop.ForAll(
		func(nums []interface{}) bool {
			num1 := nums[0].(int)
			num2 := nums[1].(int)
			result := Add(num1, num2)
			return num1 < result && result < num2
		},
		gopter.CombineGens(
			gen.IntRange(math.MinInt32/2, 0),
			gen.IntRange(0, math.MaxInt32/2),
		),
	))

	properties.TestingRun(t)
}
