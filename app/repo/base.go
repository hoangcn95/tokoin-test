package repo

import (
	"app/common/adapter"
)

// Mongo struct
type Mongo struct {
	Session    *adapter.Mongo
	Collection string
}

