package builder

import "sync"

// Row represents a single database record, where each key is the column name
// and the value is the associated field value.
type Row map[string]any

// Rows represents a collection of Row values, typically used to return
// multiple records from a query or result set.
type Rows []Row

type contextKeyType struct{}

var (
	defaultKeyContext     = contextKeyType{}
	contextKey        any = contextKeyType{}
	contextKeyLock    sync.RWMutex
)
