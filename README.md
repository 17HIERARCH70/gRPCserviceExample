# SSO Service

This Single Sign-On (SSO) service is designed to authenticate users across various systems. It's built using Go and can be easily deployed with Docker.

## Structure

- `api-contracts/`: Protobuf files and generated Go code for gRPC.
- `sso/`: Main application code including command line tools and internal libraries.
- `migrations/`: SQL migration files for database schemas.
- `storage/`: Persistent storage for SQLite databases.

## Getting Started

### Prerequisites

- Docker & Docker Compose
- Go (optional for local development)

### Installation

#### Docker

Clone the repository and navigate to the directory:

```bash
git clone <repository-url>
cd <repository-dir>
```

Start the service using Docker Compose:

```bash
docker-compose up --build
```

The service should now be accessible at `http://localhost:44044`.

#### Local Setup

To run locally without Docker:

```bash
cd sso
go run ./cmd/sso --config=config/prod.yaml
```

### API

Refer to the `api-contracts/proto/sso/sso.proto` file for gRPC service definitions.

## Authors

- **Nikita Belyakov** - *Git* - [17HIERARCH70](https://github.com/17HIERARCH70/)

