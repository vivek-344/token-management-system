#!/bin/bash

set -e
echo "Starting database migration..."

# Load env and run migrations
source /app/app.env
/app/migrate -path /app/db/migration -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@postgres:5432/${POSTGRES_DB}?sslmode=disable" -verbose up

echo "Migration completed successfully."