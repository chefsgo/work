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

	"github.com/chefsgo/event"
	_ "github.com/chefsgo/event-default"
	_ "github.com/chefsgo/event-redis"

	"github.com/chefsgo/http"
	_ "github.com/chefsgo/http-default"
	_ "github.com/chefsgo/view-default"
)

var (
// LoginOK = chef.Result(0, "login.ok", "登录成功")
)

func init() {
	data.Register("test", data.Table{
		Name: "test", Text: "test",
		Fields: Vars{
			"id":  Var{Required: false, Type: "int", Name: "id"},
			"msg": Var{Required: true, Type: "string", Name: "msg"},
		},
	})

	chef.Register("zhCN", chef.Language{
		Accepts: []string{"zh-CN", "zhCN", "cn"},
		Strings: chef.Strings{
			"user_1": "asdfadsf",
			"user_2": "asdfadsf",
		},
	})

	// s := Vars{
	// 	"name": Var{
	// 		Required: true, Name: "名称", Desc: "名称",
	// 	},
	// }

	chef.Register("mime", chef.Mime{
		"text/plain",
	})

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
		"http": Map{
			"driver": "default",
			"sites": Map{
				"www": Map{
					"name": "主站",
				},
			},
		},
	})

	chef.Register("test.Method", chef.Method{
		Name: "name", Text: "name",
		Action: func(ctx *chef.Context) Map {
			log.Info("test.Method 被调用啦！", ctx.Language())
			return Map{
				"msg": "msg from method.",
			}
		},
	})

	// chef.Register(chef.StartTrigger, chef.Method{
	// 	Name: "Start", Text: "Start",
	// 	Action: func(ctx *chef.Context) {
	// 		log.Info("系统启动了")
	// 		time.Sleep(time.Millisecond * 100)
	// 		event.Publish("test.TestEvent")
	// 	},
	// })
	// chef.Register(chef.StopTrigger, chef.Method{
	// 	Name: "Start", Text: "Start",
	// 	Action: func(ctx *chef.Context) {
	// 		log.Info("退出前的狂欢")
	// 	},
	// })

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
		},
	})

	chef.Register("test.TestEvent", event.Event{
		Name: "TestEvent", Text: "TestEvent",
		Action: func(ctx *event.Context) {
			log.Info("test.TestEvent 收到事件了")
		},
	})

	chef.Register(".index", http.Router{
		Uri: "/", Name: "index", Text: "index",
		Action: func(ctx *http.Context) {
			// ctx.Text(ctx.Url.Route(".test", Map{"{id}": 123, "page": 1}))
			ctx.Output(chef.OK)
		},
	})

	chef.Register(".test", http.Router{
		Uri: "/test/{id}", Name: "test", Text: "test",
		Action: func(ctx *http.Context) {
			ctx.View("test")
		},
	})

	chef.Register("XXX", http.Filter{
		Request: func(ctx *http.Context) {
			ctx.Next()
		},
	})

	chef.Register(".upload", http.Router{
		Uri: "/upload", Name: "upload", Text: "upload",
		Action: func(ctx *http.Context) {
			ctx.Output(nil, Map{
				"value":  ctx.Value,
				"upload": ctx.Upload,
			})
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
