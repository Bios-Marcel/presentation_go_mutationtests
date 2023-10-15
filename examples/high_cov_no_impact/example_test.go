package highcovnoimpact

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_do(t *testing.T) {
	t.Cleanup(cleanup)

	for _, x := range []int{math.MinInt32, -1, 0, 1, 5, math.MaxInt32} {
		t.Run(fmt.Sprint(x), func(t *testing.T) {
			do(x)
		})
	}
}

func Test_do_improved(t *testing.T) {
	t.Skip()

	t.Cleanup(cleanup)

	t.Run("greater than zero", func(t *testing.T) {
		for _, x := range []int{1, 5, math.MaxInt32} {
			t.Run(fmt.Sprint(x), func(t *testing.T) {
				do(x)

				{
					_, err := os.Stat(fmt.Sprintf("%d_A", x))
					assert.NoError(t, err, "File A should exist")
				}

				// Improve tests further for custom mutation
				// {
				// 	_, err := os.Stat(fmt.Sprintf("%d_B", x))
				// 	assert.NoError(t, err, "File B should exist")
				// }
			})
		}
	})

	t.Run("less than or equal to zero", func(t *testing.T) {
		for _, x := range []int{math.MinInt32, -1, 0} {
			t.Run(fmt.Sprint(x), func(t *testing.T) {
				do(x)

				{
					_, err := os.Stat(fmt.Sprintf("%d_A", x))
					assert.ErrorIs(t, err, os.ErrNotExist, "File A should not exist")
				}
			})
		}
	})
}

func cleanup() {
	aMatches, _ := filepath.Glob("**_A")
	bMatches, _ := filepath.Glob("**_B")
	for _, a := range append(aMatches, bMatches...) {
		os.Remove(a)
	}
}
