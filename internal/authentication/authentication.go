package authentication

import (
	"encoding/base64"
	"fmt"
	"github.com/SumoLogic-Incubator/sumocli/internal/encryption"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func ConfigPath() string {
	var filePath = ".sumocli/credentials/creds.json"
	homeDirectory, _ := os.UserHomeDir()
	configFile := filepath.Join(homeDirectory, filePath)
	return configFile
}

func ReadAccessId() string {
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(ConfigPath()))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No authentication credentials, please run sumocli login")
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
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(ConfigPath()))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No authentication credentials, please run sumocli login")
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
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(ConfigPath()))
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
	viper.SetConfigName("creds")
	viper.AddConfigPath(filepath.Dir(ConfigPath()))
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
