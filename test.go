package main

import (
	. "github.com/chefsgo/base"
	"github.com/chefsgo/chef"

	_ "github.com/chefsgo/cluster-default"
	_ "github.com/chefsgo/cluster-gossip"
	_ "github.com/chefsgo/log-default"
	_ "github.com/chefsgo/log-file"
)

var (
	LoginOK = chef.Result(0, "login.ok", "登录成功")
)

func init() {
	chef.Config(Map{
		"name":    "chef",
		"version": "1.0.0",
		"cluster": Map{
			"driver": "gossip",
		},
	})
}

func main() {
	chef.Go()
}
