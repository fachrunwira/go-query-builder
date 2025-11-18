package builder

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"
)

func trx(db *sql.DB, query string, args ...interface{}) error {
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

	if len(args) > 0 {
		_, err = tx.Exec(query, args...)
	} else {
		_, err = tx.Exec(query)
	}

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

func initManipulateData(q *queryStruct) (string, []interface{}) {
	switch q.manipulateType {
	case "insert":
		keys := make([]string, 0, len(q.manipulateArgs[0]))
		for k := range q.manipulateArgs[0] {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		placeholder := "(" + strings.Repeat("?,", len(keys)-1) + "?)"

		valueGroup := make([]string, 0, len(keys))
		values := make([]interface{}, 0, len(keys)*len(q.manipulateArgs))

		for _, row := range q.manipulateArgs {
			valueGroup = append(valueGroup, placeholder)
			for _, k := range keys {
				values = append(values, row[k])
			}
		}

		return fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", q.tableName, strings.Join(keys, ","), strings.Join(valueGroup, ",")), values
	case "update":
		keys := make([]string, 0, len(q.manipulateArgs[0]))
		values := make([]interface{}, 0, len(q.manipulateArgs[0]))
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

	return "", []interface{}{}
}

func printManipulateData(q *queryStruct) *string {
	query, _ := initManipulateData(q)
	return &query
}

func initQuery(q *queryStruct) (string, []interface{}) {
	fields := "*"
	if len(q.fields) > 0 {
		fields = strings.Join(q.fields, ",")
	}

	query := fmt.Sprintf("SELECT %s FROM %s", fields, q.tableName)
	if q.tableAlias != "" {
		query += " AS " + q.tableAlias
	}

	query += ";"
	args := q.whereArgs

	return query, args
}

func printQuery(q *queryStruct) *string {
	query, _ := initQuery(q)
	return &query
}
