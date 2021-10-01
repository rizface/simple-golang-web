package domain

import "database/sql"

type Component struct {
	Id sql.NullInt64
	Component string
	Include sql.NullInt64
}
