package model

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"twodo.app/condo/utility"
)

type ConfigInfo struct {
	Database   DatabaseConfigInfo `toml:"database"`
	Smtp       SmtpConfigInfo     `toml:"smtp"`
	CondoEmail string             `toml:"condo_email"`
}

func LoadConfig() ConfigInfo {
	var config ConfigInfo
	_, err := toml.DecodeFile(utility.ConfigPath, &config)
	if err != nil {
		panic(err)
	}
	return config
}

type DatabaseConfigInfo struct {
	DbName   string `toml:"db_name"`
	Host     string `toml:"host"`
	Password string `toml:"password"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
}

type SmtpConfigInfo struct {
	From     string `toml:"from"`
	Host     string `toml:"host"`
	Password string `toml:"password"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
}

type ErrorInfo struct {
	Message string `json:"message"`
}

func NewErrorInfo(err error) ErrorInfo {
	err = fmt.Errorf("%v", err)
	return ErrorInfo{Message: err.Error()}
}
