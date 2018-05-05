package judge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupOrganization(t *testing.T) {
	t.Run("it work", func(t *testing.T) {
		err := SetupNewOrganization()
		assert.Nil(t, err)
	})
}
