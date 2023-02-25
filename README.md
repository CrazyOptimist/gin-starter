# Gin Starter 

Gin boilerplate organized in a modular way.

## Table Of Contents

## DB Migration

```bash
make db_migrate
```

## Development

Create a dotenv file:

```bash
cp .env.example .env
```

Install [air](https://github.com/cosmtrek/air) for live reloading. Air config file is already inside the repo, so simply run:

```bash
air
```

## Test

```bash
make test
```

## Build

```bash
make build_linux
```

Binaries will be generated inside `PROJECT_ROOT/bin/`

## API Documentation

[gin-swagger](https://github.com/swaggo/gin-swagger) is used for API documentation.

To browse docs, open `BASE_URL/swagger/index.html`.

Generate/update docs:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
make docs_generate
```

## TODO

- [ ] Write More :D Tests
- [ ] Containerize
- [ ] Setup CI/CD
- [ ] Configure logger using zap

## License

MIT
