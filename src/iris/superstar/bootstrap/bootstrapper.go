package bootstrap

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gorilla/securecookie"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/websocket"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time

	Sessions *sessions.Sessions
}

// New returns a new Bootstrapper.
func New(appName, appOwner string, cfgs ...Configfurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// SetupViews loads the templates.
func (b *Bootstrapper) SetupViews(viewsDir string) {
	var sharedLayoutPath string
	sharedLayoutPath = "shared/layout.html"
	htmlEngine := iris.HTML(viewsDir, ".html").Layout(sharedLayoutPath)

	fmt.Println("html engine ", reflect.TypeOf(htmlEngine))

	// 每次重新加载模板（线上关闭它，影响性能）
	htmlEngine.Reload(false)
	b.RegisterView(htmlEngine)

}

// SetupSessions initializes the sessions, optionally.
func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + b.AppName,
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
	})
}

// SetupWebsockets prepares the websocket server.
func (b *Bootstrapper) SetupWebsockets(endpoint string, onConnection websocket.ConnectionFunc) {
	ws := websocket.New(websocket.Config{})
	ws.OnConnection(onConnection)

	b.Get(endpoint, ws.Handler())
	b.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})
}

// SetupErrorHandlers prepares the http error handlers
// `(context.StatusCodeNotSuccessful`,  which defaults to < 200 || >= 400 but you can change it).
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

const (
	// StaticAssets is the root directory for public assets like images, css, js.

	// go run in ./web 与 main.go 同目录 ./public/ <==> ./web/public/
	// but go run not in ./web 则./public/ <==> 当前目录下的 public/
	// 对于vscode 来说 build run 的当前目录为项目的打开根目录 例：web/public ==> /Users/fanjinlong/dev/go/golib/src/v5u.win/projectapi/web/public
	// StaticAssets = "golearn/iris/superstar/web/public/"
	StaticAssets = "./public/"

	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	// b.SetupViews("./golearn/iris/superstar/web/views")
	// go run in ./web 目录下的 main.go文件
	b.SetupViews("./views")
	b.SetupSessions(24*time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	b.SetupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
