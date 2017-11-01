# extractor

a service that allows uploading and iPhone battery page screen shot and store data to analyse battery stats

Currently the following batter stats will be stored

* Device type - iPhone type
* Time of captured info
* Current battery percentage
* Usage time
* Standby Time
* IsFull charge or partial charge

Dependency Management
----

the project no longer used Godep. Use Dep - https://github.com/golang/dep

Selfsigned certificate for local test purpose

 How to generate a self signed certificate on a unix/linux based system
 cd /etc/ssl/
 openssl req -newkey rsa:2048 -new -x509 -days 365 -nodes -out mongodb-cert.crt -keyout mongodb-cert.key
 
 Now concatenate the certificate and private key to a .pem file, as in the following example:
 
 cat mongodb-cert.key mongodb-cert.crt > mongodb.pem
 
 To run test test mode first run ssc.sh (in sudo mode) to generate the mongodb.pem in /etc/ssl folder.
 The same will be used by docker-compose