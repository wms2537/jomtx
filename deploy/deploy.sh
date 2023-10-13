#!/bin/sh
rm -rf desk/*
rm -rf sentry/*
rm -rf val/*
rm -rf public/*
rm -rf kms/*
    
docker run --rm -i \
    -v $(pwd)/desk:/root/.jomtx \
    wms2537/jomtxd_i \
    init jomtx

docker run --rm -i \
    -v $(pwd)/sentry:/root/.jomtx \
    wms2537/jomtxd_i \
    init jomtx

docker run --rm -i \
    -v $(pwd)/val:/root/.jomtx \
    wms2537/jomtxd_i \
    init jomtx

docker run --rm -i \
    -v $(pwd)/public:/root/.jomtx \
    wms2537/jomtxd_i \
    init jomtx

docker run --rm -i \
    -v $(pwd)/desk:/root/.jomtx \
    --entrypoint sed \
    wms2537/jomtxd_i \
    -Ei 's/^chain-id = .*$/chain-id = "jomtx-1"/g' \
    /root/.jomtx/config/client.toml

docker run --rm -i \
    -v $(pwd)/sentry:/root/.jomtx \
    --entrypoint sed \
    wms2537/jomtxd_i \
    -Ei 's/^chain-id = .*$/chain-id = "jomtx-1"/g' \
    /root/.jomtx/config/client.toml

docker run --rm -i \
    -v $(pwd)/val:/root/.jomtx \
    --entrypoint sed \
    wms2537/jomtxd_i \
    -Ei 's/^chain-id = .*$/chain-id = "jomtx-1"/g' \
    /root/.jomtx/config/client.toml

docker run --rm -i \
    -v $(pwd)/public:/root/.jomtx \
    --entrypoint sed \
    wms2537/jomtxd_i \
    -Ei 's/^chain-id = .*$/chain-id = "jomtx-1"/g' \
    /root/.jomtx/config/client.toml

mkdir ./desk/keys

openssl rand -base64 -out ./desk/keys/passphrase.txt 32

docker run --rm -it \
    -v $(pwd)/desk:/root/.jomtx \
    -v $(pwd)/password.sh:/password.sh \
    --entrypoint='' \
    wms2537/jomtxd_i \
    /bin/sh /password.sh

docker run --rm -it \
    -v $(pwd)/kms:/root/tmkms \
    wms2537/tmkms_i:v0.12.2 \
    init /root/tmkms

docker run --rm -i \
  -v $(pwd)/kms:/root/tmkms \
  --entrypoint sed \
  wms2537/tmkms_i:v0.12.2 \
  -Ei 's/^protocol_version = .*$/protocol_version = "v0.34"/g' \
  /root/tmkms/tmkms.toml

docker run --rm -i \
  -v $(pwd)/kms:/root/tmkms \
  --entrypoint sed \
  wms2537/tmkms_i:v0.12.2 \
  -Ei 's/path = "\/root\/tmkms\/secrets\/cosmoshub-3-consensus.key"/path = "\/root\/tmkms\/secrets\/val-consensus.key"/g' \
  /root/tmkms/tmkms.toml

docker run --rm -i \
    -v $(pwd)/kms:/root/tmkms \
    --entrypoint sed \
    wms2537/tmkms_i:v0.12.2 \
    -Ei 's/cosmoshub-3/jomtx-1/g' /root/tmkms/tmkms.toml

docker run --rm -t \
    -v $(pwd)/val:/root/.jomtx \
    wms2537/jomtxd_i \
    tendermint show-validator \
    | tr -d '\n' | tr -d '\r' \
    > desk/config/pub_validator_key-val.json

mv val/config/priv_validator_key.json \
  kms/secrets/priv_validator_key-val.json

docker run --rm -i \
    -v $(pwd)/kms:/root/tmkms \
    -w /root/tmkms \
    wms2537/tmkms_i:v0.12.2 \
    softsign import secrets/priv_validator_key-val.json \
    secrets/val-consensus.key

cp sentry/config/priv_validator_key.json \
    val/config/

docker run --rm -i \
    -v $(pwd)/kms:/root/tmkms \
    --entrypoint sed \
    wms2537/tmkms_i:v0.12.2 \
    -Ei 's/^addr = "tcp:.*$/addr = "tcp:\/\/val:26659"/g' /root/tmkms/tmkms.toml

docker run --rm -i \
  -v $(pwd)/val:/root/.jomtx \
  --entrypoint sed \
  wms2537/jomtxd_i \
  -Ei 's/priv_validator_laddr = ""/priv_validator_laddr = "tcp:\/\/0.0.0.0:26659"/g' \
  /root/.jomtx/config/config.toml

docker run --rm -i \
  -v $(pwd)/val:/root/.jomtx \
  --entrypoint sed \
  wms2537/jomtxd_i \
  -Ei 's/^priv_validator_key_file/# priv_validator_key_file/g' \
  /root/.jomtx/config/config.toml

docker run --rm -i \
  -v $(pwd)/val:/root/.jomtx \
  --entrypoint sed \
  wms2537/jomtxd_i \
  -Ei 's/^priv_validator_state_file/# priv_validator_state_file/g' \
  /root/.jomtx/config/config.toml

cp sentry/config/priv_validator_key.json \
    val/config

docker run --rm -i \
    -v $(pwd)/desk:/root/.jomtx \
    --entrypoint sed \
    wms2537/jomtxd_i \
    -Ei 's/"chain_id": "jomtx"/"chain_id": "jomtx-1"/g' \
    /root/.jomtx/config/genesis.json

ALICE=$(cat ./desk/keys/passphrase.txt | docker run --rm -i \
    -v $(pwd)/desk:/root/.jomtx \
    wms2537/jomtxd_i \
    keys \
    --keyring-backend file --keyring-dir /root/.jomtx/keys \
    show alice --address)

docker run --rm -it \
    -v $(pwd)/desk:/root/.jomtx \
    wms2537/jomtxd_i \
    add-genesis-account $ALICE 1000000000000000stake,1000000000000000token

# TODO: Prebuild genesis

cat ./desk/keys/passphrase.txt | docker run --rm -i \
    -v $(pwd)/desk:/root/.jomtx \
    wms2537/jomtxd_i \
    gentx alice 60000000stake \
    --keyring-backend file --keyring-dir /root/.jomtx/keys \
    --account-number 0 --sequence 0 \
    --pubkey $(cat desk/config/pub_validator_key-val.json) \
    --chain-id jomtx-1 \
    --gas 1000000 \
    --gas-prices 0.1stake

docker run --rm -it \
    -v $(pwd)/desk:/root/.jomtx \
    wms2537/jomtxd_i collect-gentxs

docker run --rm -it \
    -v $(pwd)/desk:/root/.jomtx \
    wms2537/jomtxd_i \
    validate-genesis

cp desk/config/genesis.json sentry/config
cp desk/config/genesis.json val/config
cp desk/config/genesis.json public/config

VAL_ID=$(docker run --rm -i \
    -v $(pwd)/val:/root/.jomtx \
    wms2537/jomtxd_i \
    tendermint show-node-id)

SENTRY_ID=$(docker run --rm -i \
    -v $(pwd)/sentry:/root/.jomtx \
    wms2537/jomtxd_i \
    tendermint show-node-id)

PUBLIC_ID=$(docker run --rm -i \
    -v $(pwd)/public:/root/.jomtx \
    wms2537/jomtxd_i \
    tendermint show-node-id)

docker run --rm -i \
  -v $(pwd)/sentry:/root/.jomtx \
  -e VAL_ID=$VAL_ID \
  --entrypoint sed \
  wms2537/jomtxd_i \
  -Ei "s/^persistent_peers = \"\"/persistent_peers = \"$VAL_ID@val:26656\"/" \
  /root/.jomtx/config/config.toml

docker run --rm -i \
  -v $(pwd)/sentry:/root/.jomtx \
  -e VAL_ID=$VAL_ID \
  --entrypoint sed \
  wms2537/jomtxd_i \
  -Ei "s/^private_peer_ids = \"\"/private_peer_ids = \"$VAL_ID\"/" \
  /root/.jomtx/config/config.toml

docker run --rm -i \
  -v $(pwd)/val:/root/.jomtx \
  -e SENTRY_ID=$SENTRY_ID \
  --entrypoint sed \
  wms2537/jomtxd_i \
  -Ei "s/^persistent_peers = \"\"/persistent_peers = \"$SENTRY_ID@sentry:26656\"/" \
  /root/.jomtx/config/config.toml

docker run --rm -i \
  -v $(pwd)/public:/root/.jomtx \
  -e SENTRY_ID=$SENTRY_ID \
  --entrypoint sed \
  wms2537/jomtxd_i \
  -Ei "s/^seeds = \"\"/seeds = \"$SENTRY_ID@sentry:26656\"/" \
  /root/.jomtx/config/config.toml

docker run --rm -i \
  -v $(pwd)/public:/root/.jomtx \
  --entrypoint sed \
  wms2537/jomtxd_i \
  -Ei 's/^laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' \
  /root/.jomtx/config/config.toml

docker run --rm -i \
  -v $(pwd)/public:/root/.jomtx \
  --entrypoint sed \
  wms2537/jomtxd_i \
  -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' \
  /root/.jomtx/config/config.toml

docker run --rm -i \
  -v $(pwd)/public:/root/.jomtx \
  --entrypoint sed \
  wms2537/jomtxd_i \
  -Ei "s/^enabled-unsafe-cors = .*$/enabled-unsafe-cors = true/g" \
  /root/.jomtx/config/config.toml

docker compose up -d

docker run --rm -it \
    -p 4500:4500 \
    --name jomtx-faucet \
    --network jomtx_net-public \
    --detach \
    wms2537/jomtx_faucet_i start http://public:26657