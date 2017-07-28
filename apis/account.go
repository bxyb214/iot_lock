package apis

import (
    "github.com/kataras/iris/context"
)

func Login(ctx context.Context){
    ctx.JSON(context.Map{"message": "Hello iris web framework."})
}
