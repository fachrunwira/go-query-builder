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
	Insert(data any) generateCreateQuery
	combinedInterface1
}

type combinedInterface1 interface {
	selectInterface
	whereInterface
	joinInterface
	groupingInterface
	orderingInterface
}

type selectInterface interface {
	Select(columns ...string) manipulateData
}

type whereInterface interface {
	WhereRaw(query string, bindings ...interface{}) manipulateOrQuerying
	Where(column string, values interface{}) manipulateOrQuerying
	WhereNot(column string, values interface{}) manipulateOrQuerying
	WhereIn(column string, values []interface{}) manipulateOrQuerying
	WhereNotIn(column string, values []interface{}) manipulateOrQuerying

	OrWhereRaw(query string, bindings ...interface{}) manipulateOrQuerying
	OrWhere(column string, values interface{}) manipulateOrQuerying
	OrWhereNot(column string, values interface{}) manipulateOrQuerying
	OrWhereIn(column string, values []interface{}) manipulateOrQuerying
	OrWhereNotIn(column string, values []interface{}) manipulateOrQuerying
}

type updateOrDeleteInterface interface {
	Update(data any) generateCreateQuery
	Delete() generateCreateQuery
}

type manipulateOrQuerying interface {
	updateOrDeleteInterface
	selectInterface
	whereInterface
	joinInterface
	groupingInterface
	orderingInterface
}

type joinInterface interface {
	Join(tableJoin, table_1, operator, table_2 string) manipulateData
	JoinWhere(tableJoin, table_1, operator string, bindings ...interface{}) manipulateData

	LeftJoin(tableJoin, table_1, operator, table_2 string) manipulateData
	RightJoin(tableJoin, table_1, operator, table_2 string) manipulateData
}

type groupingInterface interface {
	GroupBy(columns ...string) manipulateData
}

type orderingInterface interface {
	OrderByAsc(column string) manipulateData
	OrderByDesc(column string) manipulateData
}

type manipulateData interface {
	selectInterface
	whereInterface
	joinInterface
	groupingInterface
	orderingInterface
	generateSelectQuery
}

type generateSelectQuery interface {
	ToSql() (string, error)
}

type generateCreateQuery interface {
	ToSql() (string, error)
	Save() error
}
