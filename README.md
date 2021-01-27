# TWEETER API (wow such original name!)

The name of this repo is absolutely derived from twitter since I am s**t at thinking of a new name for this app.

This is a API service written in GoLang with RESTful API standards.

## Setup Instructions

The service runs in a MySQL database. Make sure mysql is running. Modify [config.toml](./config/config.toml) to match the exact db specifications. Run migrations to create tables.

```{shell}
make db_migrate_up
```

## Run App Server

Build the web server

```{shell}
make build

make run
```
