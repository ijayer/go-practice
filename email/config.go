package main

type EmailConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

func NewEmailConfig(host string, port int, username, password string) *EmailConfig {
	return &EmailConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}
