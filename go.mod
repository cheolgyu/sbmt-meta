module github.com/cheolgyu/sbmt-meta

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/cheolgyu/sbm-base v0.0.0
	github.com/cheolgyu/sbm-struct v0.0.0
	github.com/jmoiron/sqlx v1.3.4
)

replace (
	github.com/cheolgyu/sbm-base v0.0.0 => ../sbm-base
	github.com/cheolgyu/sbm-struct v0.0.0 => ../sbm-struct
)

go 1.16
