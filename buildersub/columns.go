package buildersub

func (qs *queryStruct) Select(columns ...string) SubQuery {
	if len(columns) > 0 && len(clearStrings(columns)) > 0 {
		qs.fields = columns
	}

	return qs
}
