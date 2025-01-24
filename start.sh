#!/bin/bash

set -e
echo "Starting the service..."

# Load env variables
source /app/app.env

# Start the main application
/app/main