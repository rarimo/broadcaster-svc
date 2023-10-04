package config

import (
	"github.com/rarimo/broadcaster-svc/internal/data"
	"github.com/rarimo/broadcaster-svc/internal/data/pg"
)

func (c *config) Storage() data.Storage {
	return c.storage.Do(func() interface{} {
		return pg.New(c.DB())
	}).(data.Storage)
}
