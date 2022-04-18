package main

import (
	. "github.com/chef-go/base"
	"github.com/chef-go/chef"

	_ "github.com/chef-go/cluster-default"
	_ "github.com/chef-go/cluster-gossip"
	_ "github.com/chef-go/log-default"
	_ "github.com/chef-go/log-file"
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
