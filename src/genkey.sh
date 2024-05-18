openssl ecparam -name prime256v1 -genkey -noout -out ec-prime256v1-priv-key.pem
openssl ec -in ec-prime256v1-priv-key.pem -pubout -out ec-prime256v1-pub-key.pem