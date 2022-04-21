package main

import (
	. "github.com/chefsgo/base"
	"github.com/chefsgo/chef"

	_ "github.com/chefsgo/cluster-default"
	_ "github.com/chefsgo/cluster-gossip"
	_ "github.com/chefsgo/log-default"
	_ "github.com/chefsgo/log-file"
	_ "github.com/chefsgo/mutex-default"
	_ "github.com/chefsgo/mutex-redis"
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
	})
}

func main() {
	chef.Go()
}
