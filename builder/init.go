package builder

import (
	"context"
	"database/sql"
	"fmt"
)

// Make initialize the querybuilder.
func Make(_db *sql.DB) initialStage {
	return &queryStruct{
		db:         _db,
		useContext: false,
	}
}

func initContextBuilder() *queryStruct {
	return &queryStruct{}
}

func MakeWithContext(ctx context.Context, _db *sql.DB) initialStage {
	qStruct := initContextBuilder()
	qStruct.useContext = true
	qStruct.ctx = ctx

	valueCtx := ctx.Value(getContextKey())
	if valueCtx == nil {
		qStruct.errors = fmt.Errorf("key not found in context: %v", getContextKey())
		return qStruct
	}

	db, ok := valueCtx.(*sql.DB)
	if !ok {
		qStruct.errors = fmt.Errorf("expected *sql.DB, but got %T for key %s", valueCtx, getContextKey())
		return qStruct
	}
	qStruct.db = db
	return qStruct
}

func UseDefaultKey() {
	contextKeyLock.Lock()
	defer contextKeyLock.Unlock()
	contextKey = defaultKeyContext
}

func SetDefaultKey(key any) {
	contextKeyLock.Lock()
	defer contextKeyLock.Unlock()
	contextKey = defaultKeyContext
}

func getContextKey() any {
	contextKeyLock.RLock()
	defer contextKeyLock.RUnlock()
	return contextKey
}
