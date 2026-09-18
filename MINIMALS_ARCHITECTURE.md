# Minimal Architecture — Go Example

The Go example uses `internal/` as the application-private boundary.

```text
internal/
├── core/
│   ├── config/
│   ├── middlewares/
│   └── security/
├── common/
│   └── utils/
│       └── response/
├── domain/
│   └── models/
├── data/
│   ├── providers/
│   │   └── database/
│   └── services/
├── features/
│   ├── auth/
│   ├── users/
│   ├── posts/
│   └── followers/
└── routes/
```

The existing feature flow is intentionally preserved:

`route -> controller -> function`

A use case or repository should be introduced only when a real responsibility
exists. DTOs are introduced when an external representation differs from the
domain representation.

Authentication token creation/validation is a service responsibility, while
generic response helpers live under common utilities.
