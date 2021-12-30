#!/bin/bash

# https://www.cnblogs.com/devhg/p/15435362.html

CA_PATH="ca"
SERVER_PATH="server"
CLIENT_PATH="client"

rm -R ${CA_PATH} ${SERVER_PATH} ${CLIENT_PATH}
mkdir ${CA_PATH} ${SERVER_PATH} ${CLIENT_PATH}

# CA
echo -e "创建 CA 私钥."
openssl genrsa -out ${CA_PATH}/ca.key 2048

echo -e "创建 CA 证书(含公钥)."
openssl req -new -x509 -days 7200 -subj "/C=cn/OU=dev/O=dev/CN=localhost" -key ${CA_PATH}/ca.key -out ${CA_PATH}/ca.pem

# Server
echo -e "创建 Server 私钥."
openssl genpkey -algorithm RSA -out ${SERVER_PATH}/server.key

echo -e "创建 Server 证书请求文件CSR."
openssl req -new -nodes -key ${SERVER_PATH}/server.key -out ${SERVER_PATH}/server.csr -days 3650 -subj "/C=cn/OU=dev/O=dev/CN=localhost" -config ./openssl.cnf -extensions v3_req

echo -e "创建 Server 证书."
openssl x509 -req -days 3650 -in ${SERVER_PATH}/server.csr -out ${SERVER_PATH}/server.pem -CA ${CA_PATH}/ca.pem -CAkey ${CA_PATH}/ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req

# Client
echo -e "创建 Client 私钥."
openssl genpkey -algorithm RSA -out ${CLIENT_PATH}/client.key

echo -e "创建 Client 证书请求文件CSR."
openssl req -new -nodes -key ${CLIENT_PATH}/client.key -out ${CLIENT_PATH}/client.csr -days 3650 -subj "/C=cn/OU=dev/O=dev/CN=localhost" -config ./openssl.cnf -extensions v3_req

echo -e "创建 Client 证书."
openssl x509 -req -days 3650 -in ${CLIENT_PATH}/client.csr -out ${CLIENT_PATH}/client.pem -CA ${CA_PATH}/ca.pem -CAkey ${CA_PATH}/ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req