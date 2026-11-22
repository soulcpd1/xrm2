#!/bin/bash
set -e

cd /home/soul/projects/xrm2

echo "=== Сборка xrm ==="
mkdir -p build
go build -o build/xrm ./cmd/xrm

echo "=== Сборка xrm-agent ==="
go build -o build/xrm-agent ./cmd/xrm-agent

echo "=== Остановка xrm ==="
sudo systemctl stop xrm

echo "=== Копирование бинарника xrm ==="
sudo cp build/xrm /usr/local/bin/xrm

echo "=== Копирование бинарника xrm-agent на VPS ==="
scp build/xrm-agent root@31.14.138.11:/tmp/xrm-agent
ssh root@31.14.138.11 "sudo mv /tmp/xrm-agent /usr/local/bin/xrm-agent && sudo chmod +x /usr/local/bin/xrm-agent && sudo systemctl restart xrm-agent"

echo "=== Запуск xrm ==="
sudo systemctl start xrm
sleep 2

echo "=== Перезапуск sing-box и sing-box-anytls ==="
sudo systemctl restart sing-box
sudo systemctl restart sing-box-anytls
sleep 2

echo "=== Статус ==="
sudo systemctl status xrm --no-pager | head -3
sudo systemctl status sing-box --no-pager | head -3
sudo systemctl status sing-box-anytls --no-pager | head -3

echo ""
echo "=== Порты ==="
sudo ss -tlnp | grep -E ':(443|2080|8446)' || true
