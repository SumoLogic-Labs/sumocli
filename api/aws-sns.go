package api

type AWSSNSPolicy struct {
	Version   string                  `json:"Version"`
	Statement []AWSSNSPolicyStatement `json:"Statement"`
}

type AWSSNSPolicyStatement struct {
	Sid       string                `json:"Sid"`
	Effect    string                `json:"Effect"`
	Principal AWSSNSPolicyPrincipal `json:"Principal"`
	Action    []string              `json:"Action"`
	Resource  []string              `json:"Resource"`
	Condition AWSSNSPolicyCondition `json:"Condition,omitempty"`
}

type AWSSNSPolicyPrincipal struct {
	AWS           []string `json:"AWS,omitempty"`
	CanonicalUser string   `json:"CanonicalUser,omitempty"`
	Federated     string   `json:"Federated,omitempty"`
	Service       string   `json:"Service,omitempty"`
}

type AWSSNSPolicyCondition struct {
	ArnLike      AWSSNSSourceArn     `json:"ArnLike,omitempty"`
	StringEquals AWSSNSSourceAccount `json:"StringEquals,omitempty"`
}

type AWSSNSSourceAccount struct {
	AWSSourceAccount string `json:"aws:SourceAccount"`
}

type AWSSNSSourceArn struct {
	AWSSourceArn string `json:"aws:SourceArn"`
}
