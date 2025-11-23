package builder

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strings"
)

func trxContext(ctx context.Context, db *sql.DB, query string, args ...any) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}

func trx(db *sql.DB, query string, args ...any) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec(query, args...)

	if err != nil {
		return err
	}

	err = tx.Commit()

	return err
}

func clearStrings(data []string) []string {
	n := 0
	for _, v := range data {
		v = strings.TrimSpace(v)
		if v != "" {
			data[n] = v
			n++
		}
	}

	data = data[:n]
	return data
}

func initManipulateData(q *queryStruct) (string, []any) {
	switch q.manipulateType {
	case "insert":
		keys := make([]string, 0, len(q.manipulateArgs[0]))
		for k := range q.manipulateArgs[0] {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		placeholder := "(" + strings.Repeat("?,", len(keys)-1) + "?)"

		valueGroup := make([]string, 0, len(keys))
		values := make([]any, 0, len(keys)*len(q.manipulateArgs))

		for _, row := range q.manipulateArgs {
			valueGroup = append(valueGroup, placeholder)
			for _, k := range keys {
				values = append(values, row[k])
			}
		}

		return fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", q.tableName, strings.Join(keys, ","), strings.Join(valueGroup, ",")), values
	case "update":
		keys := make([]string, 0, len(q.manipulateArgs[0]))
		values := make([]any, 0, len(q.manipulateArgs[0]))
		for k, v := range q.manipulateArgs[0] {
			keys = append(keys, fmt.Sprintf("%s = ?", k))
			values = append(values, v)
		}

		whereClauses := "WHERE" + strings.Join(q.whereClause, "")
		values = append(values, q.whereArgs...)

		return fmt.Sprintf("UPDATE %s SET %s %s;", q.tableName, strings.Join(keys, ", "), whereClauses), values
	case "delete":
		return fmt.Sprintf("DELETE FROM %s WHERE %s;", q.tableName, strings.Join(q.whereClause, "")), q.whereArgs
	}

	return "", []any{}
}

func printManipulateData(q *queryStruct) string {
	query, _ := initManipulateData(q)
	return query
}

func initQuery(q *queryStruct) (string, []any) {
	fields := "*"
	if len(q.fields) > 0 {
		fields = strings.Join(q.fields, ",")
	}

	var args []any
	query := fmt.Sprintf("SELECT %s FROM %s", fields, q.tableName)
	if q.tableAlias != "" {
		query += " AS " + q.tableAlias
	}

	if len(q.joins) > 0 {
		query += fmt.Sprintf(" %s", strings.Join(q.joins, " "))
		args = append(args, q.joinArgs...)
	}

	if len(q.whereClause) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(q.whereClause, " "))
		args = append(args, q.whereArgs...)
	}

	if len(q.grouping) > 0 {
		query += fmt.Sprintf(" GROUP BY %s", strings.Join(q.grouping, ","))
	}

	if len(q.ordering) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(q.ordering, ","))
	}

	return query, args
}

func printQuery(q *queryStruct) string {
	query, _ := initQuery(q)

	if q.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", q.limit)
	}

	if q.offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", q.offset)
	}

	return query
}

func convertRow(data any) (Row, error) {
	switch x := data.(type) {
	case Row:
		return x, nil
	case map[string]any:
		return Row(x), nil
	}
	return nil, fmt.Errorf("unsupported type")
}

func convertRows(data any) (Rows, error) {
	switch x := data.(type) {
	case Rows:
		return x, nil
	case []map[string]any:
		rows := make(Rows, len(x))
		for i := range x {
			rows[i] = Row(x[i])
		}
		return rows, nil
	case Row:
		return Rows{x}, nil
	case map[string]any:
		return Rows{Row(x)}, nil
	}
	return nil, fmt.Errorf("unsupported type")
}
