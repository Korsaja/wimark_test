#!/usr/bin/env bash

# shellcheck disable=SC2004
if (( $EUID != 0 )); then
    echo "Rights are required to run the installation as root"
    exit
fi
git clone https://github.com/Korsaja/wimark_test
# shellcheck disable=SC2164
cd wimark_test/
docker-compose build && docker-compose up -d
# shellcheck disable=SC2103
cd ..
rm -rf wimark_test/
echo "Done."
echo "for test use next commands:"
echo "curl http://localhost/hello"
echo "curl -iL http://localhost/cadvisor"
