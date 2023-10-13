#!/bin/sh

# Read the password from the file
password=$(cat /root/.jomtx/keys/passphrase.txt)

# Run jomtxd keys add and provide the password twice
{ echo "$password"; sleep 1; echo "$password"; } | jomtxd keys --keyring-backend file --keyring-dir /root/.jomtx/keys add alice"
