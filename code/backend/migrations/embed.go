package migrations

import "embed"

// Files holds SQL migrations embedded from code/backend/migrations.
//go:embed *.sql
var Files embed.FS
