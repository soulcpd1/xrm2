package xray

import (
	"errors"

	"github.com/google/uuid"
)

func (c *Config) AddClient(email string) error {

	ib := c.RealityInbound()
	if ib == nil {
		return errors.New("reality inbound not found")
	}

	for _, cl := range ib.Settings.Clients {
		if cl.Email == email {
			return errors.New("client already exists: " + email)
		}
	}

	ib.Settings.Clients = append(ib.Settings.Clients, Client{
		ID:    uuid.NewString(),
		Email: email,
		Flow:  "xtls-rprx-vision",
	})

	return nil
}
