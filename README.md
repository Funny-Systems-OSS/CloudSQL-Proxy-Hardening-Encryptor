# Cloud SQL Proxy Hardening
+ [Features](#Features)
+ [Requirements](#Requirements)
+ [Installation](#Installation)
+ [Usage](#Usage)
## Fork from
https://github.com/GoogleCloudPlatform/cloudsql-proxy/tree/v1.19.0
## Features
+ Replace plaintext credential file with encrypted one.
## Requirements
+ Go 1.15 or higher.
## Installation
1. git clone https://github.com/Funny-Systems-OSS/cloudsql-proxy-hardening.git
2. cd ./cloudsql-proxy-hardening
3. go build -o ./encrypt_funny ./encrypt_funny.go
4. cd ./cloudsql-proxy-1.19.0
5. go build -o ../cloud_sql_proxy_funny ./cmd/cloud_sql_proxy/
## Usage
### encrypt_funny
+ ./encrypt_funny <-f credential_file_path> <-i instance_ID> [-o encrpted_file_path]
  + -f:\
    The json file be used to retrieve Service Account credential in cloud_sql_proxy.
  + -i:\
    The instance ID which the cloud_sql_proxy will be set.
  + -o:\
    If provided, it is treated as the store path of encrypted file. Default to be the same place as the input with filename '<FILENAME>.encrypted'.
  + -v:\
    Print the version of the app and exit.
### cloud_sql_proxy_funny
+ ./cloud_sql_proxy_funny <-credential_file credential_file_path> [-use_plainfile]
  + -credential_file:\
    The encrypted credential file be used to retrieve Service Account credential in cloud_sql_proxy.
  + -use_plainfile:\
    Setting this flag will allow you to use not encrypted credential file.
