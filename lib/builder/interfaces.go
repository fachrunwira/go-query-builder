package builder

// Row represents a single database record, where each key is the column name
// and the value is the associated field value.
type Row map[string]interface{}

// Rows represents a collection of Row values, typically used to return
// multiple records from a query or result set.
type Rows []Row

// initialStage this is something.
type initialStage interface {
	Table(table string, alias ...string) insertOrQueryingStage
	InsertRaw(query string, bindings ...interface{}) error
	UpdateRaw(query string, bindings ...interface{}) error
	DeleteRaw(query string, bindings ...interface{}) error
}

type insertOrQueryingStage interface {
	Insert(data Rows) generateCreateQuery
	manipulateData
	errorInterface
}

type manipulateData interface {
	Select(columns ...string) manipulateData

	Update(data Row) generateCreateQuery
	Delete() generateCreateQuery

	WhereRaw(query string, bindings ...interface{}) manipulateData
	Where(column string, values interface{}) manipulateData
	WhereNot(column string, values interface{}) manipulateData
	WhereIn(column string, values []interface{}) manipulateData
	WhereNotIn(column string, values []interface{}) manipulateData

	OrWhereRaw(query string, bindings ...interface{}) manipulateData
	OrWhere(column string, values interface{}) manipulateData
	OrWhereNot(column string, values interface{}) manipulateData
	OrWhereIn(column string, values []interface{}) manipulateData
	OrWhereNotIn(column string, values []interface{}) manipulateData

	Join(tableJoin, table_1, table_2 string) manipulateData

	generateSelectQuery
	errorInterface
}

type generateSelectQuery interface {
	ToSql() (*string, error)
	errorInterface
}

type generateCreateQuery interface {
	ToSql() (*string, error)
	Save() error
	errorInterface
}

type errorInterface interface {
	Error() error
}
