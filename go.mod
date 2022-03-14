module github.com/cheolgyu/tb

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/cheolgyu/base v0.0.0
	github.com/cheolgyu/model v0.0.0
	github.com/jmoiron/sqlx v1.3.4
)

replace (
	github.com/cheolgyu/base v0.0.0 => ../base
	github.com/cheolgyu/model v0.0.0 => ../model
)

go 1.16
