#!/bin/bash
docker exec -it discord_clone_server /bin/sh -c "go test -v -p 1 ./... -coverprofile=coverage.out"
# docker exec -it discord_clone_server /bin/sh -c "go test -v -p 1 ./..."
