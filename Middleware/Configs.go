package Middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var DBConfig DatabaseConfig
var SmsConfig SmsApiConfig

type DatabaseConfig struct {
	MysqlAddress string `json:"mysql_address"`
	MysqlUsername string `json:"mysql_username"`
	MysqlPassword string `json:"mysql_password"`
	MysqlDatabase string `json:"mysql_database"`
	MysqlPort string `json:"mysql_port"`
}

func (config DatabaseConfig) GeneratePath () string {
	return strings.Join([]string{
		config.MysqlUsername, ":", config.MysqlPassword,
		"@tcp(", config.MysqlAddress, ":", config.MysqlPort, ")/",
		config.MysqlDatabase}, "")
}

func (config* DatabaseConfig) LoadConfig (path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		jsonStr := string(data)
		err = json.Unmarshal([]byte(jsonStr), config)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

type SmsApiConfig struct {
	RegionId string `json:"region_id"`
	AccessKeyId string `json:"access_key_id"`
	AccessSecret string `json:"access_secret"`
}

func (config* SmsApiConfig) LoadConfig (path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		jsonStr := string(data)
		err = json.Unmarshal([]byte(jsonStr), config)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

