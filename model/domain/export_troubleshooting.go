package domain

import "database/sql"

type Export struct {
	NamaCustomer string
	Biaya int
	TglMasuk string
	ChangeComponent sql.NullString
}
