module github.com/chefsgo/work

go 1.17

replace (
	github.com/chefsgo/base => /Users/yatniel/chefsgo/base
	github.com/chefsgo/builtin => /Users/yatniel/chefsgo/builtin

	github.com/chefsgo/cache => /Users/yatniel/chefsgo/cache
	github.com/chefsgo/cache-buntdb => /Users/yatniel/chefsgo/cache-buntdb
	github.com/chefsgo/cache-default => /Users/yatniel/chefsgo/cache-default
	github.com/chefsgo/cache-file => /Users/yatniel/chefsgo/cache-file
	github.com/chefsgo/cache-memory => /Users/yatniel/chefsgo/cache-memory
	github.com/chefsgo/cache-redis => /Users/yatniel/chefsgo/cache-redis
	github.com/chefsgo/chef => /Users/yatniel/chefsgo/chef

	// github.com/chefsgo/cluster-default => /Users/yatniel/chefsgo/cluster-default
	// github.com/chefsgo/cluster-gossip => /Users/yatniel/chefsgo/cluster-gossip

	// github.com/chefsgo/http-default => /Users/yatniel/chefsgo/http-default

	github.com/chefsgo/data => /Users/yatniel/chefsgo/data
	github.com/chefsgo/data-postgres => /Users/yatniel/chefsgo/data-postgres

	github.com/chefsgo/event => /Users/yatniel/chefsgo/event
	github.com/chefsgo/event-default => /Users/yatniel/chefsgo/event-default
	github.com/chefsgo/event-redis => /Users/yatniel/chefsgo/event-redis

	github.com/chefsgo/log => /Users/yatniel/chefsgo/log
	github.com/chefsgo/log-default => /Users/yatniel/chefsgo/log-default
	github.com/chefsgo/log-file => /Users/yatniel/chefsgo/log-file

	github.com/chefsgo/mutex => /Users/yatniel/chefsgo/mutex
	github.com/chefsgo/mutex-default => /Users/yatniel/chefsgo/mutex-default
	github.com/chefsgo/mutex-redis => /Users/yatniel/chefsgo/mutex-redis

	github.com/chefsgo/queue => /Users/yatniel/chefsgo/queue
	github.com/chefsgo/queue-default => /Users/yatniel/chefsgo/queue-default
	github.com/chefsgo/queue-redis => /Users/yatniel/chefsgo/queue-redis

	github.com/chefsgo/session => /Users/yatniel/chefsgo/session
	github.com/chefsgo/session-buntdb => /Users/yatniel/chefsgo/session-buntdb
	github.com/chefsgo/session-default => /Users/yatniel/chefsgo/session-default
	github.com/chefsgo/session-redis => /Users/yatniel/chefsgo/session-redis

	// github.com/chefsgo/session-buntdb => /Users/yatniel/chefsgo/session-buntdb
	// github.com/chefsgo/session-default => /Users/yatniel/chefsgo/session-default
	// github.com/chefsgo/session-file => /Users/yatniel/chefsgo/session-file
	// github.com/chefsgo/session-memory => /Users/yatniel/chefsgo/session-memory

	github.com/chefsgo/token => /Users/yatniel/chefsgo/token
	github.com/chefsgo/token-default => /Users/yatniel/chefsgo/token-default
	github.com/chefsgo/token-jwt => /Users/yatniel/chefsgo/token-jwt
	github.com/chefsgo/util => /Users/yatniel/chefsgo/util

// github.com/chefsgo/view-default => /Users/yatniel/chefsgo/view-default
)

require (
	github.com/chefsgo/base v0.0.0-00010101000000-000000000000
	github.com/chefsgo/builtin v0.0.0-00010101000000-000000000000
	github.com/chefsgo/cache v0.0.0-00010101000000-000000000000
	github.com/chefsgo/cache-buntdb v0.0.0-00010101000000-000000000000
	github.com/chefsgo/cache-default v0.0.0-00010101000000-000000000000
	github.com/chefsgo/cache-file v0.0.0-00010101000000-000000000000
	github.com/chefsgo/cache-memory v0.0.0-00010101000000-000000000000
	github.com/chefsgo/cache-redis v0.0.0-00010101000000-000000000000
	github.com/chefsgo/chef v0.0.0-00010101000000-000000000000
	github.com/chefsgo/data v0.0.0-00010101000000-000000000000
	github.com/chefsgo/data-postgres v0.0.0-00010101000000-000000000000
	github.com/chefsgo/event v0.0.0-00010101000000-000000000000
	github.com/chefsgo/event-default v0.0.0-00010101000000-000000000000
	github.com/chefsgo/event-redis v0.0.0-00010101000000-000000000000
	github.com/chefsgo/log v0.0.0-00010101000000-000000000000
	github.com/chefsgo/log-default v0.0.0-00010101000000-000000000000
	github.com/chefsgo/log-file v0.0.0-00010101000000-000000000000
	github.com/chefsgo/mutex v0.0.0-00010101000000-000000000000
	github.com/chefsgo/mutex-default v0.0.0-00010101000000-000000000000
	github.com/chefsgo/mutex-redis v0.0.0-00010101000000-000000000000
	github.com/chefsgo/queue v0.0.0-00010101000000-000000000000
	github.com/chefsgo/queue-default v0.0.0-00010101000000-000000000000
	github.com/chefsgo/queue-redis v0.0.0-00010101000000-000000000000
	github.com/chefsgo/session v0.0.0-00010101000000-000000000000
	github.com/chefsgo/session-buntdb v0.0.0-00010101000000-000000000000
	github.com/chefsgo/session-default v0.0.0-00010101000000-000000000000
	github.com/chefsgo/session-redis v0.0.0-00010101000000-000000000000
	github.com/chefsgo/token v0.0.0-00010101000000-000000000000
	github.com/chefsgo/token-default v0.0.0-00010101000000-000000000000
	github.com/chefsgo/token-jwt v0.0.0-00010101000000-000000000000
)

require (
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/chefsgo/util v0.0.0-00010101000000-000000000000 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gomodule/redigo v1.8.8 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lib/pq v1.10.5 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/speps/go-hashids/v2 v2.0.1 // indirect
	github.com/tidwall/btree v1.1.0 // indirect
	github.com/tidwall/buntdb v1.2.9 // indirect
	github.com/tidwall/gjson v1.12.1 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
)
