package builder

import "database/sql"

// Make initialize the querybuilder.
func Make(_db *sql.DB) initialStage {
	return &queryStruct{
		db: _db,
	}
}
