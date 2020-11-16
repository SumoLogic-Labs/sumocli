package api

type CreateRoleRequest struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description,omitempty"`
	FilterPredicate      string   `json:"filterPredicate,omitempty"`
	Users                []string `json:"users,omitempty"`
	Capabilities         []string `json:"capabilities,omitempty"`
	AutoFillDependencies bool     `json:"autoFillDependencies,omitempty"`
}

type Role struct {
	Data []RoleData `json:"data"`
}

type RoleData struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	FilterPredicate      string   `json:"filterPredicate"`
	Users                []string `json:"users"`
	Capabilities         []string `json:"capabilities"`
	AutofillDependencies bool     `json:"autofillDependencies"`
	CreatedAt            string   `json:"createdAt"`
	CreatedBy            string   `json:"createdBy"`
	ModifiedAt           string   `json:"modifiedAt"`
	ModifiedBy           string   `json:"modifiedBy"`
	Id                   string   `json:"id"`
	SystemDefined        bool     `json:"systemDefined"`
}

type UpdateRoleRequest struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	FilterPredicate      string   `json:"filterPredicate"`
	Users                []string `json:"users,omitempty"`
	Capabilities         []string `json:"capabilities,omitempty"`
	AutoFillDependencies bool     `json:"autoFillDependencies,omitempty"`
}
