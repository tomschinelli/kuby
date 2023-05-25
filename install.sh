#!/usr/bin/env bash
set -e
go build -o bin/kummy cmd/kummy.go
sudo cp ./bin/kummy /usr/local/bin/kummy

echo """
done.

To verify the installation, run


  kummy help


To enable completion in current shell session, run


  source <(kummy completion $SHELL)


verify completion is working:


  kummy <tab><tab>


"""