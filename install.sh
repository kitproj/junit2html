#!/usr/bin/env sh
set -eux

tag=$(curl --retry 3 -s "https://api.github.com/repos/kitproj/junit2html/releases/latest" | jq -r '.tag_name')
version=$(echo $tag | cut -c 2-)
url="https://github.com/kitproj/junit2html/releases/download/${tag}/kit_${version}_$(uname)_$(uname -m | sed 's/aarch64/arm64/').tar.gz"
curl --retry 3 -L $url | tar -zxvf - junit2html
chmod +x junit2html
sudo mv junit2html /usr/local/bin/junit2html
