package persistence_utils

import (
	"database/sql"
	"errors"
	"fmt"
)

func HandlePqError(err error) error {
	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		return errors.New("No rows returned")
	default:
		return fmt.Errorf("unknown error: %s", err.Error())
	}
}
