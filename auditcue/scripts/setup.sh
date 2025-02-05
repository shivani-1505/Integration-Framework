#!/bin/bash

# This line above is called a shebang - it tells the system this is a bash script

# First, let's make the script stop if any command fails
set -e

echo "Starting to create AuditCue project structure..."

# Create main application directories
echo "Creating main application directories..."
mkdir -p cmd/server
mkdir -p configs
mkdir -p scripts
mkdir -p docs

# Create internal directory structure
echo "Creating internal package directories..."
mkdir -p internal/auth
mkdir -p internal/connections/oauth
mkdir -p internal/models
mkdir -p internal/database/migrations
mkdir -p internal/signup
mkdir -p internal/types

echo "Project structure for AuditCue created successfully."