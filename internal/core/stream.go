package core

import (
	"context"
	"io"
	"net"
)

type StreamBridge struct {
	forwarder *Forwarder
}

func NewStreamBridge(f *Forwarder) *StreamBridge {
	return &StreamBridge{forwarder: f}
}

func (b *StreamBridge) Bridge(
	ctx context.Context,
	driver string,
	inbound net.Conn,
	network string,
	address string,
) error {

	out, err := b.forwarder.Dial(ctx, driver, network, address)
	if err != nil {
		return err
	}

	errCh := make(chan error, 2)

	go func() {
		_, err := io.Copy(out, inbound)
		errCh <- err
	}()

	go func() {
		_, err := io.Copy(inbound, out)
		errCh <- err
	}()

	select {
	case <-ctx.Done():
		_ = inbound.Close()
		_ = out.Close()
		return ctx.Err()
	case err := <-errCh:
		_ = inbound.Close()
		_ = out.Close()
		return err
	}
}

func (b *StreamBridge) Dial(
	ctx context.Context,
	driver string,
	network string,
	address string,
) (net.Conn, error) {

	return b.forwarder.Dial(
		ctx,
		driver,
		network,
		address,
	)
}
