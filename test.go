package main

import (
	. "github.com/chefsgo/base"
	_ "github.com/chefsgo/builtin"
	"github.com/chefsgo/chef"

	_ "github.com/chefsgo/cache-default"
	_ "github.com/chefsgo/cache-file"
	_ "github.com/chefsgo/cache-memory"
	_ "github.com/chefsgo/cluster-default"
	_ "github.com/chefsgo/cluster-gossip"
	_ "github.com/chefsgo/log-default"
	_ "github.com/chefsgo/log-file"
	_ "github.com/chefsgo/mutex-default"
	_ "github.com/chefsgo/mutex-redis"
	_ "github.com/chefsgo/session-buntdb"
	_ "github.com/chefsgo/session-default"
	_ "github.com/chefsgo/session-file"
	_ "github.com/chefsgo/session-memory"
	_ "github.com/chefsgo/session-redis"
	_ "github.com/chefsgo/token-default"
	_ "github.com/chefsgo/token-jwt"
)

var (
	LoginOK = chef.Result(0, "login.ok", "登录成功")
)

func init() {

	chef.Configure(Map{
		"name":    "chef",
		"role":    "user",
		"version": "1.0.0",
		"cluster": Map{
			"driver": "gossip",
		},
		"token": Map{
			"driver": "default",
		},
	})

}

func main() {
	chef.Go()
}
