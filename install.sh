#!/bin/bash
set -uo pipefail

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

CONNECT_TIMEOUT=10
MAX_TIME=60
RETRIES=3

# fetch <url> <outfile-or-'-'>: downloads with bounded timeouts/retries so a
# flaky network fails fast instead of hanging indefinitely.
fetch() {
    local url="$1"
    local out="$2"
    if command -v curl > /dev/null 2>&1; then
        curl -fsSL \
            --connect-timeout "$CONNECT_TIMEOUT" \
            --max-time "$MAX_TIME" \
            --retry "$RETRIES" \
            --retry-connrefused \
            -H "User-Agent: glee-installer" \
            -o "$out" "$url"
    elif command -v wget > /dev/null 2>&1; then
        wget -nv \
            --timeout="$CONNECT_TIMEOUT" \
            --tries="$RETRIES" \
            --header="User-Agent: glee-installer" \
            -O "$out" "$url"
    else
        echo -e "${RED}Neither curl nor wget is available. Please install one and retry.${NC}"
        exit 1
    fi
}

get_file() {
    dl_url="https://api.github.com/repos/HexmosTech/glee/releases/latest"
    echo "Fetching latest release metadata from GitHub..."
    api_resp=$(fetch "$dl_url" -)
    if [ $? -ne 0 ] || [ -z "$api_resp" ]; then
        echo -e "${RED}Could not reach GitHub (network timeout, or GitHub is unreachable from your network).${NC}"
        echo -e "${YELLOW}Check your connection/VPN/proxy and try again.${NC}"
        exit 1
    fi
}

get_platform() {
    architecture=""
    case $(uname -m) in
    i386) architecture="386" ;;
    i686) architecture="386" ;;
    x86_64) architecture="amd64" ;;
    arm) command -v dpkg > /dev/null 2>&1 && dpkg --print-architecture | grep -q "arm64" && architecture="arm64" || architecture="arm" ;;
    arm64) architecture="arm64" ;;
    aarch64) architecture="arm64" ;;
    *)
        echo -e "${RED}Unsupported architecture: $(uname -m)${NC}"
        exit 1
        ;;
    esac
}

get_os() {
    the_os=""
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        echo "OS is Linux"
        the_os="linux"
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        # Mac OSX
        echo "OS is Mac OSX"
        the_os="darwin"
    elif [[ "$OSTYPE" == "cygwin" ]]; then
        # POSIX compatibility layer and Linux environment emulation for Windows
        echo "OS is Cygwin"
        echo "Installer not supported yet; please use release binary"
        exit 1
    elif [[ "$OSTYPE" == "msys" ]]; then
        # Lightweight shell and GNU utilities compiled for Windows (part of MinGW)
        echo "OS is msys"
        echo "Installer not supported yet; please use release binary"
        exit 1
    elif [[ "$OSTYPE" == "win32" ]]; then
        # I'm not sure this can happen.
        echo "OS is win32"
        echo "Installer not supported yet; please use release binary"
        exit 1
    elif [[ "$OSTYPE" == "freebsd"* ]]; then
        # ...
        echo "OS is freebsd"
        echo "Installer not supported yet; please use release binary"
        exit 1
    else
        # Unknown.
        echo "Error: Unknown OS"
        exit 1
    fi
}


get_file
get_platform
get_os
suffix="${the_os}-${architecture}.tar.gz"

# find_asset_url <suffix>: extracts the browser_download_url whose value ends
# in <suffix> (e.g. "linux-amd64.tar.gz", not "...tar.gz.md5"). The old
# implementation grepped with an end-of-line anchor, which only worked when
# GitHub pretty-printed one JSON field per line; GitHub now returns the API
# response as a single minified line, so that anchor never matched and the
# script silently ended up with an empty URL. jq/python3 parse the JSON
# properly regardless of formatting; the final fallback normalizes commas to
# newlines to approximate one-field-per-line before matching.
find_asset_url() {
    local suf="$1"
    if command -v jq > /dev/null 2>&1; then
        echo "${api_resp}" | jq -r --arg suf "$suf" \
            '.assets[]? | select(.browser_download_url | endswith($suf)) | .browser_download_url' \
            | head -n1
    elif command -v python3 > /dev/null 2>&1; then
        echo "${api_resp}" | python3 -c '
import json, sys
suf = sys.argv[1]
try:
    data = json.load(sys.stdin)
except ValueError:
    sys.exit(0)
for asset in data.get("assets", []):
    url = asset.get("browser_download_url", "")
    if url.endswith(suf):
        print(url)
        break
' "$suf"
    else
        echo "${api_resp}" | tr ',' '\n' | grep "\"browser_download_url\":\"[^\"]*${suf}\"" \
            | sed -E 's/.*"browser_download_url":"([^"]*)".*/\1/' | head -n1
    fi
}

archive=$(find_asset_url "$suffix")

if [ -z "$archive" ]; then
    echo -e "${RED}Could not find a release asset matching ${the_os}-${architecture}.${NC}"
    echo -e "${YELLOW}See available assets at: https://github.com/HexmosTech/glee/releases/latest${NC}"
    exit 1
fi

echo "Downloading glee from ${archive}"
fetch "${archive}" /tmp/glee_latest.tar.gz
if [ $? -ne 0 ] || [ ! -s /tmp/glee_latest.tar.gz ]; then
    echo -e "${RED}Download failed or the network timed out. Please check your connection and try again.${NC}"
    exit 1
fi

tar -xzf /tmp/glee_latest.tar.gz -C /tmp
sudo rm -f /usr/local/bin/glee /usr/bin/glee
if ! sudo mv /tmp/glee /usr/local/bin/glee; then
    echo -e "${RED}Could not move glee into /usr/local/bin (sudo failed or was denied).${NC}"
    exit 1
fi

if [ -x /usr/local/bin/glee ] && command -v glee > /dev/null 2>&1; then
    echo -e "${GREEN}Successfully installed glee; Type 'glee <markdown_file>' to invoke glee${NC}"
else
    echo -e "${RED}Failure in installation; please report issue at github.com/HexmosTech/glee${NC}"
    exit 1
fi
