package buildersub

func (q *queryStruct) GroupBy(columns ...string) SubQuery {
	if len(columns) > 0 && len(clearStrings(columns)) > 0 {
		q.grouping = append(q.grouping, clearStrings(columns)...)
	}

	return q
}
