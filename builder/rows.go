package builder

import (
	"database/sql"
	"fmt"
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
