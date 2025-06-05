package config

import (
	"time"
)

type EnvConfig struct {
	App_Env   string
	App_Host  string
	App_Port  string
	SMTP_Host string
	SMTP_Port string
}

type JWTConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenTTL     time.Duration
	RefreshTokenTTL    time.Duration
}

type DatabaseConfig struct {
	Enabled              bool   `yaml:"Enabled"`
	Database_Name        string `yaml:"Database_Name"`
	Max_Idle_Connections int    `yaml:"Max_Idle_Connections"`
	Max_Open_Connections int    `yaml:"Max_Open_Connections"`
}

type SMTPConfig struct {
	Enabled           bool   `yaml:"Enabled"`
	Max_Retries       int    `yaml:"Max_Retries"`
	Use_TLS           bool   `yaml:"Use_TLS"`
	Base_Sender_Name  string `yaml:"Base_Sender_Name"`
	Base_Sender_Email string `yaml:"Base_Sender_Email"`
}

type LoggerConfig struct {
	Log_File_Path string `yaml:"Log_File_Path"`
	Log_Level     string `yaml:"Log_Level"`
	Compress_Logs bool   `yaml:"Compress_Logs"`
	Max_Size      int    `yaml:"Max_Size"`
	Max_Age       int    `yaml:"Max_Age"`
	Max_Backups   int    `yaml:"Max_Backup"`
}

type BigCache struct {
	Enabled            bool   `yaml:"Enabled"`
	Shards             int    `yaml:"Shards"`
	LifeWindow         string `yaml:"LifeWindow"`
	CleanWindow        string `yaml:"CleanWindow"`
	MaxEntriesInWindow int    `yaml:"MaxEntriesInWindow"`
	MaxEntrySize       int    `yaml:"MaxEntrySize"`
	Verbose            bool   `yaml:"Verbose"`
	HardMaxCacheSize   int    `yaml:"HardMaxCacheSize"`
}

type Config struct {
	Env       EnvConfig
	Database  DatabaseConfig `yaml:"Database"`
	SMTP      SMTPConfig     `yaml:"SMTP"`
	JwtConfig JWTConfig
	Logger    LoggerConfig `yaml:"Logging"`
	BigCache  BigCache     `yaml:"BigCache"`
}
