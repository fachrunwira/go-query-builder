package buildersub

import (
	"fmt"
	"strings"

	"github.com/fachrunwira/go-query-builder/clauseoperators"
)

func (qs *queryStruct) OrWhereNot(column string, operator clauseoperators.Operators, args ...any) SubQuery {
	op, placeholder, err := getClauseOperator(operator, args...)
	if err != nil {
		qs.err = fmt.Errorf("error in sub query: %w", err)
		return qs
	}

	if placeholder != "" {
		if len(qs.whereClause) > 0 {
			qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR NOT %s %s %s", column, op, placeholder))
		} else {
			qs.whereClause = append(qs.whereClause, fmt.Sprintf("NOT %s %s %s", column, op, placeholder))
		}
	} else {
		if len(qs.whereClause) > 0 {
			qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR NOT %s %s ?", column, op))
		} else {
			qs.whereClause = append(qs.whereClause, fmt.Sprintf("NOT %s %s ?", column, op))
		}
	}

	qs.whereArgs = append(qs.whereArgs, args...)
	return qs
}

func (qs *queryStruct) OrWhereNotIn(column string, args ...any) SubQuery {
	if len(args) == 0 {
		qs.err = fmt.Errorf("arguments must have at least one value")
		return qs
	}

	placeholder := "(" + strings.Repeat("?,", len(args)-1) + "?)"

	if len(qs.whereClause) > 0 {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR NOT %s IN %s", column, placeholder))
	} else {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("NOT %s IN %s", column, placeholder))
	}

	qs.whereArgs = append(qs.whereArgs, args...)
	return qs
}

func (qs *queryStruct) OrWhereNotBetween(column string, args ...any) SubQuery {
	if len(qs.whereClause) > 0 {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR NOT %s BETWEEN ? AND ?", column))
	} else {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("NOT %s BETWEEN ? AND ?", column))
	}

	if len(args) > 0 {
		qs.whereArgs = append(qs.whereArgs, args...)
	} else {
		qs.err = fmt.Errorf("arguments must at least have one value")
	}

	return qs
}

func (qs *queryStruct) OrWhereNotNull(column string) SubQuery {
	if len(qs.whereClause) > 0 {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("OR NOT %s IS NULL", column))
	} else {
		qs.whereClause = append(qs.whereClause, fmt.Sprintf("NOT %s IS NULL", column))
	}

	return qs
}
