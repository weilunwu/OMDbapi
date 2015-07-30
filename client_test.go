package OMDbapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testQuery(t *testing.T) {
	_, err := query("game of throne")
	assert.Nil(t, err)
}
