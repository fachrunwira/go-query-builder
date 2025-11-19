package builder

import (
	"fmt"
	"strings"
)

func (q *queryStruct) InsertRaw(query string, bindings ...interface{}) error {
	switch {
	case query == "":
		return fmt.Errorf("query is empty")
	case len(query) < 11:
		return fmt.Errorf("query inserted is not valid")
	case strings.ToLower(query[0:11]) != "insert into":
		return fmt.Errorf("query inserted is not valid")
	}

	return trx(q.db, query, bindings...)
}

func (q *queryStruct) UpdateRaw(query string, bindings ...interface{}) error {
	switch {
	case query == "":
		return fmt.Errorf("query is empty")
	case len(query) < 6:
		return fmt.Errorf("query inserted is not valid")
	case strings.ToLower(query[0:6]) != "update":
		return fmt.Errorf("query inserted is not valid")
	}
	return trx(q.db, query, bindings...)
}

func (q *queryStruct) DeleteRaw(query string, bindings ...interface{}) error {
	switch {
	case query == "":
		return fmt.Errorf("query is empty")
	case len(query) < 11:
		return fmt.Errorf("query inserted is not valid")
	case strings.ToLower(query[0:11]) != "insert into":
		return fmt.Errorf("query inserted is not valid")
	}
	return trx(q.db, query, bindings...)
}

func (q *queryStruct) Insert(data any) generateCreateQuery {
	rows, err := convertRows(data)
	if err != nil {
		q.errors = err
		return q
	}

	if len(rows) == 0 {
		q.errors = fmt.Errorf("row data is empty")
		return q
	}

	q.manipulateType = "insert"
	q.manipulateArgs = rows

	return q
}

func (q *queryStruct) Update(data any) generateCreateQuery {
	rows, err := convertRow(data)
	if err != nil {
		q.errors = err
		return q
	}

	if len(rows) == 0 {
		q.errors = fmt.Errorf("data is empty")
		return q
	}

	q.manipulateType = "update"
	q.manipulateArgs = append(q.manipulateArgs, rows)

	return q
}

func (q *queryStruct) Delete() generateCreateQuery {
	q.manipulateType = "delete"
	return q
}
