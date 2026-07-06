# msys_payment_app_gateway - Project Structure

## Current Structure (Observed)

```text
msys_payment_app_gateway/
├── .git/
├── .github/
├── api/
├── conf/
├── functions/
│   ├── api_functions/
│   └── helpers/
├── controllers/
│   └── services/
├── database/
│   └── migrations/
├── transport/
│   └── middlewares/
├── models/
├── routers/
├── structs/
│   ├── requests/
│   └── responses/
├── swagger/
├── tests/
├── utils/
├── go.mod
├── go.sum
├── main.go
└── msys_payment_app_gateway (compiled binary)
```

## Professional Microservice Structure (Recommended)

```text
msys_payment_app_gateway/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── app/                    # startup wiring, dependency injection
│   ├── config/                 # load env and config files
│   ├── transport/
│   │   └── http/
│   │       ├── handlers/       # moved from controllers/
│   │       ├── middleware/     # moved from middlewares/
│   │       └── router/
│   ├── domain/
│   │   ├── entities/           # core business types
│   │   ├── repositories/       # repository interfaces
│   │   └── services/           # business use cases
│   ├── infra/
│   │   ├── db/                 # db client + migrations integration
│   │   ├── repository/         # concrete repository implementations
│   │   └── external/           # outbound client integrations
│   └── dto/
│       ├── request/            # moved from structs/requests
│       └── response/           # moved from structs/responses
├── pkg/                        # optional shared reusable packages
├── api/
│   ├── openapi/
│   └── postman/
├── deployments/
│   ├── docker/
│   ├── k8s/
│   └── scripts/
├── database/
│   ├── migrations/
│   └── seeds/
├── tests/
│   ├── integration/
│   ├── contract/
│   └── fixtures/
├── .github/
│   └── workflows/
├── docs/
│   ├── architecture/
│   └── runbooks/
├── Makefile
├── .env.example
├── go.mod
└── go.sum
```

## Suggested Changes From Current Layout

1. Move `main.go` into `cmd/server/main.go` to separate entrypoint from business code.
2. Move `controllers/`, `middlewares/`, `routers/` under `internal/transport/http/`.
3. Replace `models/` with clearer separation:
   - `internal/domain/entities`
   - `internal/infra/repository` (DB models + persistence logic)
4. Move `structs/requests` and `structs/responses` to `internal/dto/request` and `internal/dto/response`.
5. Keep DB migrations in `database/migrations`, add `database/seeds` for bootstrapping.
6. Add `internal/app` for startup wiring and dependency graph.
7. Add `deployments/` for runtime manifests and ops scripts.
8. Add `docs/architecture` and `docs/runbooks` for maintainability.

## Practical Transition Plan

1. Create `cmd/server/main.go` and move startup logic from root `main.go`.
2. Introduce `internal/transport/http` and move router + handlers gradually.
3. Split model concerns into domain entities and repository implementations.
4. Keep API behavior unchanged while relocating packages.
5. Add Make targets for `build`, `test`, `lint`, and `run`.
6. Remove tracked local binary (`msys_payment_app_gateway`) from git and add to `.gitignore`.
