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

func (q *queryStruct) ToRaw() (string, error) {
	fields := "*"
	if len(q.fields) > 0 {
		fields = strings.Join(q.fields, ", ")
	}

	query := fmt.Sprintf("SELECT %s FROM %s", fields, q.tableName)
	if q.tableAlias != "" {
		query += " AS " + q.tableAlias
	}

	if len(q.whereClause) > 0 {
		var errWhere error
		for k, v := range q.whereClause {
			clauseValue := v
			for range q.whereArgs {
				if strings.Contains(clauseValue, "?") {
					var popValue any
					popValue, q.whereArgs = q.whereArgs[0], q.whereArgs[1:]

					clauseValue, errWhere = getArgsValue(clauseValue, popValue)
					if errWhere != nil {
						return "", fmt.Errorf("failed to generate raw query sql: %w", errWhere)
					}
				} else {
					continue
				}
			}
			q.whereClause[k] = clauseValue
		}

		query += fmt.Sprintf(" WHERE %s", strings.Join(q.whereClause, " "))
	}

	return query + ";", nil
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

		if q.useContext {
			return trxContext(q.ctx, q.db, query, args...)
		}

		return trx(q.db, query, args...)
	case "update":
		query, args := initManipulateData(q)

		if q.useContext {
			return trxContext(q.ctx, q.db, query, args...)
		}

		return trx(q.db, query, args...)
	case "delete":
		query, args := initManipulateData(q)

		if q.useContext {
			return trxContext(q.ctx, q.db, query, args...)
		}

		return trx(q.db, query, args...)
	}

	return nil
}

func (q *queryStruct) Error() error {
	return q.errors
}
