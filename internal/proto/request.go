package proto

import (
	"encoding/binary"
	"io"
)

type Request struct {
	Network string
	Address string
}

func WriteRequest(w io.Writer, r Request) error {
	if err := binary.Write(w, binary.BigEndian, uint16(len(r.Network))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(r.Network)); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint16(len(r.Address))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(r.Address)); err != nil {
		return err
	}

	return nil
}

func ReadRequest(r io.Reader) (Request, error) {
	var req Request

	var n uint16

	if err := binary.Read(r, binary.BigEndian, &n); err != nil {
		return req, err
	}

	buf := make([]byte, n)
	if _, err := io.ReadFull(r, buf); err != nil {
		return req, err
	}
	req.Network = string(buf)

	if err := binary.Read(r, binary.BigEndian, &n); err != nil {
		return req, err
	}

	buf = make([]byte, n)
	if _, err := io.ReadFull(r, buf); err != nil {
		return req, err
	}
	req.Address = string(buf)

	return req, nil
}
