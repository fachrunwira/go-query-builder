package builder

import "database/sql"

type queryStruct struct {
	db          *sql.DB
	tableName   string
	tableAlias  string
	fields      []string
	whereClause []string
	whereArgs   []interface{}

	joins []string

	manipulateType string
	manipulateArgs Rows

	errors error
}
