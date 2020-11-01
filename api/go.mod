module github.com/yumechi/GoAPIPractice

go 1.15

require (
	github.com/0xAX/notificator v0.0.0-20191016112426-3962a5ea8da1 // indirect
	github.com/codegangsta/envy v0.0.0-20141216192214-4b78388c8ce4 // indirect
	github.com/codegangsta/gin v0.0.0-20171026143024-cafe2ce98974
	github.com/gin-gonic/gin v1.6.3
	github.com/mattn/go-shellwords v1.0.10 // indirect
	gopkg.in/urfave/cli.v1 v1.20.0 // indirect
	local.packages/api/controller v0.0.0-00010101000000-000000000000
)

replace local.packages/api/controller => ./controller
