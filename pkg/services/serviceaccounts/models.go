package serviceaccounts

import "github.com/grafana/grafana/pkg/services/accesscontrol"

var (
	ScopeAll = "serviceaccounts:*"
	ScopeID  = accesscontrol.Scope("serviceaccounts", "id", accesscontrol.Parameter(":serviceaccountId"))
)

const (
	ActionDelete = "serviceaccounts:delete"
	ActionRead   = "serviceaccounts:read"
)

type Status struct {
	Enabled bool `json:"enabled"`
}