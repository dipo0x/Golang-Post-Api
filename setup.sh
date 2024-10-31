#!/bin/bash

# Prompt the user for server and database details
echo "Enter server post (default: localhost:8080):"
read -r PORT
PORT=${PORT:-localhost:8080}

echo "Enter database host (default: mongodb://localhost:27017/?connect=direct):"
read -r MONGO_URI
MONGO_URI=${MONGO_URI:-mongodb://localhost:27017/?connect=direct}

echo "Enter database name (default: golang-rest-api):"
read -r MONGO_DATABASE
MONGO_DATABASE=${MONGO_DATABASE:-golang-rest-api}

# Create the .env file
cat <<EOF > .env
PORT=$PORT
MONGO_URI=$MONGO_URI
MONGO_DATABASE=$MONGO_DATABASE
EOF

echo ".env file created successfully."