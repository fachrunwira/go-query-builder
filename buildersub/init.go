package buildersub

import (
	"fmt"
	"strings"
)

func Make() initialStage {
	return &queryStruct{}
}

func Build(subquery SubQuery) (string, []any, error) {
	sb, ok := subquery.(*queryStruct)
	if !ok {
		return "", nil, fmt.Errorf("invalid type for sub query")
	}

	if sb.tableName == "" {
		return "", nil, fmt.Errorf("sub query table name is empty")
	}
	query := "SELECT"
	var args []any

	if len(sb.fields) > 0 {
		query += fmt.Sprintf(" %s", strings.Join(sb.fields, ","))
	} else {
		query += " *"
	}

	query += fmt.Sprintf(" FROM %s", sb.tableName)
	if sb.tableAlias != "" {
		query += fmt.Sprintf(" AS %s", sb.tableAlias)
	}

	if len(sb.joins) > 0 {
		query += fmt.Sprintf(" %s", strings.Join(sb.joins, " "))
		args = append(args, sb.joinArgs...)
	}

	if len(sb.whereClause) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(sb.whereClause, " "))
		args = append(args, sb.whereArgs...)
	}

	if len(sb.grouping) > 0 {
		query += fmt.Sprintf(" GROUP BY %s", strings.Join(sb.grouping, ","))
	}

	if len(sb.ordering) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(sb.ordering, ","))
	}

	if sb.limits > 0 {
		query += fmt.Sprintf(" LIMIT %d", sb.limits)
	}

	return query, args, nil
}
