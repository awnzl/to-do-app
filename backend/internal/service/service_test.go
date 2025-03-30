package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.Implements(t, (*TodoService)(nil), &todoService{})

	//TODO AW: add testing for NewTodoService
}
