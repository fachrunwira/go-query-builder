package builder

import (
	"context"
	"database/sql"
)

// Make initialize the querybuilder.
func Make(_db *sql.DB) initialStage {
	return &queryStruct{
		db:         _db,
		useContext: false,
	}
}

func MakeWithContext(ctx context.Context, _db *sql.DB) initialStage {
	return &queryStruct{
		db:         _db,
		useContext: true,
		ctx:        ctx,
	}
}
