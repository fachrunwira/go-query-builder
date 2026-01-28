package builder

import (
	"fmt"
	"strings"

	"github.com/fachrunwira/go-query-builder/buildersub"
	"github.com/fachrunwira/go-query-builder/clauseoperators"
)

func (q *queryStruct) Where(column string, operator clauseoperators.Operators, args ...any) manipulateOrQuerying {
	op, placeholder, err := getClauseOperator(operator, args...)
	if err != nil {
		q.errors = fmt.Errorf("error in query builder: %w", err)
		return q
	}

	if placeholder != "" {
		if len(q.whereClause) > 0 {
			q.whereClause = append(q.whereClause, fmt.Sprintf("AND %s %s %s", column, op, placeholder))
		} else {
			q.whereClause = append(q.whereClause, fmt.Sprintf("%s %s %s", column, op, placeholder))
		}
	} else {
		if len(q.whereClause) > 0 {
			q.whereClause = append(q.whereClause, fmt.Sprintf("AND %s %s ?", column, op))
		} else {
			q.whereClause = append(q.whereClause, fmt.Sprintf("%s %s ?", column, op))
		}
	}

	q.whereArgs = append(q.whereArgs, args...)
	return q
}

func (q *queryStruct) WhereBetween(column string, args ...any) manipulateOrQuerying {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("AND %s BETWEEN ? AND ?", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("%s BETWEEN ? AND ?", column))
	}

	if len(args) > 0 {
		q.whereArgs = append(q.whereArgs, args...)
	} else {
		q.errors = fmt.Errorf("arguments must at least have one value")
	}
	return q
}

func (q *queryStruct) WhereIn(column string, args ...any) manipulateOrQuerying {
	if len(args) == 0 {
		q.errors = fmt.Errorf("arguments must have at least one value")
		return q
	}

	placeholder := "(" + strings.Repeat("?,", len(args)-1) + "?)"

	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("AND %s IN %s", column, placeholder))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("%s IN %s", column, placeholder))
	}

	q.whereArgs = append(q.whereArgs, args...)
	return q
}

func (q *queryStruct) WhereExists(callback func() buildersub.SubQuery) manipulateOrQuerying {
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
		q.whereClause = append(q.whereClause, fmt.Sprintf("AND EXISTS (%s)", subquery))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("EXISTS (%s)", subquery))
	}

	q.whereArgs = append(q.whereArgs, args...)

	return q
}

func (q *queryStruct) WhereNull(column string) manipulateOrQuerying {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("AND %s IS NULL", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("%s IS NULL", column))
	}

	return q
}

func (q *queryStruct) WhereSub(column string, operator clauseoperators.Operators, callback func() buildersub.SubQuery) manipulateOrQuerying {
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

	fmt.Println(subquery)

	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("AND %s %s (%s)", column, op, subquery))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("%s %s (%s)", column, op, subquery))
	}

	q.whereArgs = append(q.whereArgs, args...)
	return q
}

func (q *queryStruct) WhereRaw(query string, args ...any) manipulateOrQuerying {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf("AND %s", query))
	} else {
		q.whereClause = append(q.whereClause, query)
	}

	q.whereArgs = append(q.whereArgs, args...)

	return q
}
