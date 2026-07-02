package xray

func sanitizeReality(rs *RealitySettings) {
	if rs.ServerName == "" {
		rs.ServerName = "www.sberbank.ru"
	}
}
