#!/bin/bash
docker exec -it discord_clone_server /bin/sh -c "go run main.go migrate up"
