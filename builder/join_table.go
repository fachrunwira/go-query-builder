package builder

import "fmt"

func (q *queryStruct) Join(tableJoin, table_1, operator, table_2 string) manipulateData {
	switch {
	case tableJoin == "":
		q.errors = fmt.Errorf("join table is empty")
	case table_1 == "":
		q.errors = fmt.Errorf("join column is empty")
	case operator == "":
		q.errors = fmt.Errorf("join operator is empty")
	case table_2 == "":
		q.errors = fmt.Errorf("join column is empty")
	default:
		q.joins = append(q.joins, fmt.Sprintf("JOIN %s ON %s %s %s", tableJoin, table_1, operator, table_2))
	}
	return q
}

func (q *queryStruct) LeftJoin(tableJoin, table_1, operator, table_2 string) manipulateData {
	switch {
	case tableJoin == "":
		q.errors = fmt.Errorf("join table is empty")
	case table_1 == "":
		q.errors = fmt.Errorf("join column is empty")
	case operator == "":
		q.errors = fmt.Errorf("join operator is empty")
	case table_2 == "":
		q.errors = fmt.Errorf("join column is empty")
	default:
		q.joins = append(q.joins, fmt.Sprintf("LEFT JOIN %s ON %s %s %s", tableJoin, table_1, operator, table_2))
	}
	return q
}

func (q *queryStruct) RightJoin(tableJoin, table_1, operator, table_2 string) manipulateData {
	switch {
	case tableJoin == "":
		q.errors = fmt.Errorf("join table is empty")
	case table_1 == "":
		q.errors = fmt.Errorf("join column is empty")
	case operator == "":
		q.errors = fmt.Errorf("join operator is empty")
	case table_2 == "":
		q.errors = fmt.Errorf("join column is empty")
	default:
		q.joins = append(q.joins, fmt.Sprintf("RIGHT JOIN %s ON %s %s %s", tableJoin, table_1, operator, table_2))
	}
	return q
}

func (q *queryStruct) JoinWhere(tableJoin, table_1, operator string, bindings ...interface{}) manipulateData {
	switch {
	case tableJoin == "":
		q.errors = fmt.Errorf("join table is empty")
	case table_1 == "":
		q.errors = fmt.Errorf("join column is empty")
	case operator == "":
		q.errors = fmt.Errorf("join operator is empty")
	}

	if len(bindings) > 0 {
		q.joinArgs = append(q.joinArgs, bindings...)
	}

	return q
}
