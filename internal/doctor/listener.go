package doctor

import (
	"fmt"
	"net"
	"time"
)

type ListenerChecker struct{}

func (ListenerChecker) Run() Result {

	ports := []int{1080, 10808}

	var ok []int

	for _, port := range ports {

		conn, err := net.DialTimeout(
			"tcp",
			fmt.Sprintf("127.0.0.1:%d", port),
			500*time.Millisecond,
		)

		if err == nil {
			conn.Close()
			ok = append(ok, port)
		}
	}

	if len(ok) == 0 {
		return Result{
			Name:   "Local SOCKS",
			OK:     false,
			Detail: "no listeners",
		}
	}

	return Result{
		Name:   "Local SOCKS",
		OK:     true,
		Detail: fmt.Sprintf("%v", ok),
	}
}
