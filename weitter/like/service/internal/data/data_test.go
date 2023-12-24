package data

import (
	"testing"

	"github.com/zxq97/design/weitter/like/pkg/method"
	"gorm.io/gen"
)

func TestNewDB(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../../pkg/query",
		ModelPkgPath: "../../../pkg/model",
	})
	g.UseDB(NewDB())

	g.ApplyInterface(func(method.LikeMethod) {}, g.GenerateModel("like"))
	g.ApplyInterface(func(method.LikeCountMethod) {}, g.GenerateModel("like_count"))

	g.Execute()
}
