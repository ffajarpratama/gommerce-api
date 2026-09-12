# gommerce-api

A REST API for a simple e-commerce backend, written in Go. Console (admin) and customer auth, product catalog, and media uploads today; cart/checkout coming.

## Tech stack

- Go 1.22
- [chi](https://github.com/go-chi/chi) — HTTP router
- [GORM](https://gorm.io/) + MySQL — persistence
- [Cloudinary](https://cloudinary.com/) — media storage
- [golang-jwt](https://github.com/golang-jwt/jwt) — auth
- [go-playground/validator](https://github.com/go-playground/validator) — request validation
- [viper](https://github.com/spf13/viper) — config

## Project structure

```plaintext
cmd/            entrypoint + app wiring
config/         env config loading
constant/       shared + per-domain constants
internal/http/  handlers (console, customer, media), middleware, request/response DTOs
internal/model/ gorm models
internal/repository/ data access
internal/usecase/    business logic
lib/            integrations: mysql, cloudinary, jwt, hash, custom_error, custom_validator
util/           generic helpers
```

See `.claude/references/architecture.md` for a more detailed breakdown.

## Getting started

1. Copy the env file and fill in your own values:

   ```bash
   cp .env.example .env
   ```

   You'll need a MySQL database and a [Cloudinary](https://cloudinary.com/) account (free tier works).

2. Install dependencies and run:

   ```bash
   go run ./cmd/main.go
   ```

   Or with live-reload during development ([air](https://github.com/air-verse/air)):

   ```bash
   air
   ```

   Server listens on `APP_PORT` from `.env` (defaults to `7007`).

## API overview

All routes are prefixed with `/api/v1`. 🔒 = requires `Authorization: Bearer <token>`.

### Console (admin)

| Method | Path                            |     |
| ------ | ------------------------------- | --- |
| POST   | `/console/auth/login`           |     |
| GET    | `/console/auth/profile`         | 🔒  |
| POST   | `/console/product`              | 🔒  |
| GET    | `/console/product`              | 🔒  |
| GET    | `/console/product/{product_id}` | 🔒  |
| PUT    | `/console/product/{product_id}` | 🔒  |
| DELETE | `/console/product/{product_id}` | 🔒  |

### Customer

| Method | Path                       |     |
| ------ | -------------------------- | --- |
| POST   | `/customer/auth/register`  |     |
| POST   | `/customer/auth/login`     |     |
| GET    | `/customer/auth/profile`   | 🔒  |

### Media

| Method | Path      |     |
| ------ | --------- | --- |
| POST   | `/media`  | 🔒  |

Multipart upload — `file` + `location` (`avatar` or `product`), 2MB max.

All responses use a consistent envelope: `{ "success": bool, "data": ..., "paging": ..., "error": ... }`.

## Roadmap

This is a work in progress. Planned next:

- [ ] Normalize product categories into their own table (currently a free-text `category_name` field)
- [ ] Payment method table + integration
- [ ] Cart handling (add/remove/update items, view cart)
- [ ] Order / checkout flow (create order from cart, order status, order history)

## License

MIT — see [LICENSE](./LICENSE).
