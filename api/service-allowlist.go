package api

type ListServiceAllowlist struct {
	Data []allowlistCIDR `json:"data"`
}

type allowlistCIDR struct {
	Cidr        string `json:"cidr"`
	Description string `json:"description"`
}
