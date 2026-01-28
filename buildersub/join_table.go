package buildersub

import (
	"fmt"

	"github.com/fachrunwira/go-query-builder/joinbuilder"
)

func (qs *queryStruct) Join(callback func() joinbuilder.JoinQuery) SubQuery {
	cb := callback()
	if cb == nil {
		qs.err = fmt.Errorf("no join builder detected in sub query")
		return qs
	}

	query, args, err := joinbuilder.Build(cb)
	if err != nil {
		qs.err = fmt.Errorf("error while trying to make join in sub query: %w", err)
		return qs
	}

	qs.joins = append(qs.joins, fmt.Sprintf("INNER JOIN %s", query))
	qs.joinArgs = append(qs.joinArgs, args...)
	return qs
}

func (qs *queryStruct) LeftJoin(callback func() joinbuilder.JoinQuery) SubQuery {
	cb := callback()
	if cb == nil {
		qs.err = fmt.Errorf("no join builder detected in sub query")
		return qs
	}

	query, args, err := joinbuilder.Build(cb)
	if err != nil {
		qs.err = fmt.Errorf("error while trying to make join in sub query: %w", err)
		return qs
	}

	qs.joins = append(qs.joins, fmt.Sprintf("LEFT JOIN %s", query))
	qs.joinArgs = append(qs.joinArgs, args...)
	return qs
}

func (qs *queryStruct) RightJoin(callback func() joinbuilder.JoinQuery) SubQuery {
	cb := callback()
	if cb == nil {
		qs.err = fmt.Errorf("no join builder detected in sub query")
		return qs
	}

	query, args, err := joinbuilder.Build(cb)
	if err != nil {
		qs.err = fmt.Errorf("error while trying to make join in sub query: %w", err)
		return qs
	}

	qs.joins = append(qs.joins, fmt.Sprintf("RIGHT JOIN %s", query))
	qs.joinArgs = append(qs.joinArgs, args...)
	return qs
}

func (qs *queryStruct) CrossJoin(tableJoin string) SubQuery {
	if tableJoin == "" {
		qs.err = fmt.Errorf("joined table is empty")
	}
	qs.joins = append(qs.joins, fmt.Sprintf("CROSS JOIN %s", tableJoin))
	return qs
}
