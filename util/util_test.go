package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	testString := "convert to camel case"
	assert.Equal(t, toCamelCase(testString), "convertToCamelCase")
}
func TestCovertPathtoCamelCaseMethodName(t *testing.T) {
	testString := "convert to camel case"
	method := "post"
	assert.Equal(t, CovertPathtoCamelCaseMethodName(testString, method), "postConvertToCamelCase")
}
