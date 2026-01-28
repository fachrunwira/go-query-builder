package joinbuilder

import (
	"fmt"
	"strings"

	"github.com/fachrunwira/go-query-builder/clauseoperators"
)

func (q *queryStruct) OrWhereNot(column string, operator clauseoperators.Operators, args ...any) JoinQuery {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR NOT %s = ?", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("ON NOT %s = ?", column))
	}

	q.whereArgs = append(q.whereArgs, args)
	return q
}

func (q *queryStruct) OrWhereNotBetween(column string, args ...any) JoinQuery {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR NOT %s BETWEEN ? AND ?", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("ON NOT %s BETWEEN ? AND ?", column))
	}

	if len(args) > 0 {
		q.whereArgs = append(q.whereArgs, args...)
	} else {
		q.errors = fmt.Errorf("arguments must at least have one value")
	}
	return q
}

func (q *queryStruct) OrWhereNotIn(column string, args ...any) JoinQuery {
	if len(args) == 0 {
		q.errors = fmt.Errorf("arguments must have at least one value")
		return q
	}

	placeholder := "(" + strings.Repeat("?,", len(args)-1) + "?)"

	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR NOT %s IN %s", column, placeholder))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("ON NOT %s IN %s", column, placeholder))
	}

	q.whereArgs = append(q.whereArgs, args...)
	return q
}

func (q *queryStruct) OrWhereNotNull(column string) JoinQuery {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR %s IS NOT NULL", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("ON %s IS NOT NULL", column))
	}

	return q
}
