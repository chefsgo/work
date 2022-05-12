package main

import (
	. "github.com/chefsgo/base"

	_ "github.com/chefsgo/builtin"
	"github.com/chefsgo/chef"
	"github.com/chefsgo/http"

	"github.com/chefsgo/store"
	_ "github.com/chefsgo/store-default"
)

func init() {

	// chef.Register("email", chef.Regular{
	// 	`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$`,
	// })

	// chef.Register(chef.Regulars{
	// 	"mobile": chef.Regular{`^1[0-9]{10}$`},
	// 	"date": chef.Regular{
	// 		`^(\d{4})(\d{2})(\d{2})$`,
	// 		`^(\d{4})-(\d{2})-(\d{2})$`,
	// 	},
	// })

	// chef.Register("mp4", chef.Mime{"video/mp4"})

	// chef.Register(chef.Mimes{
	// 	"mp4": chef.Mime{"video/mp4"},
	// 	"apk": chef.Mime{
	// 		"application/vnd.android.package-archive",
	// 		"application/vnd.android",
	// 	},
	// })

	// chef.Register("send.ok", chef.State(0))
	// chef.Register(chef.States{
	// 	"login.ok":   0,
	// 	"login.fail": 1,
	// })

	// chef.Register("zhCN", chef.Language{
	// 	Accepts: []string{"zh", "zhCN", "zh-cn"},
	// 	Name:    "简体中文", Text: "简体中文",
	// 	Strings: chef.Strings{
	// 		"login.ok":   "登录成功",
	// 		"login.fail": "登录失败",
	// 	},
	// })

	// chef.Register("en", chef.Language{
	// 	Accepts: []string{"en", "enUs", "en-US"},
	// 	Name:    "English", Text: "English",
	// 	Strings: chef.Strings{
	// 		"login.ok":   "Login successful",
	// 		"login.fail": "Login failed",
	// 	},
	// })

	// chef.Register("zhCN", chef.Strings{
	// 	"login.ok":   "登录成功",
	// 	"login.fail": "登录失败",
	// })
	// chef.Register("en", chef.Strings{
	// 	"login.ok":   "Login successful",
	// 	"login.fail": "Login failed",
	// })

	// 获取语言字段
	// str1 = 登录成功
	// str2 = Login successful
	// str1 := chef.String("zhCN", "login.ok")
	// str2 := chef.String("en", "login.ok")

	// chef.Register(store.Configs{
	// 	Driver: chef.DEFAULT, Setting: Map{
	// 		"asdf": "asdfasdf",
	// 	},
	// })

	chef.Configure(Map{
		"http": Map{
			"port": 80,
		},
	})

	// chef.Register("test", cron.Job{
	// 	Time: "0 * * * * *",
	// 	Action: func(ctx *cron.Context) {
	// 		log.Info("测试计划")
	// 	},
	// })

	chef.Register(chef.Mimes{
		"png": chef.Mime{"image/png"},
	})

	chef.Register(".file", http.Router{
		Uri: "/file/{code}",
		Args: Vars{
			"code": Var{Type: "string", Required: true, Name: "代码"},
		},
		Action: func(ctx *http.Context) {
			code := ctx.Args["code"].(string)
			file, err := store.Download(code)
			if err != nil {
				ctx.Text(err.Error())
			} else {
				ctx.File(file.File())
			}
		},
	})

	chef.Register(".upload", http.Router{
		Uri: "/upload",
		Action: func(ctx *http.Context) {
			codes := []string{}
			for _, val := range ctx.Upload {
				file, _, err := store.Upload(val)
				if err == nil {
					codes = append(codes, file.Code())
				}
			}
			ctx.JSON(codes)
		},
	})

	//

	// chef.Register(".test", http.Router{
	// 	Uri: "/test",
	// 	Action: func(ctx *http.Context) {
	// 		code := "jH0xdH05EXlZEXMojqfZK3dZdi4zKiCpjPGOdXEqei4YE3hwKihwd0QslaIR2H4WEHpvdv=="
	// 		file, err := store.Download(code)

	// 		log.Debug(err, file)

	// 		ctx.File(file)
	// 	},
	// })

	// chef.Register(".download", http.Router{
	// 	Uri: "/download",
	// 	Action: func(ctx *http.Context) {
	// 		code := "jH0xdH05EXlZEXMojqfZK3dZdi4zKiCpjPGOdXEqei4YE3hwKihwd0QslaIR2H4WEHpvdv=="
	// 		file, _ := store.Download(code)
	// 		ctx.File(file)
	// 	},
	// })

}

func main() {
	chef.Go()
}
