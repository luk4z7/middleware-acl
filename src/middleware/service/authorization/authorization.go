package authorization

import (
	"middleware/core/adapter"
	"github.com/casbin/casbin"
)

var adp = adapter.New()

// Because the DB is empty at first,
// so we need to load the policy from the file adapter (.CSV) first
func initPolicy() (*casbin.Enforcer, error) {
	e := casbin.NewEnforcer("/go/src/middleware/setting/authorization/rbac_model.conf", "/go/src/middleware/setting/authorization/rbac_policy.csv")
	err := adp.SavePolicy(e.GetModel())
	if err != nil {
		return e, err
	}
	return e, nil
}

func GetEnforcer() *casbin.Enforcer {
	e := casbin.NewEnforcer("/go/src/middleware/setting/authorization/rbac_model.conf", adp)
	if len(e.GetPolicy()) == 0 {
		init, err := initPolicy()
		if err != nil {
			panic(err)
		}
		e = init
	}
	return e
}