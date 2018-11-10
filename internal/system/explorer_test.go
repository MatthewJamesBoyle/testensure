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
	t.Run("Correctly returns untested files", func(t *testing.T) {
		untest, err := Explore("../")
		assert.NoError(t, err)
		assert.Equal(t, []string{"notTested.go"}, untest)
	})
}
