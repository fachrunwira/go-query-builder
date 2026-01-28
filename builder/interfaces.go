package builder

import (
	"github.com/fachrunwira/go-query-builder/buildersub"
	"github.com/fachrunwira/go-query-builder/clauseoperators"
	"github.com/fachrunwira/go-query-builder/joinbuilder"
)

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
	Where(column string, operator clauseoperators.Operators, args ...any) manipulateOrQuerying
	WhereIn(column string, args ...any) manipulateOrQuerying
	WhereExists(callback func() buildersub.SubQuery) manipulateOrQuerying
	WhereBetween(column string, args ...any) manipulateOrQuerying
	WhereNull(column string) manipulateOrQuerying
	WhereSub(column string, operator clauseoperators.Operators, callback func() buildersub.SubQuery) manipulateOrQuerying
	WhereRaw(query string, args ...any) manipulateOrQuerying

	WhereNot(column string, operator clauseoperators.Operators, args ...any) manipulateOrQuerying
	WhereNotIn(column string, args ...any) manipulateOrQuerying
	WhereNotExists(callback func() buildersub.SubQuery) manipulateOrQuerying
	WhereNotBetween(column string, args ...any) manipulateOrQuerying
	WhereNotNull(column string) manipulateOrQuerying
	WhereNotSub(column string, operator clauseoperators.Operators, callback func() buildersub.SubQuery) manipulateOrQuerying

	OrWhere(column string, operator clauseoperators.Operators, args ...any) manipulateOrQuerying
	OrWhereIn(column string, args ...any) manipulateOrQuerying
	OrWhereExists(callback func() buildersub.SubQuery) manipulateOrQuerying
	OrWhereBetween(column string, args ...any) manipulateOrQuerying
	OrWhereNull(column string) manipulateOrQuerying
	OrWhereSub(column string, operator clauseoperators.Operators, callback func() buildersub.SubQuery) manipulateOrQuerying
	OrWhereRaw(query string, args ...any) manipulateOrQuerying

	OrWhereNot(column string, operator clauseoperators.Operators, args ...any) manipulateOrQuerying
	OrWhereNotIn(column string, args ...any) manipulateOrQuerying
	OrWhereNotExists(callback func() buildersub.SubQuery) manipulateOrQuerying
	OrWhereNotBetween(column string, args ...any) manipulateOrQuerying
	OrWhereNotNull(column string) manipulateOrQuerying
	OrWhereNotSub(column string, operator clauseoperators.Operators, callback func() buildersub.SubQuery) manipulateOrQuerying
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
	limiterInterface
}

type joinInterface interface {
	Join(tableJoin string, callback func() joinbuilder.JoinQuery) manipulateData
	LeftJoin(tableJoin string, callback func() joinbuilder.JoinQuery) manipulateData
	RightJoin(tableJoin string, callback func() joinbuilder.JoinQuery) manipulateData

	JoinSub(query func() buildersub.SubQuery, as string, clause func() joinbuilder.JoinQuery) manipulateData
	LeftJoinSub(query func() buildersub.SubQuery, as string, clause func() joinbuilder.JoinQuery) manipulateData
	RightJoinSub(query func() buildersub.SubQuery, as string, clause func() joinbuilder.JoinQuery) manipulateData

	CrossJoin(tableJoin string) manipulateData
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
	ToRaw() (string, error)
	Get() (Rows, error)
	First() (Row, error)
}

type generateCreateQuery interface {
	ToSql() (string, error)
	ToRaw() (string, error)
	Save() error
}
