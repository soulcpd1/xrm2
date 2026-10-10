#!/bin/bash
set -e

cd /home/soul/projects/xrm2

echo "=== Сборка xrm ==="
mkdir -p build
go build -o build/xrm ./cmd/xrm

echo "=== Остановка xrm ==="
sudo systemctl stop xrm

echo "=== Копирование бинарника ==="
sudo cp build/xrm /usr/local/bin/xrm

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
