package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitRBAC() {
}
func InitRBACWithFile() {
	//e, err := casbin.NewEnforcer("conf/model.conf", "conf/policy.csv")
}
func InitRBACWithAdapter() {
	// 使用MySQL数据库初始化一个Xorm适配器
	a, err := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/casbin")
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}

	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	sub := "alice" // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read"  // 用户对资源执行的操作。

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		// 处理err
	}

	if ok == true {
		// 允许alice读取data1
	} else {
		// 拒绝请求，抛出异常
	}
}
