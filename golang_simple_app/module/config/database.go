package config

import (
	"fmt"
)

type DatabaseConfig struct {
	Driver      string `default:"postgres" envconfig:"DRIVER"`
	Host        string `default:"127.0.0.1" envconfig:"HOST"`
	Port        int    `default:"5432" envconfig:"PORT"`
	Username    string `default:"postgres" envconfig:"USERNAME"`
	Password    string `default:"postgres" envconfig:"PASSWORD"`
	Database    string `default:"sku_management" envconfig:"DATABASE"`
	QueryString string `default:"sslmode=disable" envconfig:"QUERYSTRING"`
}

func (c *DatabaseConfig) RWDataSourceName() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?%s",
		c.Driver,
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.QueryString,
	)
}
