package joinbuilder

import (
	"fmt"

	"github.com/fachrunwira/go-query-builder/clauseoperators"
)

func (jq *queryStruct) On(clause_1 string, operator clauseoperators.Operators, clause_2 string) JoinQuery {
	op, err := getClauseOperatorSub(operator)
	if err != nil {
		jq.errors = fmt.Errorf("error in join query: %w", err)
		return jq
	}

	if len(jq.whereClause) > 0 {
		jq.whereClause = append(jq.whereClause, fmt.Sprintf("AND %s %s %s", clause_1, op, clause_2))
	} else {
		jq.whereClause = append(jq.whereClause, fmt.Sprintf("ON %s %s %s", clause_1, op, clause_2))
	}

	return jq
}

func (jq *queryStruct) Or(clause_1 string, operator clauseoperators.Operators, clause_2 string) JoinQuery {
	op, err := getClauseOperatorSub(operator)
	if err != nil {
		jq.errors = fmt.Errorf("error in join query: %w", err)
		return jq
	}

	if len(jq.whereClause) > 0 {
		jq.whereClause = append(jq.whereClause, fmt.Sprintf("OR %s %s %s", clause_1, op, clause_2))
	} else {
		jq.whereClause = append(jq.whereClause, fmt.Sprintf("ON %s %s %s", clause_1, op, clause_2))
	}

	return jq
}
