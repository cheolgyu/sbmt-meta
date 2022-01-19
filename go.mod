module github.com/cheolgyu/stock-write-module-meta

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/cheolgyu/stock-write-common v0.0.0
	github.com/cheolgyu/stock-write-model v0.0.0
	github.com/jmoiron/sqlx v1.3.4
)

replace (
	github.com/cheolgyu/stock-write-common v0.0.0 => ../stock-write-common
	github.com/cheolgyu/stock-write-model v0.0.0 => ../stock-write-model
)

go 1.16
