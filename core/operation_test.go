package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperation(t *testing.T) {

	tests := map[string]func(*testing.T){
		"Should initialize entity": func(t *testing.T) {
			//given
			id := 4
			description := "PAGAMENTO"
			credit := true

			// when
			operation := &Operation{
				ID:          id,
				Description: description,
				Credit:      credit,
			}

			// then
			assert.NotEmpty(t, operation)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
