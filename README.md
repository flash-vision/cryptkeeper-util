# cryptkeeper-util
 Compiled tool that works with cryptkeeper

 This app encrypts your config file for use with the cryptkeeper module.
 build:

> go build -o /path/cryptkeeper-util ./main.go

 usage:

 > ./app/cryptkeeper-util -passkey <plaintext_password> -configfile ./path/to/config.yml

 This will create an encrypted version of the config file with .crypt added.
