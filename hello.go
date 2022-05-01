package main

import (
	. "github.com/chefsgo/base"
	_ "github.com/chefsgo/builtin"
	"github.com/chefsgo/chef"
	"github.com/chefsgo/http"
)

func init() {
	chef.Register(".index", http.Router{
		Uri: "/",
		Action: func(ctx *http.Context) {
			ctx.Output(chef.OK, Map{
				"msg": "msg",
			})
		},
	})

	chef.Register("email", chef.Regular{
		`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$`,
	})

	chef.Register(chef.Regulars{
		"mobile": chef.Regular{`^1[0-9]{10}$`},
		"date": chef.Regular{
			`^(\d{4})(\d{2})(\d{2})$`,
			`^(\d{4})-(\d{2})-(\d{2})$`,
		},
	})

	chef.Register("mp4", chef.Mime{"video/mp4"})

	chef.Register(chef.Mimes{
		"mp4": chef.Mime{"video/mp4"},
		"apk": chef.Mime{
			"application/vnd.android.package-archive",
			"application/vnd.android",
		},
	})

	chef.Register("send.ok", chef.State(0))
	chef.Register(chef.States{
		"login.ok":   0,
		"login.fail": 1,
	})

	chef.Register("zhCN", chef.Language{
		Accepts: []string{"zh", "zhCN", "zh-cn"},
		Name:    "简体中文", Text: "简体中文",
		Strings: chef.Strings{
			"login.ok":   "登录成功",
			"login.fail": "登录失败",
		},
	})

	chef.Register("en", chef.Language{
		Accepts: []string{"en", "enUs", "en-US"},
		Name:    "English", Text: "English",
		Strings: chef.Strings{
			"login.ok":   "Login successful",
			"login.fail": "Login failed",
		},
	})

	chef.Register("zhCN", chef.Strings{
		"login.ok":   "登录成功",
		"login.fail": "登录失败",
	})
	chef.Register("en", chef.Strings{
		"login.ok":   "Login successful",
		"login.fail": "Login failed",
	})

	// 获取语言字段
	// str1 = 登录成功
	// str2 = Login successful
	// str1 := chef.String("zhCN", "login.ok")
	// str2 := chef.String("en", "login.ok")

}

func main() {
	chef.Go()
}
