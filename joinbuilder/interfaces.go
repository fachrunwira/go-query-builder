package joinbuilder

import (
	"github.com/fachrunwira/go-query-builder/clauseoperators"
)

type clauseInterface interface {
	On(clause_1 string, operator clauseoperators.Operators, clause_2 string) JoinQuery
	Or(clause_1 string, operator clauseoperators.Operators, clause_2 string) JoinQuery
}

type whereInterface interface {
	Where(column string, operator clauseoperators.Operators, args ...any) JoinQuery
	WhereBetween(column string, args ...any) JoinQuery
	WhereIn(column string, args ...any) JoinQuery
	WhereNull(column string) JoinQuery
	WhereRaw(query string, args ...any) JoinQuery

	OrWhere(column string, operator clauseoperators.Operators, args ...any) JoinQuery
	OrWhereBetween(column string, args ...any) JoinQuery
	OrWhereIn(column string, args ...any) JoinQuery
	OrWhereNull(column string) JoinQuery
	OrWhereRaw(query string, args ...any) JoinQuery

	WhereNot(column string, operator clauseoperators.Operators, args ...any) JoinQuery
	WhereNotBetween(column string, args ...any) JoinQuery
	WhereNotIn(column string, args ...any) JoinQuery
	WhereNotNull(column string) JoinQuery

	OrWhereNot(column string, operator clauseoperators.Operators, args ...any) JoinQuery
	OrWhereNotBetween(column string, args ...any) JoinQuery
	OrWhereNotIn(column string, args ...any) JoinQuery
	OrWhereNotNull(column string) JoinQuery
}

type JoinQuery interface {
	clauseInterface
	whereInterface
}
