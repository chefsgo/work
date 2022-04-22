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
	_ "github.com/chefsgo/data-postgres"
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

	// s := Vars{
	// 	"name": Var{
	// 		Required: true, Name: "名称", Desc: "名称",
	// 	},
	// }

	chef.Configure(Map{
		"name":    "chef",
		"role":    "user",
		"version": "1.0.0",
		"data": Map{
			"driver": "postgres",
		},
		"cluster": Map{
			"driver": "gossip",
		},
		"token": Map{
			"driver": "default",
		},
	})

	chef.Register("test.Method", chef.Method{
		Name: "name", Desc: "name",
		Action: func(ctx *chef.Process) {
			ctx.Info("test.Method 被调用啦！")
		},
	})

}

func main() {
	chef.Go()
}
