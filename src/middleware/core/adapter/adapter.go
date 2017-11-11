package adapter

import (
	"github.com/casbin/mongodb-adapter"
	"github.com/casbin/casbin/persist"
)

func New() persist.Adapter {
	a := mongodbadapter.NewAdapter("admin:admin@mongo:27017/middleware-acl")
	return a
}