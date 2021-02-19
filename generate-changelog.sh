#!/bin/bash
version=v0.0.4-dev.1
curl -s https://api.github.com/repos/barelyhuman/commitlog/releases/tags/$version | grep browser_download_url | grep linux-amd64 | cut -d '"' -f 4 | wget -qi -
tar -xvzf commitlog-linux-amd64.tar.gz
chmod +x commitlog
./commitlog -i fix,feat,refactor > CHANGELOG.txt