package update

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wizedkyle/sumocli/api"
	"github.com/wizedkyle/sumocli/pkg/cmd/factory"
	"github.com/wizedkyle/sumocli/pkg/logging"
	"io"
)

func NewCmdPasswordPolicyUpdate() *cobra.Command {
	var (
		minLength                      int
		maxLength                      int
		mustContainLowercase           bool
		mustContainUppercase           bool
		mustContainDigits              bool
		mustContainSpecialCharacters   bool
		maxPasswordAgeInDays           int
		minUniquePasswords             int
		accountLockoutThreshold        int
		failedLoginResetDurationInMins int
		accountLockoutDurationInMins   int
		requireMfa                     bool
		rememberMfa                    bool
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update the current password policy.",
		Run: func(cmd *cobra.Command, args []string) {
			updatePasswordPolicy(minLength, maxLength, mustContainLowercase, mustContainUppercase, mustContainDigits,
				mustContainSpecialCharacters, maxPasswordAgeInDays, minUniquePasswords, accountLockoutThreshold,
				failedLoginResetDurationInMins, accountLockoutDurationInMins, requireMfa, rememberMfa)
		},
	}
	cmd.Flags().IntVar(&minLength, "minLength", 8, "Specify the minimum password length")
	cmd.Flags().IntVar(&maxLength, "maxLength", 128, "Specify the maximum password length")
	cmd.Flags().BoolVar(&mustContainLowercase, "mustContainLowercase", true, "Set to false if you don't require passwords to contain lower case characters")
	cmd.Flags().BoolVar(&mustContainUppercase, "mustContainUppercase", true, "Set to false if you don't require passwords to contain upper case characters")
	cmd.Flags().BoolVar(&mustContainDigits, "mustContainDigits", true, "Set to false if you don't require passwords to contain digits")
	cmd.Flags().BoolVar(&mustContainSpecialCharacters, "mustContainSpecialCharacters", true, "set to false if you don't require passwords to contain special characters")
	cmd.Flags().IntVar(&maxPasswordAgeInDays, "maxPasswordAgeInDays", 365, "Specify the maximum password age (in days)")
	cmd.Flags().IntVar(&minUniquePasswords, "minUniquePasswords", 10, "Specify the minimum number of unique new passwords that a user must use before an old password can be reused.")
	cmd.Flags().IntVar(&accountLockoutThreshold, "accountLockoutThreshold", 6, "Specify the number of failed login attempts allowed before account is locked-out.")
	cmd.Flags().IntVar(&failedLoginResetDurationInMins, "failedLoginResetDurationInMins", 10, "Specify the duration of time in minutes that must elapse from the first failed login attempt after which failed login count is reset to 0.")
	cmd.Flags().IntVar(&accountLockoutDurationInMins, "accountLockoutDurationInMins", 30, "Specify the duration of time in minutes that a locked-out account remained locked before getting unlocked automatically.")
	cmd.Flags().BoolVar(&requireMfa, "requireMfa", false, "Set to true if you require users to have MFA enabled")
	cmd.Flags().BoolVar(&rememberMfa, "rememberMfa", true, "Set to false if MFA should not be remembered on the browser.")
	return cmd
}

func updatePasswordPolicy(minLength int, maxLength int, mustContainLowercase bool, mustContainUppercase bool,
	mustContainDigits bool, mustContainSpecialCharacters bool, maxPasswordAgeInDays int, minUniquePasswords int,
	accountLockoutThreshold int, failedLoginResetDurationInMins int, accountLockoutDurationInMins int,
	requireMfa bool, rememberMfa bool) {
	var passwordPolicyResponse api.GetPasswordPolicy
	log := logging.GetConsoleLogger()
	requestBodySchema := api.GetPasswordPolicy{
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
	}
	requestBody, err := json.Marshal(requestBodySchema)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request body")
	}
	requestUrl := "v1/passwordPolicy"
	client, request := factory.NewHttpRequestWithBody("PUT", requestUrl, requestBody)
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to make http request to " + requestUrl)
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &passwordPolicyResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response body")
	}

	passwordPolicyResponseJson, err := json.MarshalIndent(passwordPolicyResponse, "", "    ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
	}

	if response.StatusCode != 200 {
		factory.HttpError(response.StatusCode, responseBody, log)
	} else {
		fmt.Println(string(passwordPolicyResponseJson))
	}
}
