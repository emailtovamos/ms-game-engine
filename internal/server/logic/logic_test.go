package logic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAverage(t *testing.T){
	average := Average()
	assert.Equal(t, 3.0, average)
}