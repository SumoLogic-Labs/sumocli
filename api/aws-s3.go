package api

type AWSPolicy struct {
	Version   string               `json:"Version"`
	Statement []AWSPolicyStatement `json:"Statement"`
}

type AWSPolicyStatement struct {
	Sid       string             `json:"Sid"`
	Effect    string             `json:"Effect"`
	Principal AWSPolicyPrincipal `json:"Principal"`
	Action    []string           `json:"Action"`
	Resource  []string           `json:"Resource"`
	Condition AWSPolicyCondition `json:"Condition,omitempty"`
}

type AWSPolicyPrincipal struct {
	AWS           []string `json:"AWS,omitempty"`
	CanonicalUser string   `json:"CanonicalUser,omitempty"`
	Federated     string   `json:"Federated,omitempty"`
	Service       string   `json:"Service,omitempty"`
}

type AWSPolicyCondition struct {
	StringEquals S3XAmzAcl                `json:"StringEquals,omitempty"`
	StringLike   KMSEncryptionContext     `json:"StringLike,omitempty"`
	Null         KMSEncryptionContextBool `json:"Null,omitempty"`
}

type S3XAmzAcl struct {
	S3XAmzAcl string `json:"s3:x-amz-acl,omitempty"`
}

type KMSEncryptionContext struct {
	KMSEncryption []string `json:"kms:EncryptionContext:aws:cloudtrail:arn,omitempty"`
}

type KMSEncryptionContextBool struct {
	KMSEncryption bool `json:"kms:EncryptionContext:aws:cloudtrail:arn,omitempty"`
}
