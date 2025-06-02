# !/bin/bash

cd deployment/docker

echo "start copy env files"
bash copy-env.sh

echo "start copy k8s secret files"
bash copy-k8s-secret.sh

sleep 2

echo "start build docker containers"
docker compose up -d

sleep 5

echo "start migrate"
bash migrate-up.sh

echo "start seed"
bash seed.sh
