package builder

import (
	"fmt"

	"github.com/fachrunwira/go-query-builder/buildersub"
	"github.com/fachrunwira/go-query-builder/joinbuilder"
)

func (q *queryStruct) CrossJoin(tableJoin string) manipulateData {
	if tableJoin == "" {
		q.errors = fmt.Errorf("joined table is empty")
	}
	q.joins = append(q.joins, fmt.Sprintf("CROSS JOIN %s", tableJoin))
	return q
}

func (q *queryStruct) Join(tableJoin string, callback func() joinbuilder.JoinQuery) manipulateData {
	cb := callback()
	if cb == nil {
		q.errors = fmt.Errorf("no join builder detected")
		return q
	}

	query, args, err := joinbuilder.Build(cb)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make join sub: %w", err)
		return q
	}

	q.joins = append(q.joins, fmt.Sprintf("INNER JOIN %s %s", tableJoin, query))
	q.joinArgs = append(q.joinArgs, args...)

	return q
}

func (q *queryStruct) LeftJoin(tableJoin string, callback func() joinbuilder.JoinQuery) manipulateData {
	cb := callback()
	if cb == nil {
		q.errors = fmt.Errorf("no join builder detected")
		return q
	}

	query, args, err := joinbuilder.Build(cb)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make join sub: %w", err)
		return q
	}

	q.joins = append(q.joins, fmt.Sprintf("LEFT JOIN %s %s", tableJoin, query))
	q.joinArgs = append(q.joinArgs, args...)

	return q
}

func (q *queryStruct) RightJoin(tableJoin string, callback func() joinbuilder.JoinQuery) manipulateData {
	cb := callback()
	if cb == nil {
		q.errors = fmt.Errorf("no join builder detected")
		return q
	}

	query, args, err := joinbuilder.Build(cb)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make join sub: %w", err)
		return q
	}

	q.joins = append(q.joins, fmt.Sprintf("RIGHT JOIN %s %s", tableJoin, query))
	q.joinArgs = append(q.joinArgs, args...)

	return q
}

func (q *queryStruct) JoinSub(query func() buildersub.SubQuery, as string, clause func() joinbuilder.JoinQuery) manipulateData {
	cbQuery := query()
	if cbQuery == nil {
		q.errors = fmt.Errorf("no sub query detected")
		return q
	}

	sQuery, sArgs, err := buildersub.Build(cbQuery)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make sub query: %w", err)
		return q
	}

	cbJoin := clause()
	if cbJoin == nil {
		q.errors = fmt.Errorf("no join builder detected")
		return q
	}

	jQuery, jArgs, err := joinbuilder.Build(cbJoin)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make join builder: %w", err)
		return q
	}

	q.joinArgs = append(q.joinArgs, sArgs...)
	q.joinArgs = append(q.joinArgs, jArgs...)

	q.joins = append(q.joins, fmt.Sprintf("INNER JOIN (%s) AS %s %s", sQuery, as, jQuery))
	return q
}

func (q *queryStruct) LeftJoinSub(query func() buildersub.SubQuery, as string, clause func() joinbuilder.JoinQuery) manipulateData {
	cbQuery := query()
	if cbQuery == nil {
		q.errors = fmt.Errorf("no sub query detected")
		return q
	}

	sQuery, sArgs, err := buildersub.Build(cbQuery)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make sub query: %w", err)
		return q
	}

	cbJoin := clause()
	if cbJoin == nil {
		q.errors = fmt.Errorf("no join builder detected")
		return q
	}

	jQuery, jArgs, err := joinbuilder.Build(cbJoin)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make join builder: %w", err)
		return q
	}

	q.joinArgs = append(q.joinArgs, sArgs...)
	q.joinArgs = append(q.joinArgs, jArgs...)

	q.joins = append(q.joins, fmt.Sprintf("LEFT JOIN (%s) AS %s %s", sQuery, as, jQuery))
	return q
}

func (q *queryStruct) RightJoinSub(query func() buildersub.SubQuery, as string, clause func() joinbuilder.JoinQuery) manipulateData {
	cbQuery := query()
	if cbQuery == nil {
		q.errors = fmt.Errorf("no sub query detected")
		return q
	}

	sQuery, sArgs, err := buildersub.Build(cbQuery)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make sub query: %w", err)
		return q
	}

	cbJoin := clause()
	if cbJoin == nil {
		q.errors = fmt.Errorf("no join builder detected")
		return q
	}

	jQuery, jArgs, err := joinbuilder.Build(cbJoin)
	if err != nil {
		q.errors = fmt.Errorf("error while trying to make join builder: %w", err)
		return q
	}

	q.joinArgs = append(q.joinArgs, sArgs...)
	q.joinArgs = append(q.joinArgs, jArgs...)

	q.joins = append(q.joins, fmt.Sprintf("RIGHT JOIN (%s) AS %s %s", sQuery, as, jQuery))
	return q
}
