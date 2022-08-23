#!/bin/bash

echo "$ID_RSA" | base64 --decode > "$HOME/.ssh/id_rsa"
echo "$ID_RSA_PUB" | base64 --decode > "$HOME/.ssh/id_rsa.pub"

chmod 400 "$HOME/.ssh/id_rsa"
chmod 400 "$HOME/.ssh/id_rsa.pub"

ssh-keyscan -H $SSH_HOST >> $HOME/.ssh/known_hosts

"$HOME"/dist/app
