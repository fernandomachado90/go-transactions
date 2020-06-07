package core

import "time"

type Transaction struct {
	ID          int
	AccountID   int
	OperationID int
	Amount      float64
	EventDate   time.Time
}
