// Read and Write JSON only.
package main

func main() {
    app := iris.New()
    app.OnErrorCode(iris.StatusNotFound, notFound)

    app.Get("/", index)
    app.Get("/ping", status)

    // IMPORTANT:
    runner, configurator := gateway.New(gateway.Options{
        URLPathParameter: "path",
    })
    app.Run(runner, configurator)
}

func notFound(ctx iris.Context){
    code := ctx.GetStatusCode()
    msg := iris.StatusText(code)
    if err := ctx.GetErr(); err!=nil{
        msg = err.Error(),
    }

    ctx.JSON(iris.Map{
        "Message": msg,
        "Code": code,
    })
}

func index(ctx iris.Context) {
    var req map[string]interface{}
    ctx.ReadJSON(req)
    ctx.JSON(req)
}

func status(ctx iris.Context) {
    ctx.JSON(iris.Map{"Message": "OK"})
}
