package joinbuilder

import (
	"fmt"
	"strings"

	"github.com/fachrunwira/go-query-builder/clauseoperators"
)

func getClauseOperator(operator clauseoperators.Operators, args ...any) (string, string, error) {
	if len(args) == 0 {
		return "", "", fmt.Errorf("arguments value is at least one")
	}

	switch operator {
	case clauseoperators.EQUAL,
		clauseoperators.GREATER_THAN,
		clauseoperators.GREATER_THAN_EQUAL,
		clauseoperators.LESS_THAN,
		clauseoperators.LESS_THAN_EQUAL,
		clauseoperators.LIKE:
		return string(operator), "", nil
	case clauseoperators.IN:
		placeholder := "(" + strings.Repeat("?,", len(args)-1) + "?)"

		return string(operator), placeholder, nil
	}

	return "", "", fmt.Errorf("invalid operators %q", operator)
}

func getClauseOperatorSub(operator clauseoperators.Operators) (string, error) {
	switch operator {
	case clauseoperators.EQUAL,
		clauseoperators.GREATER_THAN,
		clauseoperators.GREATER_THAN_EQUAL,
		clauseoperators.LESS_THAN,
		clauseoperators.LESS_THAN_EQUAL,
		clauseoperators.LIKE,
		clauseoperators.IN:
		return string(operator), nil
	}

	return "", fmt.Errorf("invalid operators %q", operator)
}
