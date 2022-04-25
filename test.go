package main

import (
	"time"

	"github.com/chefsgo/cache"
	"github.com/chefsgo/chef"
	"github.com/chefsgo/mutex"
	"github.com/chefsgo/session"
	"github.com/chefsgo/token"

	. "github.com/chefsgo/base"
	_ "github.com/chefsgo/builtin"

	"github.com/chefsgo/log"
	_ "github.com/chefsgo/log-default"
	_ "github.com/chefsgo/log-file"

	_ "github.com/chefsgo/session-buntdb"
	_ "github.com/chefsgo/session-default"
	_ "github.com/chefsgo/session-redis"

	_ "github.com/chefsgo/mutex-default"
	_ "github.com/chefsgo/mutex-redis"

	_ "github.com/chefsgo/token-default"
	_ "github.com/chefsgo/token-jwt"

	_ "github.com/chefsgo/cache-buntdb"
	_ "github.com/chefsgo/cache-default"
	_ "github.com/chefsgo/cache-file"
	_ "github.com/chefsgo/cache-memory"
	_ "github.com/chefsgo/cache-redis"

	"github.com/chefsgo/data"
	_ "github.com/chefsgo/data-postgres"

	"github.com/chefsgo/queue"
	_ "github.com/chefsgo/queue-default"
	_ "github.com/chefsgo/queue-redis"
)

var (
// LoginOK = chef.Result(0, "login.ok", "登录成功")
)

func init() {
	chef.Configure(Map{
		"queue": Map{
			"driver": "redis",
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

	chef.Register("test.Method", chef.Method{
		Name: "name", Text: "name",
		Action: func(ctx *chef.Context) {
			log.Info("test.Method 被调用啦！")
		},
	})

	chef.Register(chef.StartTrigger, chef.Method{
		Name: "Start", Text: "Start",
		Action: func(ctx *chef.Context) {
			log.Info("系统启动了")
		},
	})
	chef.Register(chef.StopTrigger, chef.Method{
		Name: "Start", Text: "Start",
		Action: func(ctx *chef.Context) {
			log.Info("退出前的狂欢")
		},
	})

	// chef.Register("queue.Filter", queue.Filter{
	// 	Name: "filter", Text: "filter",
	// 	Request: func(ctx *queue.Context) {
	// 		ctx.Next()
	// 	},
	// })

	chef.Register("test.SendMsg", queue.Queue{
		Retry: 5, Name: "name", Text: "name",
		Action: func(ctx *queue.Context) {
			log.Info("test.SendMsg 被调用啦！", ctx.Retries())
			ctx.Invoke("test.Method")
			// log.Info("what", ctx.Retries())
			ctx.Retry()
			// log.Info("retry what", ctx.Retries())
		},
	})

}

func main() {
	chef.Go()
}

func test() {
	chef.Ready()
	log.Debug("cool.")

	ddd, res := chef.Execute("test.Method")
	log.Info("execute", res, ddd)

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

	ttt, eee := token.Sign(&token.Token{ActId: "m8tVfvzhz"})
	log.Debug("token", eee, ttt)
	chef.Go()
}
