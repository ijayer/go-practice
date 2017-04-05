package myiris

import (
	"github.com/kataras/iris"
)

func MainIris(port *string) {
	iris.Get("/hi_json", func(c *iris.Context) {
		c.JSON(200, iris.Map{"Username": "hzqx", "Age": 12})
	})

	// router
	iris.Get ("/home",  testGet)
	iris.Post("/login", testPost)
	iris.Put ("/add",   testPut)

	// api (RESTFul)
	iris.API("/users/id", UserAPI{})

	// web socket
	WebSocketTest()

	// deal with json
	iris.API("/api/test/deal-with-json", DealWithJson{})

	// static file server
	iris.Static("/update/resource", "/resource/version/", 2)

	iris.Listen(":" + *port)
}