package buildersub

import "fmt"

func (qs *queryStruct) Table(table string, alias ...string) SubQuery {
	if len(alias) > 0 {
		qs.tableAlias = alias[0]
	}

	if table == "" {
		qs.err = fmt.Errorf("table name cannot be empty")
	} else {
		qs.tableName = table
	}

	return qs
}
