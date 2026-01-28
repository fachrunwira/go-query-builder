package joinbuilder

import (
	"fmt"
	"strings"

	"github.com/fachrunwira/go-query-builder/clauseoperators"
)

func (q *queryStruct) OrWhere(column string, operator clauseoperators.Operators, args ...any) JoinQuery {
	op, placeholder, err := getClauseOperator(operator, args...)
	if err != nil {
		q.errors = fmt.Errorf("error in join builder: %w", err)
		return q
	}

	if placeholder != "" {
		if len(q.whereClause) > 0 {
			q.whereClause = append(q.whereClause, fmt.Sprintf("OR %s %s %s", column, op, placeholder))
		} else {
			q.whereClause = append(q.whereClause, fmt.Sprintf("ON %s %s %s", column, op, placeholder))
		}
	} else {
		if len(q.whereClause) > 0 {
			q.whereClause = append(q.whereClause, fmt.Sprintf("OR %s %s ?", column, op))
		} else {
			q.whereClause = append(q.whereClause, fmt.Sprintf("ON %s %s ?", column, op))
		}
	}

	q.whereArgs = append(q.whereArgs, args)
	return q
}

func (q *queryStruct) OrWhereBetween(column string, args ...any) JoinQuery {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR %s BETWEEN ? AND ?", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("ON %s BETWEEN ? AND ?", column))
	}

	if len(args) > 0 {
		q.whereArgs = append(q.whereArgs, args...)
	} else {
		q.errors = fmt.Errorf("arguments must at least have one value")
	}
	return q
}
func (q *queryStruct) OrWhereIn(column string, args ...any) JoinQuery {
	if len(args) == 0 {
		q.errors = fmt.Errorf("arguments must have at least one value")
		return q
	}

	placeholder := "(" + strings.Repeat("?,", len(args)-1) + "?)"

	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR %s IN %s", column, placeholder))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("ON %s IN %s", column, placeholder))
	}

	q.whereArgs = append(q.whereArgs, args...)
	return q
}

func (q *queryStruct) OrWhereNull(column string) JoinQuery {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR %s IS NULL", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("ON %s IS NULL", column))
	}

	return q
}

func (q *queryStruct) OrWhereRaw(query string, args ...any) JoinQuery {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("OR %s", query))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("ON %s", query))
	}

	if len(args) > 0 {
		q.whereArgs = append(q.whereArgs, args...)
	}

	return q
}
