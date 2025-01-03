package postgres

import "fmt"

type Config struct {
	DBName             string
	Hostname           string
	Port               int
	Username           string
	Password           string
	MaxOpenConnections int
	MaxIdleConnections int
}

func (p Config) formatDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.Hostname, p.Port, p.Username, p.Password, p.DBName)
}
