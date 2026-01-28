package joinbuilder

import (
	"fmt"
	"strings"
)

func Make() JoinQuery {
	return &queryStruct{}
}

func Build(joinQuery JoinQuery) (string, []any, error) {
	jq, ok := joinQuery.(*queryStruct)
	if !ok {
		return "", nil, fmt.Errorf("invalid type for join query")
	}

	query := strings.Join(jq.whereClause, " ")
	var args []any

	args = append(args, jq.whereArgs...)

	return query, args, nil
}
