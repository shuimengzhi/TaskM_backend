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
	tableName := "tm_project_user"
	model.Generator(tableName)
}
