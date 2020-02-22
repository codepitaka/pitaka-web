package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_URLs(t *testing.T) {
	assert.Equal(t, DevServerURL, "https://pitaka-server-dev.herokuapp.com")
	assert.Equal(t, PrdServerURL, "https://pitaka-server.herokuapp.com")
}
