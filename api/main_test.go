package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	tests := map[string]func(*testing.T){
		"Hello": func(t *testing.T) {
			//given

			// when
			expected := Main()

			// then
			assert.True(t, expected)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
