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
		query += fmt.Sprintf(" LIMIT %d", q.limit)
	}

	if q.offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", q.offset)
	}

	rows, err := q.db.Query(query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no data found")
		}
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := Rows{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuesPtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuesPtrs[i] = &values[i]
		}

		if err = rows.Scan(valuesPtrs...); err != nil {
			return nil, err
		}
		row := Row{}
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

	return result, nil
}

func (q *queryStruct) First() (Row, error) {
	if q.errors != nil {
		return nil, q.errors
	}

	query, args := initQuery(q)
	query += " LIMIT 1;"

	rows, err := q.db.Query(query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no data found")
		}
		return nil, err
	}
	defer rows.Close()

	column, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := Row{}
	for rows.Next() {
		values := make([]interface{}, len(column))
		valuesPtrs := make([]interface{}, len(column))

		for i := range column {
			valuesPtrs[i] = &values[i]
		}

		if err = rows.Scan(valuesPtrs...); err != nil {
			return nil, err
		}

		for i, col := range column {
			val := values[i]
			if b, ok := val.([]byte); ok {
				result[col] = string(b)
			} else {
				result[col] = val
			}
		}
	}

	return result, nil
}
