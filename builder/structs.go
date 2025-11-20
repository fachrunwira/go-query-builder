package builder

import "database/sql"

type queryStruct struct {
	db          *sql.DB
	tableName   string
	tableAlias  string
	fields      []string
	whereClause []string
	whereArgs   []interface{}

	joins    []string
	joinArgs []interface{}

	manipulateType string
	manipulateArgs Rows

	grouping []string

	ordering []string

	limit  uint
	offset uint

	errors error
}
