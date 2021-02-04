#!/usr/bin/env bash

# Set up the tweeter_api as a systemd service
if [ -f ./bin/application ]; then
  echo
  echo "Configuring tweeter_api service..."
  # Stop and Delete any pre-existing service
  sudo systemctl stop tweeter > /dev/null 2>&1
  sudo systemctl disable tweeter > /dev/null 2>&1
  # sudo useradd wordfreqservice -s /sbin/nologin -M > /dev/null 2>&1
  sudo cp tweeter.service /lib/systemd/system/
  sudo chmod 755 /lib/systemd/system/tweeter.service
  sudo systemctl enable tweeter.service
  echo "Starting tweeter service..."
  sudo systemctl start tweeter.service
  sudo systemctl is-active -q tweeter.service
  if [ $? -eq 0 ]; then
    echo "Service started successfully."
    echo
    echo "To see log output, type:"
    echo "sudo journalctl -f -u tweeter"
    echo "(CTRL+C to exit log reader)"
  else
    echo "Service could not be started. If this persists, please re-run setup."
  fi
else
  echo
  echo "No file at bin/application - run setup first."
fi
echo
