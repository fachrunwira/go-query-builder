package builder

import (
	"fmt"
	"strings"
)

// Table Set the table name and optional alias.
func (q *queryStruct) Table(table string, alias ...string) insertOrQueryingStage {
	if len(alias) > 0 && len(clearStrings(alias)) > 0 {
		q.tableAlias = clearStrings(alias)[0]
	}

	if table == "" {
		q.errors = fmt.Errorf("table name is not set")
	} else {
		q.tableName = table
	}

	return q
}

func (q *queryStruct) Select(columns ...string) manipulateData {
	if len(columns) > 0 && len(clearStrings(columns)) > 0 {
		q.fields = append(q.fields, columns...)
	}
	return q
}

func (q *queryStruct) ToRaw() string {
	fields := "*"
	if len(q.fields) > 0 {
		fields = strings.Join(q.fields, ", ")
	}

	query := fmt.Sprintf("SELECT %s FROM %s", fields, q.tableName)
	if q.tableAlias != "" {
		query += " AS " + q.tableAlias
	}

	return query + ";"
}

func (q *queryStruct) ToSql() (string, error) {
	if q.errors != nil {
		return "", q.errors
	}

	if q.manipulateType != "" {
		return printManipulateData(q), nil
	}

	return printQuery(q), nil
}

func (q *queryStruct) Save() error {
	switch q.manipulateType {
	case "insert":
		query, args := initManipulateData(q)

		return trx(q.db, query, args...)
	case "update":
		query, args := initManipulateData(q)
		return trx(q.db, query, args...)
	case "delete":
		query, args := initManipulateData(q)
		return trx(q.db, query, args...)
	}

	return nil
}

func (q *queryStruct) Error() error {
	return q.errors
}
