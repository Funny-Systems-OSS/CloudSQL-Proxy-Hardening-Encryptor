# Encryptor for CloudSQL Proxy Hardening
CloudSQL Proxy Hardening:
  https://github.com/Funny-Systems-OSS/cloudsql-proxy-hardening/tree/1.0.0
+ [功能](#功能)
+ [需求](#需求)
+ [安裝](#安裝)
+ [使用](#使用)
## 功能
+ 將 Cloud SQL Proxy 使用的明文金鑰加密，使其只能在指定的 GCE 執行個體上配合 [CloudSQL Proxy Hardening](https://github.com/Funny-Systems-OSS/cloudsql-proxy-hardening/tree/1.0.0) 使用。
## 需求
+ Go 1.15 以上
## 安裝
1. git clone https://github.com/Funny-Systems-OSS/Encryptor-for-CloudSQL-Proxy-Hardening.git
2. cd ./cloudsql-proxy-hardening
3. go build -o ./encrypt_funny ./encrypt_funny.go
## 使用
+ ./encrypt_funny <-f 明文金鑰檔案路徑> <-i 執行個體 ID> [-o 加密金鑰檔案路徑]
  + -f:\
    在 cloud_sql_proxy 中用於認證服務帳號所使用的明文金鑰 json 格式檔案。
  + -i:\
    執行 cloud_sql_proxy_funny 所使用的執行個體 ID。
  + -o:\
    可選參數。設置時將以指定檔名輸出至指定路徑。\
    預設值為 "*原始檔名*.encrypted" 於原始檔案位置。
  + -v:\
    印出版本號並離開程式。
