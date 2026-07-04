package hycore

import (
	"crypto/tls"
	"net"
	"time"

	hyserver "github.com/apernet/hysteria/core/v2/server"

	"github.com/soulcpd1/xrm2/internal/hysteria"
)

func BuildConfig(cfg *hysteria.Config) (*hyserver.Config, error) {
	cert, err := tls.LoadX509KeyPair(
		cfg.TLS.Cert,
		cfg.TLS.Key,
	)
	if err != nil {
		return nil, err
	}

	pc, err := net.ListenPacket("udp", cfg.Listen)
	if err != nil {
		return nil, err
	}

	return &hyserver.Config{
		Conn: pc,
		TLSConfig: hyserver.TLSConfig{
			Certificates: []tls.Certificate{cert},
		},
		QUICConfig: hyserver.QUICConfig{
			InitialStreamReceiveWindow:     cfg.QUIC.InitStreamReceiveWindow,
			MaxStreamReceiveWindow:         cfg.QUIC.MaxStreamReceiveWindow,
			InitialConnectionReceiveWindow: cfg.QUIC.InitConnectionReceiveWindow,
			MaxConnectionReceiveWindow:     cfg.QUIC.MaxConnectionReceiveWindow,
			MaxIdleTimeout:                 30 * time.Second,
		},
		IgnoreClientBandwidth: cfg.IgnoreClient,
		Authenticator:         NewAuthenticator(&cfg.Auth),
	}, nil
}
