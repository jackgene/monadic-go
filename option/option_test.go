package option

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("Maps non-empty Option", func(t *testing.T) {
		// Set up
		some := Wrap(1)

		// Test
		actualOutput := Map[int, string](some, func(in int) string {
			return fmt.Sprintf("%d", in)
		})

		// Verify
		expectedOutput := Wrap[string]("1")
		assert.Equal(t, expectedOutput, actualOutput)
	})

	t.Run("Maps empty Option", func(t *testing.T) {
		// Set up
		some := None[int]()

		// Test
		actualOutput := Map[int, string](some, func(in int) string {
			return fmt.Sprintf("%d", in)
		})

		// Verify
		expectedOutput := None[string]()
		assert.Equal(t, expectedOutput, actualOutput)
	})
}
