# apex-validator (go-playground/validator v10) — Conventions

## Language & Framework

- Go 1.13+ (primary)
- golangci-lint v1.41.1 — Code linting

## Project Structure

- Layout: single-package
- Organization: flat
- `//` — Main validator package
- `non-standard/validators//` — Optional validators not in baked_in.go
- `_examples//` — Usage examples
- `translations//` — Localized error messages
- `testdata//` — Test data files

## Naming Conventions

- Files: snake_case
- Functions/methods: camelCase
- Components: PascalCase
- Tests: *_test.go

## Error Handling

- Pattern: error-codes
- Returns ValidationErrors (slice of FieldError) for field failures, InvalidValidationError for bad input. Validators return bool.
- Example in `errors.go`: ValidationErrors []FieldError with Namespace, Field, Tag, Param, Value methods
- Example in `baked_in.go`: Func(fl FieldLevel) bool — all validators return bool

## Testing

- Framework: Go testing stdlib
- Assertions: go-playground/assert/v2 with custom AssertError helpers
- File pattern: *_test.go
- Fixtures: Table-driven tests with tagged test structs
- Coverage: go tool cover

## Import Patterns

- Style: direct
- reflect for struct introspection
- sync.Pool for validator pooling
- context.Context support

## API Patterns

- Example in `validator_instance.go`: New(), Struct(), StructCtx(), Var(), RegisterValidation(), RegisterAlias()
- Example in `baked_in.go`: bakedInValidators map[string]Func — tag name to validation function

## CI/CD

- Platform: GitHub Actions
- Lint: golangci-lint
- Build: go test -race -cover

