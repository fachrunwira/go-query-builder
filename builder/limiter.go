package builder

func (q *queryStruct) Limit(limit uint) offsetInterface {
	q.limit = limit
	return q
}

func (q *queryStruct) Offset(offset uint) generateSelectQuery {
	q.offset = offset
	return q
}
