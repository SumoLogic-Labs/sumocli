package api

type GetAllowlistStatus struct {
	ContentEnabled bool `json:"contentEnabled"`
	LoginEnabled   bool `json:"loginEnabled"`
}

type ListServiceAllowlist struct {
	Data []allowlistCIDR `json:"data"`
}

type allowlistCIDR struct {
	Cidr        string `json:"cidr"`
	Description string `json:"description"`
}
