package postgres_connection

import (
	"fmt"
	"os"
)

type Options struct {
	databaseName *string
	host         *string
	port         *int
	user         *string
	password     *string
}

func Config() *Options {
	return new(Options)
}

func (c *Options) DatabaseName(databaseName string) *Options {
	c.databaseName = &databaseName
	return c
}

func (c *Options) Host(host string) *Options {
	c.host = &host
	return c
}

func (c *Options) Port(port int) *Options {
	c.port = &port
	return c
}

func (c *Options) User(user string) *Options {
	c.user = &user
	return c
}

func (c *Options) Password(password string) *Options {
	c.password = &password
	return c
}

func MergeOptions(opts ...*Options) *Options {
	option := new(Options)

	for _, opt := range opts {
		if opt.databaseName != nil {
			option.databaseName = opt.databaseName
		}
		if opt.host != nil {
			option.host = opt.host
		}
		if opt.port != nil {
			option.port = opt.port
		}
		if opt.user != nil {
			option.user = opt.user
		}
		if opt.password != nil {
			option.password = opt.password
		}
	}
	return option
}

var (
	defaultPort = 5432
)

func (d *Options) GetDsnConnection() string {
	DsnFormat := "host=%v user=%v password=%v dbname=%v port=%v"

	if d.port == nil {
		d.port = &defaultPort
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" || environment == "" {
		DsnFormat = "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable"
	}

	// log.Info("Connection: %s", fmt.Sprintf(DsnFormat, *d.host, *d.user, "************", *d.databaseName, *d.port))
	return fmt.Sprintf(DsnFormat, *d.host, *d.user, *d.password, *d.databaseName, *d.port)
}
