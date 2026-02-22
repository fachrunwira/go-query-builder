package builder

import (
	"database/sql"
	"fmt"
	"strings"
)

func (q *queryStruct) Get() (Rows, error) {
	if q.errors != nil {
		return nil, q.errors
	}

	query, args := initQuery(q)

	if q.limit > 0 {
		query += " LIMIT ?"
		args = append(args, q.limit)
	}

	if q.offset > 0 {
		query += " OFFSET ?"
		args = append(args, q.offset)
	}

	var (
		rows *sql.Rows
		err  error
	)

	if q.useContext {
		rows, err = q.db.QueryContext(q.ctx, query, args...)
	} else {
		rows, err = q.db.Query(query, args...)
	}

	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	result := make(Rows, 0, 100)
	for rows.Next() {
		values := make([]any, len(columns))
		valuesPtrs := make([]any, len(columns))

		for i := range columns {
			valuesPtrs[i] = &values[i]
		}

		if err = rows.Scan(valuesPtrs...); err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		row := make(Row)
		for i, col := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				row[col] = string(b)
			} else {
				row[col] = val
			}
		}
		result = append(result, row)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed: %w", err)
	}

	return result, nil
}

func (q *queryStruct) First() (Row, error) {
	if q.errors != nil {
		return nil, q.errors
	}

	query, args := initQuery(q)
	query += " LIMIT 1;"

	var (
		rows *sql.Rows
		err  error
	)

	if q.useContext {
		rows, err = q.db.QueryContext(q.ctx, query, args...)
	} else {
		rows, err = q.db.Query(query, args...)
	}

	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("no data found: %w", sql.ErrNoRows)
	}

	column, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	values := make([]any, len(column))
	valuesPtrs := make([]any, len(column))
	result := Row{}

	for i := range column {
		valuesPtrs[i] = &values[i]
	}

	if err = rows.Scan(valuesPtrs...); err != nil {
		return nil, fmt.Errorf("row scan failed: %w", err)
	}

	for i, col := range column {
		val := values[i]
		if b, ok := val.([]byte); ok {
			result[col] = string(b)
		} else {
			result[col] = val
		}
	}

	return result, nil
}

func (q *queryStruct) Exists() (bool, error) {
	if q.errors != nil {
		return false, q.errors
	}

	var args []any
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s", q.tableName)
	if q.tableAlias != "" {
		query += fmt.Sprintf(" AS %s", q.tableAlias)
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

	query += ")"

	var (
		row    *sql.Row
		exists bool
		err    error
	)

	if q.useContext {
		row = q.db.QueryRowContext(q.ctx, query, args...)
	} else {
		row = q.db.QueryRow(query, args...)
	}

	err = row.Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("query execution failed: %w", err)
	}

	return exists, nil
}

func (q *queryStruct) Count() (int64, error) {
	if q.errors != nil {
		return 0, q.errors
	}

	var (
		args  []any
		query string
		row   *sql.Row
		total int64
		err   error
	)

	if len(q.grouping) > 0 {
		fields := "*"
		if len(q.fields) > 0 {
			fields = strings.Join(q.fields, ",")
		}

		query = fmt.Sprintf("SELECT COUNT(*) FROM (SELECT %s FROM %s", fields, q.tableName)
		if q.tableAlias != "" {
			query += fmt.Sprintf(" AS %s", q.tableAlias)
		}

		if len(q.joins) > 0 {
			query += fmt.Sprintf(" %s", strings.Join(q.joins, " "))
			args = append(args, q.joinArgs...)
		}

		if len(q.whereClause) > 0 {
			query += fmt.Sprintf(" WHERE %s", strings.Join(q.whereClause, " "))
			args = append(args, q.whereArgs...)
		}

		query += fmt.Sprintf(" GROUP BY %s) AS total_row", strings.Join(q.grouping, ","))
	} else {
		query = fmt.Sprintf("SELECT COUNT(*) FROM %s", q.tableName)
		if q.tableAlias != "" {
			query += fmt.Sprintf(" AS %s", q.tableAlias)
		}

		if len(q.joins) > 0 {
			query += fmt.Sprintf(" %s", strings.Join(q.joins, " "))
			args = append(args, q.joinArgs...)
		}

		if len(q.whereClause) > 0 {
			query += fmt.Sprintf(" WHERE %s", strings.Join(q.whereClause, " "))
			args = append(args, q.whereArgs...)
		}
	}

	fmt.Print(args...)

	if q.useContext {
		row = q.db.QueryRowContext(q.ctx, query, args...)
	} else {
		row = q.db.QueryRow(query, args...)
	}

	err = row.Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("query execution failed: %w", err)
	}

	return total, nil
}
