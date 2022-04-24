package main

import (
	"time"

	"github.com/chefsgo/chef"

	. "github.com/chefsgo/base"
	_ "github.com/chefsgo/builtin"

	"github.com/chefsgo/log"
	_ "github.com/chefsgo/log-default"
	_ "github.com/chefsgo/log-file"

	"github.com/chefsgo/session"
	_ "github.com/chefsgo/session-buntdb"
	_ "github.com/chefsgo/session-default"
	_ "github.com/chefsgo/session-redis"

	"github.com/chefsgo/mutex"
	_ "github.com/chefsgo/mutex-default"
	_ "github.com/chefsgo/mutex-redis"

	"github.com/chefsgo/token"
	_ "github.com/chefsgo/token-default"
	_ "github.com/chefsgo/token-jwt"

	"github.com/chefsgo/cache"
	_ "github.com/chefsgo/cache-buntdb"
	_ "github.com/chefsgo/cache-default"
	_ "github.com/chefsgo/cache-file"
	_ "github.com/chefsgo/cache-memory"
	_ "github.com/chefsgo/cache-redis"

	"github.com/chefsgo/data"
	_ "github.com/chefsgo/data-postgres"
)

var (
// LoginOK = chef.Result(0, "login.ok", "登录成功")
)

func init() {

	chef.Configure(Map{
		"data": Map{
			"driver": "postgres",
			"url":    "postgres://chefsgo:chefsgo@127.0.0.1:5432/chefsgo?sslmode=disable",
		},
	})

	data.Register("test", data.Table{
		Name: "test", Text: "test",
		Fields: Vars{
			"id":  Var{Required: false, Type: "int", Name: "id"},
			"msg": Var{Required: true, Type: "string", Name: "msg"},
		},
	})

	// s := Vars{
	// 	"name": Var{
	// 		Required: true, Name: "名称", Desc: "名称",
	// 	},
	// }

	// chef.Configure(Map{
	// 	"name":    "chef",
	// 	"role":    "user",
	// 	"version": "1.0.0",
	// 	"data": Map{
	// 		"driver": "postgres",
	// 	},
	// 	"cluster": Map{
	// 		"driver": "gossip",
	// 	},
	// 	"token": Map{
	// 		"driver": "default",
	// 	},
	// })

	// chef.Register("test.Method", chef.Method{
	// 	Name: "name", Desc: "name",
	// 	Action: func(ctx *chef.Process) {
	// 		ctx.Info("test.Method 被调用啦！")
	// 	},
	// })

}

func main() {
	chef.Ready()
	log.Debug("cool.")

	db := data.Base()

	db.Table("test").Create(Map{"msg": "msg from work111"})

	item := db.Table("test").Entity(1)
	log.Debug("item", item)

	cce1 := cache.Write("key", "asdfaf")
	ccv, cce := cache.Read("key")
	log.Debug("cache", cce1, cce, ccv)

	session.Write("key", Map{"msg": "msg123123123"})
	s, e := session.Read("key")
	log.Debug("session", e, s)

	if mutex.Locked("key", time.Second) {
		log.Debug("key loaded 1")
	} else {
		log.Debug("key unloaded 1")
	}
	if mutex.Locked("key", time.Second) {
		log.Debug("key loaded 2")
	} else {
		log.Debug("key unloaded 2")
	}

	time.Sleep(time.Second)
	if mutex.Locked("key", time.Second) {
		log.Debug("key loaded 3")
	} else {
		log.Debug("key unloaded 3")
	}

	ttt, eee := token.Sign(&token.Token{ActId: "EkTWtdXf1"})
	log.Debug("token", eee, ttt)

	chef.Go()
}
