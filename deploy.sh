#!/bin/bash
cd /home/soul/projects/xrm2
make build
sudo systemctl stop xrm
sudo cp build/xrm /usr/local/bin/xrm
sudo systemctl start xrm
sleep 2
sudo systemctl restart sing-box
sudo systemctl status xrm --no-pager | head -3
sudo systemctl status sing-box --no-pager | head -3
