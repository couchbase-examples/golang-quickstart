#!/bin/bash

# Check if .env file exists
if [ -f .env ]; then
  # Load .env file
  export $(cat .env | grep -v '^#' | xargs)
else
  echo ".env file not found"
fi