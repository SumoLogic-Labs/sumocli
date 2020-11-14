package api

type CreateUserRequest struct {
	Firstname    string   `json:"firstName"`
	Lastname     string   `json:"lastName"`
	Emailaddress string   `json:"email"`
	Roleids      []string `json:"roleIds"`
}

type CreateUserResponse struct {
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
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
	LastLoginTimestamp string   `json:"lastLoginTimeStamp"`
}
