# 🛠 Terraform HTTP Backend

A lightweight Go-based HTTP server for managing and persisting Terraform state files.  
Designed for internal use as a backend for remote Terraform state storage.

Built with:

- [Go](https://golang.org/)
- [Chi Router](https://github.com/go-chi/chi)
- Custom logging and file-based storage layers

## 🚀 Features

- 🌐 RESTful API to `GET` and `POST` Terraform state
- 🗂 Stores state files per project and environment
- 🔒 Modular architecture with clean logging and local storage
- 🐳 Dockerized for container-based deployments
- 📊 Integrates with **Grafana Loki** for centralized log aggregation

## 📦 API Endpoints

### `GET /state/{projectName}/{environment}`

Fetches the Terraform state for the specified project and environment.

### `POST /state/{projectName}/{environment}`

Stores or updates the Terraform state for the specified project and environment.

## Environmental Variables

This application supports the following environment variables for configuration.

| Variable               | Description                                            | Default    |
| ---------------------- | ------------------------------------------------------ | ---------- |
| `LOG_LEVEL`            | Log verbosity level (`debug`, `info`, `warn`, `error`) | `error`    |
| `STORE_TYPE`           | Backend store type (e.g., `file`)                      | `file`     |
| `FILE_STORE_PATH`      | Directory where state files are stored                 | `./states` |
| `FILE_STORE_EXTENSION` | File extension for saved state files                   | `.json`    |
| `LOKI_URL`             | URL to the Grafana Loki endpoint                       | _(unset)_  |
| `LOKI_ENV`             | Environment name sent to Loki (e.g., `dev`)            | `dev`      |
| `LOKI_APP_VERSION`     | Application version reported in Loki logs              | `0.0.0`    |
