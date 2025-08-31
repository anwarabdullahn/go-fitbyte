#!/bin/bash
# Ensure .air directory exists
mkdir -p /workspace/.air
# Run air in background
echo "Start development..."
# Open a shell so VS Code terminal stays active
exec "$SHELL"
