# https-bi-directional-auth

一个通过自签名证书实现https双向认证的方法，本demo使用go语言实现。  
需要按照下面的方式分别生成server与client的证书，在server和client当中分别使用。  
server与client只是名称上的区别，实际使用过程中没有方向区别。  

## 建立CA

### 生成CA私钥

```bash
$ openssl genrsa -out ca.key 2048
```

### 用CA私钥生成CA证书

```bash
$ openssl req -new -x509 -days 36500 -key ca.key -out ca.crt -subj "/C=CN/ST=hubei/L=wuhan/O=yunify/OU=appcenter/CN=openpitrix.io"
```

### 建立CA相应目录

```bash
$ cd /etc/pki/CA/
$ touch index.txt
```

## 生成server端证书

### 生成server私钥

openssl genrsa -out server.key 2048

### 使用server私钥生成server端证书文件

```bash
$ openssl req -new -key server.key -out server.csr -subj "/C=CN/ST=hubei/L=wuhan/O=yunify/OU=appcenter/CN=openpitrix.io"
```

### 使用server证书请求文件通过CA生成自签名证书
```bash
$ openssl ca -in server.csr -out server.crt -cert ca.crt -keyfile ca.key
```

### 验证server证书

```bash
$ openssl verify -CAfile ca.crt server.crt 
server.crt: OK
```

## 生成client端证书

### 生成client私钥

openssl genrsa -out server.key 2048

### 使用client私钥生成client端证书文件

```bash
$ openssl req -new -key client.key -out client.csr -subj "/C=CN/ST=hubei/L=wuhan/O=yunify/OU=appcenter/CN=openpitrix.io"
```

### 使用client证书请求文件通过CA生成自签名证书

```bash
$ openssl ca -in client.csr -out client.crt -cert ca.crt -keyfile ca.key
```

### 验证client证书

```bash
$ openssl verify -CAfile ca.crt client.crt 
client.crt: OK
```
