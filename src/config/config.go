package config

import "github.com/couchbase/gocb/v2"

var SharedScope *gocb.Scope

// InitializeSharedScope initializes the shared scope.
func InitializeSharedScope(scope *gocb.Scope) {
	SharedScope = scope
}
