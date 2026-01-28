package builder

import (
	"fmt"
	"strings"

	"github.com/fachrunwira/go-query-builder/buildersub"
	"github.com/fachrunwira/go-query-builder/clauseoperators"
)

func (q *queryStruct) OrWhereNot(column string, operator clauseoperators.Operators, args ...any) manipulateOrQuerying {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR NOT %s = ?", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("NOT %s = ?", column))
	}

	q.whereArgs = append(q.whereArgs, args...)
	return q
}

func (q *queryStruct) OrWhereNotBetween(column string, args ...any) manipulateOrQuerying {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR NOT %s BETWEEN ? AND ?", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("NOT %s BETWEEN ? AND ?", column))
	}

	if len(args) > 0 {
		q.whereArgs = append(q.whereArgs, args...)
	} else {
		q.errors = fmt.Errorf("arguments must at least have one value")
	}
	return q
}

func (q *queryStruct) OrWhereNotIn(column string, args ...any) manipulateOrQuerying {
	if len(args) == 0 {
		q.errors = fmt.Errorf("arguments must have at least one value")
		return q
	}

	placeholder := "(" + strings.Repeat("?,", len(args)-1) + "?)"

	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR NOT %s IN %s", column, placeholder))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("NOT %s IN %s", column, placeholder))
	}

	q.whereArgs = append(q.whereArgs, args...)
	return q
}

func (q *queryStruct) OrWhereNotExists(callback func() buildersub.SubQuery) manipulateOrQuerying {
	cb := callback()
	if cb == nil {
		q.errors = fmt.Errorf("no subquery value detected")
		return q
	}

	subquery, args, err := buildersub.Build(cb)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make sub query: %w", err)
		return q
	}

	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR NOT EXISTS (%s)", subquery))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("NOT EXISTS (%s)", subquery))
	}

	q.whereArgs = append(q.whereArgs, args...)

	return q
}

func (q *queryStruct) OrWhereNotNull(column string) manipulateOrQuerying {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR %s IS NOT NULL", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("%s IS NOT NULL", column))
	}

	return q
}

func (q *queryStruct) OrWhereNotSub(column string, operator clauseoperators.Operators, callback func() buildersub.SubQuery) manipulateOrQuerying {
	op, err := getClauseOperatorSub(operator)
	if err != nil {
		q.errors = fmt.Errorf("error in query: %w", err)
		return q
	}

	cb := callback()
	if cb == nil {
		q.errors = fmt.Errorf("no sub query detected")
		return q
	}

	subquery, args, err := buildersub.Build(cb)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make sub query: %w", err)
		return q
	}

	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR NOT %s %s (%s)", column, op, subquery))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("NOT %s %s (%s)", column, op, subquery))
	}

	q.whereArgs = append(q.whereArgs, args...)
	return q
}
