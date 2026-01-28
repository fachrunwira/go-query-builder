package joinbuilder

type queryStruct struct {
	whereClause []string
	whereArgs   []any

	errors error
}
