package entity

type Status int

const (
	Unreconciled Status = iota
	Reconciled
	Void
	FollowUp
	Duplicate
)

var Statuses = []string{"Unreconciled", "Reconciled", "Void", "Follow Up", "Duplicate"}

type StatusRegistry struct {
	Unreconciled Status
	Reconciled   Status
	Void         Status
	FollowUp     Status
	Duplicate    Status
}
