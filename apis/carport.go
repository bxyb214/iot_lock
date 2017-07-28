package apis

import (
	"github.com/kataras/iris/context"
	"strings"
	"github.com/bxyb214/iot-lock-server/model"
	"github.com/kataras/iris"
	"fmt"
	"github.com/bxyb214/iot-lock-server/common"
)

// Create 创建分类
func GetAllCarports(ctx context.Context) {
	var carports []model.CarPort
	model.DB.C(common.CARPORT).Find(nil).All(&carports)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(map[string]interface{}{
		"errNo" : model.ErrorCode.SUCCESS,
		"msg"   : "success",
		"data"  : carports,
	})
	return
}

// Create 创建分类
func CreateCarport(ctx context.Context) {
	save(ctx, false)
}

// Update 更新分类
func UpdateCarport(ctx context.Context) {
	save(ctx, true)
}


// Save 保存分类（创建或更新）
func save(ctx context.Context, isEdit bool) {


	carport := new(model.CarPort)
	err := ctx.ReadJSON(carport)

	fmt.Println("carport = " + carport.Name)

	if err != nil {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]interface{}{
			"errNo" : model.ErrorCode.ERROR,
			"msg"   : "参数无效",
			"data"  : "",
		})
		return
	}

	isError := false
	msg     := ""

	carport.Name = strings.TrimSpace(carport.Name)
	if carport.Name == "" {
		isError = true
		msg     = "分类名称不能为空"
	}

	if isError {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]interface{}{
			"errNo" : model.ErrorCode.ERROR,
			"msg"   : msg,
			"data"  : "",
		})
		return
	}

	var saveErr error
	var errMsg = "error."
	if !isEdit {
		//创建
		saveErr = model.DB.C(common.CARPORT).Insert(carport)
	}
	if saveErr != nil {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]interface{}{
			"errNo" : model.ErrorCode.ERROR,
			"msg"   : errMsg,
			"data"  : "",
		})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(map[string]interface{}{
		"errNo" : model.ErrorCode.SUCCESS,
		"msg"   : "success",
		"data"  : "",
	})
	return
}
