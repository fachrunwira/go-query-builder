package clauseoperators

type whereOperator string
type Operators whereOperator

const (
	EQUAL              Operators = "="
	LIKE               Operators = "LIKE"
	LESS_THAN          Operators = "<"
	LESS_THAN_EQUAL    Operators = "<="
	GREATER_THAN       Operators = ">"
	GREATER_THAN_EQUAL Operators = ">="
	IN                 Operators = "IN"
)
