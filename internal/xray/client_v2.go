package xray

import (
	"errors"

	"github.com/google/uuid"
)

type ClientMutation struct {
	ID    string
	Email string
	Flow  string
}

func (cfg *Config) RealityInbound() *Inbound {
	for i := range cfg.Inbounds {
		inb := &cfg.Inbounds[i]

		if inb.StreamSettings.Security == "reality" {
			return inb
		}
	}

	return nil
}

func (cfg *Config) FindClient(email string) *Client {
	inb := cfg.RealityInbound()
	if inb == nil {
		return nil
	}

	for i := range inb.Settings.Clients {
		if inb.Settings.Clients[i].Email == email {
			return &inb.Settings.Clients[i]
		}
	}

	return nil
}

func (cfg *Config) AddClient(m ClientMutation) error {

	inb := cfg.RealityInbound()
	if inb == nil {
		return errors.New("reality inbound not found")
	}

	if cfg.FindClient(m.Email) != nil {
		return errors.New("client already exists: " + m.Email)
	}

	if m.ID == "" {
		m.ID = uuid.NewString()
	}

	inb.Settings.Clients = append(inb.Settings.Clients, Client{
		ID:    m.ID,
		Email: m.Email,
		Flow:  m.Flow,
	})

	return nil
}

func (cfg *Config) RemoveClient(email string) error {

	inb := cfg.RealityInbound()
	if inb == nil {
		return errors.New("reality inbound not found")
	}

	for i, c := range inb.Settings.Clients {

		if c.Email != email {
			continue
		}

		inb.Settings.Clients = append(
			inb.Settings.Clients[:i],
			inb.Settings.Clients[i+1:]...,
		)

		return nil
	}

	return errors.New("client not found: " + email)
}
