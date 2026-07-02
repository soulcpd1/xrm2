package outbound

type Direct struct{}

func NewDirect() *Direct {
	return &Direct{}
}

func (d *Direct) Name() string {
	return "direct"
}

func (d *Direct) Up() error {
	return nil
}

func (d *Direct) Down() error {
	return nil
}

func (d *Direct) Ready() bool {
	return true
}

func (d *Direct) Forward(conn any) error {
	return nil
}
