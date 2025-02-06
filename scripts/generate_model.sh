#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
SERVER_DIR="${PROJECT_DIR}/go"
ENV_FILE="${PROJECT_DIR}/.env"

which dotenv > /dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "dotenv-cli がインストールされていません！"
    echo "dotenv-cli をインストールしてから再度実行してください。"
    echo "\`npm install -g dotenv-cli\`"
    exit 1
fi

docker compose ps | grep go >/dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "Docker コンテナが起動していません！"
    echo "Docker コンテナを起動してから再度実行してください。"
    exit 1
fi

cd "${SERVER_DIR}"

pwd

dotenv -e "${ENV_FILE}" -- go run "cmd/database/main.go"
