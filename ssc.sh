!#/binsh
cd /etc/ssl/
openssl req -newkey rsa:2048 -new -x509 -days 365 -nodes -out mongodb-cert.crt -keyout mongodb-cert.key
 
#Now concatenate the certificate and private key to a .pem file, as in the following example:
 cat mongodb-cert.key mongodb-cert.crt > mongodb.pem
