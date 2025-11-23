package builder

import (
	"context"
	"database/sql"
)

type queryStruct struct {
	db          *sql.DB
	useContext  bool
	ctx         context.Context
	tableName   string
	tableAlias  string
	fields      []string
	whereClause []string
	whereArgs   []any

	joins    []string
	joinArgs []any

	manipulateType string
	manipulateArgs Rows

	grouping []string

	ordering []string

	limit  uint
	offset uint

	errors error
}
