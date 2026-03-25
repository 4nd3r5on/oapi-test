package postgres

import (
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/4nd3rs0n/oapi-test/pkg/errs"
)

// mapErr translates postgres-specific errors to domain sentinel errors.
func mapErr(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return errs.SafeMsg(errs.Mark(err, errs.ErrNotFound), "not found")
	}
	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
		switch pgErr.Code {
		case "23505": // unique_violation
			errMsg := strings.Builder{}
			field, ok := fieldFromConstraint(pgErr.ConstraintName)
			errMsg.WriteString("entity ")
			if ok {
				errMsg.WriteString("with conflicting ")
				errMsg.WriteString(field)
				errMsg.WriteByte(' ')
			}
			errMsg.WriteString("exists")
			return errs.SafeMsg(errs.Mark(err, errs.ErrExists), errMsg.String())
		}
	}
	return err
}

// fieldFromConstraint extracts a human-readable field name from a postgres
// constraint name following the <table>_<field>_<suffix> naming convention.
// e.g. "users_username_idx" → "username"
func fieldFromConstraint(constraint string) (string, bool) {
	first := strings.Index(constraint, "_")
	last := strings.LastIndex(constraint, "_")
	if first == last {
		return "", false
	}
	return constraint[first+1 : last], true
}
