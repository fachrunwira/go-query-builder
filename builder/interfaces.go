package builder

// initialStage this is something.
type initialStage interface {
	Table(table string, alias ...string) insertOrQueryingStage
	InsertRaw(query string, bindings ...any) error
	UpdateRaw(query string, bindings ...any) error
	DeleteRaw(query string, bindings ...any) error
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
	limiterInterface
	generateSelectQuery
}

type selectInterface interface {
	Select(columns ...string) manipulateData
}

type whereInterface interface {
	WhereRaw(query string, bindings ...any) manipulateOrQuerying
	Where(column string, values any) manipulateOrQuerying
	WhereNot(column string, values any) manipulateOrQuerying
	WhereIn(column string, values []any) manipulateOrQuerying
	WhereNotIn(column string, values []any) manipulateOrQuerying

	OrWhereRaw(query string, bindings ...any) manipulateOrQuerying
	OrWhere(column string, values any) manipulateOrQuerying
	OrWhereNot(column string, values any) manipulateOrQuerying
	OrWhereIn(column string, values []any) manipulateOrQuerying
	OrWhereNotIn(column string, values []any) manipulateOrQuerying
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
	generateSelectQuery
}

type joinInterface interface {
	Join(tableJoin, table_1, operator, table_2 string) manipulateData
	JoinWhere(tableJoin, table_1, operator string, bindings ...any) manipulateData

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

type limiterInterface interface {
	Limit(limit uint) offsetInterface
	generateSelectQuery
}

type offsetInterface interface {
	Offset(offset uint) generateSelectQuery
	generateSelectQuery
}

type manipulateData interface {
	selectInterface
	whereInterface
	joinInterface
	groupingInterface
	orderingInterface
	generateSelectQuery
	limiterInterface
}

type generateSelectQuery interface {
	ToSql() (string, error)
	Get() (Rows, error)
	First() (Row, error)
}

type generateCreateQuery interface {
	ToSql() (string, error)
	Save() error
}
