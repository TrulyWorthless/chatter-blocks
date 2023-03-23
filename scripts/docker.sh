#!/bin/sh
docker run --name chatterBlocksDb -p 5432:5432 -e POSTGRES_USER=postgresUser -e POSTGRES_PASSWORD=postgresPW -e POSTGRES_DB=postgresDB -d postgres
docker exec myPostgresDb env