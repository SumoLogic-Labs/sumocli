package authentication

import (
	"encoding/base64"
	"fmt"
	"github.com/SumoLogic-Labs/sumocli/internal/encryption"
	"github.com/SumoLogic-Labs/sumologic-go-sdk/service/cip"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func ConfirmCredentialsSet(client *cip.APIClient) {
	if client.Cfg.Authentication.AccessId == "" || client.Cfg.Authentication.AccessKey == "" || client.Cfg.BasePath == "" {
		fmt.Println("No authentication credentials set, please run sumocli configure or set environment variables.")
		os.Exit(1)
	}
}

func ConfigName() string {
	var envConfigName = os.Getenv("SUMO_CONFIG_NAME")
	if envConfigName != "" {
		return envConfigName
	}
	return "creds"
}

func ConfigDirPath() string {
	homeDirectory, _ := os.UserHomeDir()
	return filepath.Join(homeDirectory, ".sumocli/credentials")
}

func ConfigFilePath() string {
	return filepath.Join(ConfigDirPath(), ConfigName() + ".json")
}

func ReadAccessId() string {
	viper.SetConfigName(ConfigName())
	viper.AddConfigPath(ConfigDirPath())
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No authentication credentials, please run sumocli configure")
		return ""
	} else {
		version := viper.GetString("version")
		accessId := viper.GetString("accessid")
		if version == "v1" {
			accessIdDecrypted := encryption.DecryptData(accessId)
			return accessIdDecrypted
		} else {
			return accessId
		}
	}
}

func ReadAccessKeys() (string, string, string) {
	viper.SetConfigName(ConfigName())
	viper.AddConfigPath(ConfigDirPath())
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No authentication credentials, please run sumocli configure")
		return "", "", ""
	} else {
		version := viper.GetString("version")
		accessId := viper.GetString("accessid")
		accessKey := viper.GetString("accesskey")
		endpoint := viper.GetString("endpoint")
		if version == "v1" {
			accessIdDecrypted := encryption.DecryptData(accessId)
			accessKeyDecrypted := encryption.DecryptData(accessKey)
			return accessIdDecrypted, accessKeyDecrypted, endpoint
		} else {
			return accessId, accessKey, endpoint
		}
	}
}

func ReadCredentials() (string, string) {
	var accessCredentialsComplete string
	var endpoint string
	viper.SetConfigName(ConfigName())
	viper.AddConfigPath(ConfigDirPath())
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		accessidenv := viper.GetString("SUMO_ACCESS_ID")
		accesskeyenv := viper.GetString("SUMO_ACCESS_KEY")
		endpointenv := viper.GetString("SUMO_ENDPOINT")

		accessCredentials := accessidenv + ":" + accesskeyenv
		accessCredentialsEnc := base64.StdEncoding.EncodeToString([]byte(accessCredentials))
		accessCredentialsComplete = "Basic " + accessCredentialsEnc
		return accessCredentialsComplete, endpointenv
	} else {
		version := viper.GetString("version")
		accessid := viper.GetString("accessid")
		accesskey := viper.GetString("accesskey")
		endpoint = viper.GetString("endpoint")
		// Determines if the access key/access id are encrypted at rest and need to be decrypted
		// before being used in requests.
		if version == "v1" {
			accessIdDecrypted := encryption.DecryptData(accessid)
			accessKeyDecrypted := encryption.DecryptData(accesskey)
			accessCredentials := accessIdDecrypted + ":" + accessKeyDecrypted
			accessCredentialsEnc := base64.StdEncoding.EncodeToString([]byte(accessCredentials))
			accessCredentialsComplete = "Basic " + accessCredentialsEnc
		} else {
			accessCredentials := accessid + ":" + accesskey
			accessCredentialsEnc := base64.StdEncoding.EncodeToString([]byte(accessCredentials))
			accessCredentialsComplete = "Basic " + accessCredentialsEnc
		}
	}
	return accessCredentialsComplete, endpoint
}

func ReadAuthCredentials() (string, string, string) {
	var endpoint string
	viper.SetConfigName(ConfigName())
	viper.AddConfigPath(ConfigDirPath())
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		accessIdEnv := viper.GetString("SUMO_ACCESS_ID")
		accessKeyEnv := viper.GetString("SUMO_ACCESS_KEY")
		endpointEnv := viper.GetString("SUMO_ENDPOINT")
		return accessIdEnv, accessKeyEnv, endpointEnv
	} else {
		version := viper.GetString("version")
		accessId := viper.GetString("accessid")
		accessKey := viper.GetString("accesskey")
		endpoint = viper.GetString("endpoint")
		// Determines if the access key/access id are encrypted at rest and need to be decrypted
		// before being used in requests.
		if version == "v1" {
			accessIdDecrypted := encryption.DecryptData(accessId)
			accessKeyDecrypted := encryption.DecryptData(accessKey)
			return accessIdDecrypted, accessKeyDecrypted, endpoint
		} else {
			return accessId, accessKey, endpoint
		}
	}
}
