package api

type GetPasswordPolicy struct {
	MinLength                      int  `json:"minLength"`
	MaxLength                      int  `json:"maxLength"`
	MustContainLowercase           bool `json:"mustContainLowercase"`
	MustContainUppercase           bool `json:"mustContainUppercase"`
	MustContainDigits              bool `json:"mustContainDigits"`
	MustContainSpecialChars        bool `json:"mustContainSpecialChars"`
	MaxPasswordAgeInDays           int  `json:"maxPasswordAgeInDays"`
	MinUniquePasswords             int  `json:"minUniquePasswords"`
	AccountLockoutThreshold        int  `json:"accountLockoutThreshold"`
	FailedLoginResetDurationInMins int  `json:"failedLoginResetDurationInMins"`
	AccountLockoutDurationInMins   int  `json:"accountLockoutDurationInMins"`
	RequireMfa                     bool `json:"requireMfa"`
	RememberMfa                    bool `json:"rememberMfa"`
}
