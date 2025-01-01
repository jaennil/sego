package entity

import (
	"errors"
	"time"
)

type Transaction struct {
	Date   time.Time
	Status Status
	Typ    Type
	Amount float64
}

type Status string

const (
	Unreconciled Status = "Unreconciled"
	Reconciled   Status = "Reconciled"
	Void         Status = "Void"
	FollowUp     Status = "Follow Up"
	Duplicate    Status = "Duplicate"
)

func Statuses() []string {
	return []string{string(Unreconciled), string(Reconciled), string(Void), string(FollowUp), string(Duplicate)}
}

func NewStatus(status string) (Status, error) {
	switch status {
	case string(Unreconciled):
		return Unreconciled, nil
	case string(Reconciled):
		return Reconciled, nil
	case string(Void):
		return Void, nil
	case string(FollowUp):
		return FollowUp, nil
	case string(Duplicate):
		return Duplicate, nil
	default:
		return "", errors.New("invalid status")
	}
}

type Type string

const (
	Withdrawal Type = "Withdrawal"
	Deposit    Type = "Deposit"
)

func Types() []string {
	return []string{string(Withdrawal), string(Deposit)}
}

func NewType(typ string) (Type, error) {
	switch typ {
	case string(Withdrawal):
		return Withdrawal, nil
	case string(Deposit):
		return Deposit, nil
	default:
		return "", errors.New("invalid type")
	}
}
