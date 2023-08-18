module github.com/jefferyjob/go-redis-lock

go 1.18

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-redis/redismock/v8 v8.11.5
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)

retract (
	v1.1.0
	v1.0.1
	v1.0.0
)
