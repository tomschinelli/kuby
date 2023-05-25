#!/usr/bin/env bash
set -e
go build -o bin/kummy cmd/kummy.go
sudo cp ./bin/kummy /usr/local/bin/kummy

echo """
done.

Run the following command to enable completion in current session:

  source <(kummy completion zsh)

Then try run

  kummy <tab><tab>

"""