package OMDbapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	resp, err := query("game of throne")
	assert.Nil(t, err)
	assert.Equal(t, resp.Title, "Game of Throne")
}
