# CloudSQL Proxy Hardening Encryptor
[中文版](https://github.com/Funny-Systems-OSS/CloudSQL-Proxy-Hardening-Encryptor/blob/master/README.zh.md)\
CloudSQL Proxy Hardening: https://github.com/Funny-Systems-OSS/cloudsql-proxy-hardening
+ [Features](#Features)
+ [Requirements](#Requirements)
+ [Installation](#Installation)
+ [Usage](#Usage)
+ [Todo](#Todo)
## Features
+ Replace plaintext credential file with encrypted one which bound to instance ID.
## Requirements
+ Go 1.15 or higher.
## Installation
1. git clone https://github.com/Funny-Systems-OSS/CloudSQL-Proxy-Hardening-Encryptor.git
2. cd ./cloudsql-proxy-hardening
3. go build -o ./encrypt_funny ./encrypt_funny.go
## Usage
+ ./encrypt_funny <-f credential_file_path> <-i instance_ID> [-o encrypted_file_path]
  + -f:\
    The json file be used to retrieve Service Account credential in cloud_sql_proxy.
  + -i:\
    The instance ID which the cloud_sql_proxy_funny will be set.
  + -o:\
    If provided, it is treated as the store path of encrypted file. Default to be the same place as the input with filename '*FILENAME*.encrypted'.
  + -v:\
    Print the version of the app and exit.
## Todo
+ Encrypt credential file using OpenSSL.
