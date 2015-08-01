package OMDbapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {

	testData := []struct {
		queryTitle     string
		expectedGenre  string
		expectedRating string
		expectedVotes  string
		expectedResult string
	}{
		{
			queryTitle:     "game of throne",
			expectedGenre:  "Short, Comedy, Crime",
			expectedRating: "5.7",
			expectedVotes:  "11",
			expectedResult: "#tt3728462: Game of Throne (2014) Type: movie",
		},
	}

	for _, test := range testData {
		resp, err := QueryByTitle(test.queryTitle)
		assert.Nil(t, err)

		assert.Equal(t, resp.Genre, test.expectedGenre)
		assert.Equal(t, resp.IMDBRating, test.expectedRating)
		assert.Equal(t, resp.IMDBVotes, test.expectedVotes)
		assert.Equal(t, resp.String(), test.expectedResult)
	}
}
