package buildersub

import (
	"fmt"
	"strings"

	"github.com/fachrunwira/go-query-builder/clauseoperators"
)

func (qs *queryStruct) OrWhere(column string, operator clauseoperators.Operators, args ...any) SubQuery {
	op, placeholder, err := getClauseOperator(operator, args...)
	if err != nil {
		qs.err = fmt.Errorf("error in sub query: %w", err)
		return qs
	}

	if placeholder != "" {
		if len(qs.whereClause) > 0 {
			qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR %s %s %s", column, op, placeholder))
		} else {
			qs.whereClause = append(qs.whereClause, fmt.Sprintf("%s %s %s", column, op, placeholder))
		}
	} else {
		if len(qs.whereClause) > 0 {
			qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR %s %s ?", column, op))
		} else {
			qs.whereClause = append(qs.whereClause, fmt.Sprintf("%s %s ?", column, op))
		}
	}

	qs.whereArgs = append(qs.whereArgs, args...)
	return qs
}

func (qs *queryStruct) OrWhereRaw(query string, args ...any) SubQuery {
	if len(qs.whereClause) > 0 {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR %s", query))
	} else {
		qs.whereClause = append(qs.whereClause, query)
	}

	if len(args) > 0 {
		qs.whereArgs = append(qs.whereArgs, args...)
	}

	return qs
}

func (qs *queryStruct) OrWhereIn(column string, args ...any) SubQuery {
	if len(args) == 0 {
		qs.err = fmt.Errorf("arguments must have at least one value")
		return qs
	}

	placeholder := "(" + strings.Repeat("?,", len(args)-1) + "?)"

	if len(qs.whereClause) > 0 {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR %s IN %s", column, placeholder))
	} else {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("%s IN %s", column, placeholder))
	}

	qs.whereArgs = append(qs.whereArgs, args...)
	return qs
}

func (qs *queryStruct) OrWhereBetween(column string, args ...any) SubQuery {
	if len(qs.whereClause) > 0 {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR %s BETWEEN ? AND ?", column))
	} else {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("%s BETWEEN ? AND ?", column))
	}

	if len(args) > 0 {
		qs.whereArgs = append(qs.whereArgs, args...)
	} else {
		qs.err = fmt.Errorf("arguments must at least have one value")
	}

	return qs
}

func (qs *queryStruct) OrWhereNull(column string) SubQuery {
	if len(qs.whereClause) > 0 {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR %s IS NULL", column))
	} else {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("%s IS NULL", column))
	}

	return qs
}
