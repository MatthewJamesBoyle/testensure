package system

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExplore(t *testing.T) {
	t.Run("returns error if error occurs in walk func", func(t *testing.T) {
		untest, err := Explore("someInvalidPath")
		assert.Error(t, err)
		assert.Nil(t, untest)
	})
	t.Run("A directory with no go files returns nil", func(t *testing.T) {
		untest, err := Explore("../testPack/nongodir")
		assert.NoError(t, err)
		assert.Equal(t, 0, len(untest))
	})
	t.Run("Correctly returns untested files", func(t *testing.T) {
		var expectedUntested = []string{"notTested.go", "alsoNotTested.go"}
		untest, err := Explore("../")
		assert.NoError(t, err)
		assert.Equal(t, expectedUntested, untest)
	})
}
