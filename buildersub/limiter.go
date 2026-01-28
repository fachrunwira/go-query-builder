package buildersub

func (qs *queryStruct) Limit(limit uint) SubQuery {
	qs.limits = limit
	return qs
}
