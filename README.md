# TWEETER API (wow such original name!)

The name of this repo is absolutely derived from twitter since I am s**t at thinking of a new name for this app.

This is a API service written in GoLang with RESTful API standards.

## Setup Instructions

The service runs in a MySQL database. Make sure mysql is running. Modify [config.toml](./config/config.toml) to match the exact db specifications. Run migrations to create tables.

```{shell}
make db_migrate_up
```

You will also need to create a RSA Key Pair to create JWT sessions. To do so, cd into the secrets directory and execute the commands below.

```{shell}
# Generate private key
openssl genrsa -out tweeter-private.pem 2048

# Generate public key
openssl rsa -in tweeter-private.pem -pubout -out tweeter-public.pem
```

## Run App Server

Build the web server

```{shell}
make build

make run
```
