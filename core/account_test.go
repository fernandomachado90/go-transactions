package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {

	tests := map[string]func(*testing.T){
		"Should initialize entity": func(t *testing.T) {
			//given
			id := 1
			documentNumber := "12345678900"

			// when
			account := &Account{
				ID:             id,
				DocumentNumber: documentNumber,
			}

			// then
			assert.NotEmpty(t, account)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
