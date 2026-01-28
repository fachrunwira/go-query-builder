package buildersub

type queryStruct struct {
	tableName   string
	tableAlias  string
	fields      []string
	whereClause []string
	whereArgs   []any
	ordering    []string
	grouping    []string
	limits      uint

	joins    []string
	joinArgs []any

	err error
}
