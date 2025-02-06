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

MYSQL_USER=$(cat ${ENV_FILE} | grep MYSQL_USER | cut -d '=' -f 2)
MYSQL_PASSWORD=$(cat ${ENV_FILE} | grep MYSQL_PASSWORD | cut -d '=' -f 2)
MYSQL_HOST=$(cat ${ENV_FILE} | grep MYSQL_HOST | cut -d '=' -f 2)
MYSQL_PORT=$(cat ${ENV_FILE} | grep MYSQL_PORT | cut -d '=' -f 2)
MYSQL_DATABASE=$(cat ${ENV_FILE} | grep MYSQL_DATABASE | cut -d '=' -f 2) 

dotenv -e "${ENV_FILE}" -- docker compose exec go migrate -path sql/migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" -verbose up
