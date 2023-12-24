package data

import (
	"testing"

	"github.com/zxq97/design/weitter/account/pkg/method"
	"gorm.io/gen"
)

func TestNewDB(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../../pkg/query",
		ModelPkgPath: "../../../pkg/model",
	})
	g.UseDB(NewDB())

	g.ApplyInterface(func(method.UsersMethod) {}, g.GenerateModel("users"))

	g.Execute()
}
