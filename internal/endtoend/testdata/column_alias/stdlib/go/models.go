// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querytest

import (
	"time"
)

type User struct {
	ID        int32
	Fname     string
	Lname     string
	Email     string
	EncPasswd string
	CreatedAt time.Time
}
