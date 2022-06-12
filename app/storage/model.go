package storage

import (
	"google.golang.org/genproto/googleapis/type/datetime"
)

type Transaction struct {
	ID            uint64            `db:"id" json:"id"`
	UserID        string            `db:"userid" json:"userid"`
	UserEmail     string            `db:"useremail" json:"useremail"`
	Sum           float64           `db:"sum" json:"sum"`
	Currency      string            `db:"currency" json:"currency"`
	DTCreated     datetime.DateTime `db:"dtcreated" json:"dtcreated"`
	DTLastChanged datetime.DateTime `db:"dtlastchanged" json:"dtlastchanged"`
	Status        string            `db:"status" json:"status"`
}
