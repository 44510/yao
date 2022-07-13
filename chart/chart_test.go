package chart

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/model"
	"github.com/yaoapp/yao/query"
	"github.com/yaoapp/yao/share"
)

func TestLoad(t *testing.T) {
	share.DBConnect(config.Conf.DB)
	model.Load(config.Conf)
	query.Load(config.Conf)

	Load(config.Conf)
	LoadFrom("not a path", "404.")
	check(t)
}

func check(t *testing.T) {
	keys := []string{}
	for key := range Charts {
		keys = append(keys, key)
	}
	assert.Equal(t, 2, len(keys))
}
