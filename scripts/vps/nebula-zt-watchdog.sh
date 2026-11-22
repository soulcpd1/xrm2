#!/bin/bash
LOG="/var/log/nebula-zt-watchdog.log"
NEBULA_PEER="192.168.100.2"
ZT_PEER="10.147.17.73"
PING_COUNT=2
PING_TIMEOUT=3
MAX_DOWN_TIME=300

log() { echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1" >> "$LOG"; }

check_peer() {
    # Явно указываем интерфейс, чтобы ping шёл строго через туннель
    if [ "$1" == "$NEBULA_PEER" ]; then
        ping -I nebula1 -c "$PING_COUNT" -W "$PING_TIMEOUT" "$1" > /dev/null 2>&1
    else
        ping -I ztss4sxsjb -c "$PING_COUNT" -W "$PING_TIMEOUT" "$1" > /dev/null 2>&1
    fi
}

check_and_restart() {
    local peer="$1"
    local service="$2"
    local down_file="/var/run/${service}-down-since"

    if check_peer "$peer"; then
        [ -f "$down_file" ] && rm -f "$down_file"
        return 0
    fi

    local now=$(date +%s)
    if [ ! -f "$down_file" ]; then
        echo "$now" > "$down_file"
        log "$service channel DOWN (peer $peer unreachable), starting 5-min grace period"
        return 0
    fi

    local down_since=$(cat "$down_file")
    local down_duration=$((now - down_since))

    if [ "$down_duration" -lt "$MAX_DOWN_TIME" ]; then
        log "$service still DOWN for ${down_duration}s, waiting (grace period)"
        return 0
    fi

    log "$service DOWN for ${down_duration}s (exceeded 5min). Restarting $service..."
    systemctl restart "$service"
    sleep 5

    if check_peer "$peer"; then
        log "$service RESTORED after restart"
        rm -f "$down_file"
    else
        log "ERROR: $service still DOWN after restart"
    fi
}

check_and_restart "$NEBULA_PEER" "nebula"
check_and_restart "$ZT_PEER" "zerotier-one"
