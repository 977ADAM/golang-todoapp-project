package corepostgrespool

import "errors"

var (
	ErrNoRows = errors.New("no rows")
	ErrViolatesForeignKey = errors.New("violates foreign key constraint")
	ErrUnknown = errors.New("unknown error")
)
