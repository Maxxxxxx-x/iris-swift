package config

import "time"

type AppConfig struct {
	Host string
	Port string
}

type JWTTokenConfig struct {
	Secret string
	TTL    time.Duration
}

type JWTConfig struct {
	AccessToken        JWTTokenConfig
	RefreshToken       JWTTokenConfig
	VerifyEmailToken   JWTTokenConfig
	ResetPasswordToken JWTTokenConfig
}

type DatabaseConfig struct {
	Enabled              bool   `yaml:"Enabled"`
	Database_Name        string `yaml:"Database_Name"`
	Max_Idle_Connections int    `yaml:"Max_Idle_Connections"`
	Max_Open_Connections int    `yaml:"Max_Open_Connections"`
}

type SMTPConfig struct {
	Enabled           bool `yaml:"Enabled"`
	SMTP_Host         string
	SMTP_Port         string
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

type Config struct {
	Env       string
	App       AppConfig
	Database  DatabaseConfig `yaml:"Database"`
	SMTP      SMTPConfig     `yaml:"SMTP"`
	JwtConfig JWTConfig
	Logger    LoggerConfig `yaml:"Logging"`
}
