package builder

func (q *queryStruct) GroupBy(columns ...string) manipulateData {
	if len(columns) > 0 && len(clearStrings(columns)) > 0 {
		q.grouping = append(q.grouping, clearStrings(columns)...)
	}

	return q
}
