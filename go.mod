module github.com/tiyee/holydramon

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/allegro/bigcache v1.2.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/klauspost/compress v1.11.0 // indirect
	github.com/valyala/fasthttp v1.16.0
	go.uber.org/zap v1.16.0
)

replace (
	github.com/tiyee/holydramon/src/api => ./api
	github.com/tiyee/holydramon/src/engine => ./engine
	github.com/tiyee/holydramon/src/hook => ./hook
	github.com/tiyee/holydramon/src/model => ./model
	github.com/tiyee/holydramon/src/service => ./service
)
