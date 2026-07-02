package xray

import "github.com/google/uuid"

type ClientMutation struct {
	ID    string
	Email string
	Flow  string
}

func AppendRealityClient(cfg *Config, client ClientMutation) error {

	if client.ID == "" {
		client.ID = uuid.NewString()
	}

	inb := &cfg.Inbounds[0] // reality inbound (упрощённо v2 stage 1)

	cfg.Inbounds[0].Settings.Clients = append(inb.Settings.Clients, Client{
		ID:    client.ID,
		Email: client.Email,
		Flow:  client.Flow,
	})

	return nil
}
