package builder

import "fmt"

func (q *queryStruct) Where(column string, values interface{}) manipulateData {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf(" AND %s = ?", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("%s = ?", column))
	}

	q.whereArgs = append(q.whereArgs, values)
	return q
}

func (q *queryStruct) WhereRaw(query string, bindings ...interface{}) manipulateData {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf(" AND %s", query))
	} else {
		q.whereClause = append(q.whereClause, query)
	}

	if len(bindings) > 0 {
		q.whereArgs = append(q.whereArgs, bindings...)
	}
	return q
}

func (q *queryStruct) WhereNot(column string, values interface{}) manipulateData {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf(" AND NOT %s = ?", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("NOT %s = ?", column))
	}

	q.whereArgs = append(q.whereArgs, values)
	return q
}

func (q *queryStruct) OrWhereRaw(query string, bindings ...interface{}) manipulateData {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf(" OR %s", query))
	} else {
		q.whereClause = append(q.whereClause, query)
	}

	if len(bindings) > 0 {
		q.whereArgs = append(q.whereArgs, bindings...)
	}
	return q
}

func (q *queryStruct) OrWhere(column string, values interface{}) manipulateData {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf(" OR %s = ?", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("%s = ?", column))
	}

	q.whereArgs = append(q.whereArgs, values)
	return q
}

func (q *queryStruct) OrWhereNot(column string, values interface{}) manipulateData {
	if len(q.whereClause) > 0 {
		q.whereClause = append(q.whereClause, fmt.Sprintf(" OR NOT %s = ?", column))
	} else {
		q.whereClause = append(q.whereClause, fmt.Sprintf("NOT %s = ?", column))
	}

	q.whereArgs = append(q.whereArgs, values)
	return q
}
