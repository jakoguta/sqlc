// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package override

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bar struct {
	ID      pgtype.Text `type:"id"`
	OtherID pgtype.Text `type:"other_id"`
	About   pgtype.Text
	Other   pgtype.Text `type:"other"`
}

type Foo struct {
	ID      pgtype.Text `source:"foo" type:"id"`
	OtherID pgtype.Text `type:"other_id"`
	About   pgtype.Text `type:"about"`
	Other   pgtype.Text `type:"this"`
}
