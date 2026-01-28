package buildersub

import (
	"fmt"
	"strings"

	"github.com/fachrunwira/go-query-builder/clauseoperators"
)

func clearStrings(data []string) []string {
	n := 0
	for _, v := range data {
		v = strings.TrimSpace(v)
		if v != "" {
			data[n] = v
			n++
		}
	}

	data = data[:n]
	return data
}

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
