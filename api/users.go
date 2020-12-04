package api

type CreateUserRequest struct {
	Firstname    string   `json:"firstName"`
	Lastname     string   `json:"lastName"`
	Emailaddress string   `json:"email"`
	Roleids      []string `json:"roleIds"`
}

type DisableUserMfa struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserEmail struct {
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Firstname string   `json:"firstName"`
	Lastname  string   `json:"lastName"`
	IsActive  bool     `json:"isActive"`
	Roleids   []string `json:"roleIds"`
}

type UserResponse struct {
	Firstname          string   `json:"firstName"`
	Lastname           string   `json:"lastName"`
	Email              string   `json:"email"`
	RoleIds            []string `json:"roleIds"`
	CreatedAt          string   `json:"createdAt"`
	CreatedBy          string   `json:"createdBy"`
	ModifiedAt         string   `json:"modifiedAt"`
	ModifiedBy         string   `json:"modifiedBy"`
	Id                 string   `json:"id"`
	IsActive           bool     `json:"isActive"`
	IsLocked           bool     `json:"isLocked"`
	IsMfaEnabled       bool     `json:"isMfaEnabled"`
	LastLoginTimestamp string   `json:"lastLoginTimestamp"`
}

type Users struct {
	Data []UserResponse `json:"data"`
	Next string         `json:"next"`
}
