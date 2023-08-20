package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	var data = map[string]interface{}{
		"Name": "John Wick",
		"Age":  27,
	}
	var result = data["Name"]
	var resultString = result.(string)

	assert.Equal(t, "John Wick", resultString)
}
