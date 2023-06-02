#!/bin/sh
echo "Stopping the application..."
# Find the process IDs (PIDs) of the application
app_pids=$(pgrep -f "/home/server")

# Check if any application process is running
if [ -n "$app_pids" ]; then
  # Terminate each application process individually
  for app_pid in $app_pids; do
    sudo kill "$app_pid"
  done
  echo "Application stopped successfully."
else
  echo "Application is not running."
fi
