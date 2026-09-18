# Minimal Architecture — Go Example

This project adapts the existing API to a modular architecture focused on
separation of responsibilities without forcing every feature to contain the
same layers.

## Structure

- `internal/` — private application code, following the Go ecosystem.
- `internal/core/` — application-wide infrastructure and configuration.
- `internal/common/` — generic reusable utilities.
- `internal/domain/` — business concepts such as models, enums and types.
- `internal/data/` — DTOs, providers, repositories and services when those
  responsibilities are actually needed.
- `internal/features/` — functional application modules.
- `internal/routes/` — HTTP route composition.

## Feature flow

A feature can remain simple:

`route -> controller -> function`

When a real application responsibility requires it, it can grow to:

`route -> controller -> use case -> repository -> provider`

Repositories, use cases and DTOs are not mandatory ceremony. They are added
when their responsibility exists.

Authentication-related behavior belongs to an authentication service rather
than a generic helper.
