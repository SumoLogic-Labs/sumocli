package update

import (
	"github.com/SumoLogic-Labs/sumocli/pkg/cmdutils"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip/types"
	"github.com/spf13/cobra"
)

func NewCmdPasswordPolicyUpdate(client *cip.APIClient) *cobra.Command {
	var (
		minLength                      int32
		maxLength                      int32
		mustContainLowercase           bool
		mustContainUppercase           bool
		mustContainDigits              bool
		mustContainSpecialCharacters   bool
		maxPasswordAgeInDays           int32
		minUniquePasswords             int32
		accountLockoutThreshold        int32
		failedLoginResetDurationInMins int32
		accountLockoutDurationInMins   int32
		requireMfa                     bool
		rememberMfa                    bool
	)
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the current password policy.",
		Run: func(cmd *cobra.Command, args []string) {
			updatePasswordPolicy(minLength, maxLength, mustContainLowercase, mustContainUppercase, mustContainDigits,
				mustContainSpecialCharacters, maxPasswordAgeInDays, minUniquePasswords, accountLockoutThreshold,
				failedLoginResetDurationInMins, accountLockoutDurationInMins, requireMfa, rememberMfa, client)
		},
	}
	cmd.Flags().Int32Var(&minLength, "minLength", 8, "Specify the minimum password length")
	cmd.Flags().Int32Var(&maxLength, "maxLength", 128, "Specify the maximum password length")
	cmd.Flags().BoolVar(&mustContainLowercase, "mustContainLowercase", true, "Set to false if you don't require passwords to contain lower case characters")
	cmd.Flags().BoolVar(&mustContainUppercase, "mustContainUppercase", true, "Set to false if you don't require passwords to contain upper case characters")
	cmd.Flags().BoolVar(&mustContainDigits, "mustContainDigits", true, "Set to false if you don't require passwords to contain digits")
	cmd.Flags().BoolVar(&mustContainSpecialCharacters, "mustContainSpecialCharacters", true, "set to false if you don't require passwords to contain special characters")
	cmd.Flags().Int32Var(&maxPasswordAgeInDays, "maxPasswordAgeInDays", 365, "Specify the maximum password age (in days)")
	cmd.Flags().Int32Var(&minUniquePasswords, "minUniquePasswords", 10, "Specify the minimum number of unique new passwords that a user must use before an old password can be reused.")
	cmd.Flags().Int32Var(&accountLockoutThreshold, "accountLockoutThreshold", 6, "Specify the number of failed login attempts allowed before account is locked-out.")
	cmd.Flags().Int32Var(&failedLoginResetDurationInMins, "failedLoginResetDurationInMins", 10, "Specify the duration of time in minutes that must elapse from the first failed login attempt after which failed login count is reset to 0.")
	cmd.Flags().Int32Var(&accountLockoutDurationInMins, "accountLockoutDurationInMins", 30, "Specify the duration of time in minutes that a locked-out account remained locked before getting unlocked automatically.")
	cmd.Flags().BoolVar(&requireMfa, "requireMfa", false, "Set to true if you require users to have MFA enabled")
	cmd.Flags().BoolVar(&rememberMfa, "rememberMfa", true, "Set to false if MFA should not be remembered on the browser.")
	return cmd
}

func updatePasswordPolicy(minLength int32, maxLength int32, mustContainLowercase bool, mustContainUppercase bool,
	mustContainDigits bool, mustContainSpecialCharacters bool, maxPasswordAgeInDays int32, minUniquePasswords int32,
	accountLockoutThreshold int32, failedLoginResetDurationInMins int32, accountLockoutDurationInMins int32,
	requireMfa bool, rememberMfa bool, client *cip.APIClient) {
	data, response, err := client.SetPasswordPolicy(types.PasswordPolicy{
		MinLength:                      minLength,
		MaxLength:                      maxLength,
		MustContainLowercase:           mustContainLowercase,
		MustContainUppercase:           mustContainUppercase,
		MustContainDigits:              mustContainDigits,
		MustContainSpecialChars:        mustContainSpecialCharacters,
		MaxPasswordAgeInDays:           maxPasswordAgeInDays,
		MinUniquePasswords:             minUniquePasswords,
		AccountLockoutThreshold:        accountLockoutThreshold,
		FailedLoginResetDurationInMins: failedLoginResetDurationInMins,
		AccountLockoutDurationInMins:   accountLockoutDurationInMins,
		RequireMfa:                     requireMfa,
		RememberMfa:                    rememberMfa,
	})
	if err != nil {
		cmdutils.OutputError(response, err)
	} else {
		cmdutils.Output(data, response, err, "")
	}
}
