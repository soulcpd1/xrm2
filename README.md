# XRM2 - eXtensible Routing Module

Маршрутизатор трафика с поддержкой SOCKS5, доменных и IP-правил.

## Возможности

- SOCKS5-прокси на `127.0.0.1:2080`
- Маршрутизация по доменам и IP
- Поддержка нескольких upstream (Direct, VPS)
- Загрузка правил из xray-конфига
- Горячая перезагрузка правил

## Установка

```bash
git clone https://github.com/soulcpd1/xrm2.git
cd xrm2
make build
sudo cp build/xrm /usr/local/bin/
var DefaultRules = []Rule{
    {Type: SuffixRule, Value: ".ru", Action: Direct},
    // ...
}
var LocalRules = []Rule{
    {Type: DomainRule, Value: "domain:max.ru", Action: Block},
    // ...
}
[Unit]
Description=XRM Runtime
After=network.target

[Service]
Type=simple
WorkingDirectory=/opt/xrm2
ExecStart=/usr/local/bin/xrm
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
# В конфиге Hysteria указать SOCKS5 upstream:
socks5:
  listen: 127.0.0.1:2080
# Через proxychains:
proxychains4 -f /etc/proxychains.conf ss-server -c /etc/shadowsocks/config.json
# Через iptables + redsocks:
iptables -t nat -A PREROUTING -s 10.100.0.0/24 -p tcp -j REDSOCKS_XRM
MIT
