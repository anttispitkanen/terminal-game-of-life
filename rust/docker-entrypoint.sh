#!/bin/sh
# Sleep is needed as a hacky fix for a Docker tty bug, see README
sleep 1; ./rust "$@"