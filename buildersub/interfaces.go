package buildersub

import (
	"github.com/fachrunwira/go-query-builder/clauseoperators"
	"github.com/fachrunwira/go-query-builder/joinbuilder"
)

type initialStage interface {
	Table(table string, alias ...string) SubQuery
}

type selectInterface interface {
	Select(columns ...string) SubQuery
}

type whereInterface interface {
	Where(column string, operator clauseoperators.Operators, args ...any) SubQuery
	WhereRaw(query string, args ...any) SubQuery
	WhereIn(column string, args ...any) SubQuery
	WhereBetween(column string, args ...any) SubQuery
	WhereNull(column string) SubQuery

	WhereNot(column string, operator clauseoperators.Operators, args ...any) SubQuery
	WhereNotIn(column string, args ...any) SubQuery
	WhereNotBetween(column string, args ...any) SubQuery
	WhereNotNull(column string) SubQuery

	OrWhere(column string, operator clauseoperators.Operators, args ...any) SubQuery
	OrWhereRaw(query string, args ...any) SubQuery
	OrWhereIn(column string, args ...any) SubQuery
	OrWhereBetween(column string, args ...any) SubQuery
	OrWhereNull(column string) SubQuery

	OrWhereNot(column string, operator clauseoperators.Operators, args ...any) SubQuery
	OrWhereNotIn(column string, args ...any) SubQuery
	OrWhereNotBetween(column string, args ...any) SubQuery
	OrWhereNotNull(column string) SubQuery
}

type groupingInterface interface {
	GroupBy(columns ...string) SubQuery
}

type orderingInterface interface {
	OrderByAsc(column string) SubQuery
	OrderByDesc(column string) SubQuery
}

type joinInterface interface {
	Join(callback func() joinbuilder.JoinQuery) SubQuery
	LeftJoin(callback func() joinbuilder.JoinQuery) SubQuery
	RightJoin(callback func() joinbuilder.JoinQuery) SubQuery

	CrossJoin(tableJoin string) SubQuery
}

type limiterInterface interface {
	Limit(limit uint) SubQuery
}

type SubQuery interface {
	selectInterface
	joinInterface
	whereInterface
	groupingInterface
	orderingInterface
	limiterInterface
}
