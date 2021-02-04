#!/usr/bin/env bash

# Set up the GOLANG environment
echo "Installing the GO language environment..."

cd /tmp || exit
wget https://golang.org/dl/go1.15.8.linux-amd64.tar.gz
tar -xvf go1.15.8.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo mv go /usr/local/
cd - || exit
echo

# Copy over and run the profile for environment variables
cp -f go-profile ~/.profile
if [ -f ~/.bashrc ]; then cat go-profile >> ~/.bashrc; fi
# shellcheck source=/dev/null
source ~/.profile
echo


#  location / {
#                 proxy_pass http://127.0.0.1:3000;
#         }

echo "Building tweeter_api bin/application..."
go get -u github.com/gin-gonic/gin
rm -f ./bin/application > /dev/null 2>&1
go build -o bin/application ./main.go
if [ -f ./bin/application ]; then echo "Setup complete."; fi
echo
