package model_test

import (
	"os"
	"taskm/core"
	"taskm/model"
	"testing"
)

func TestGenerator(t *testing.T) {
	envPath, _ := os.Getwd()
	envPath = envPath + "/../.env"
	core.LoadCore(envPath)
	tableName := "ot_task_material_library"
	model.Generator(tableName)
}
