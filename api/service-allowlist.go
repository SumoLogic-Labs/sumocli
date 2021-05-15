package api

type GetAllowlistStatus struct {
	ContentEnabled bool `json:"contentEnabled"`
	LoginEnabled   bool `json:"loginEnabled"`
}

type ListServiceAllowlist struct {
	Data []AllowlistCIDR `json:"data"`
}

type AllowlistCIDR struct {
	Cidr        string `json:"cidr"`
	Description string `json:"description"`
}
