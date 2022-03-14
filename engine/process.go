package engine

import (
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
	"github.com/yaoapp/gou"
	"github.com/yaoapp/kun/exception"
	"github.com/yaoapp/yao/config"
	"github.com/yaoapp/yao/share"
	"github.com/yaoapp/yao/xfs"
)

func init() {
	// 注册处理器
	gou.RegisterProcessHandler("xiang.main.Ping", processPing)
	gou.AliasProcess("xiang.main.Ping", "xiang.sys.Ping")

	gou.RegisterProcessHandler("xiang.main.FileContent", processFileContent)
	gou.RegisterProcessHandler("xiang.main.AppFileContent", processAppFileContent)

	gou.RegisterProcessHandler("xiang.main.Inspect", processInspect)
	gou.AliasProcess("xiang.main.Inspect", "xiang.sys.Inspect")

	gou.RegisterProcessHandler("xiang.main.Favicon", processFavicon)
}

// processCreate 运行模型 MustCreate
func processPing(process *gou.Process) interface{} {
	res := map[string]interface{}{
		"engine":  share.BUILDNAME,
		"version": share.VERSION,
	}
	return res
}

// processInspect 返回系统信息
func processInspect(process *gou.Process) interface{} {
	// share.App.Icons.Set("favicon", "/api/xiang/favicon.ico")
	// return share.App.Public()

	// Return app.json for xgen-next debug
	info := map[string]interface{}{}
	fs := xfs.New(config.Conf.Root)
	if fs.MustExists("/app.json") {
		err := jsoniter.Unmarshal(fs.MustReadFile("/app.json"), &info)
		if err != nil {
			exception.New("解析应用失败 %s", 500, err).Throw()
		}
	}
	return info
}

// processFavicon 运行模型 MustCreate
func processFavicon(process *gou.Process) interface{} {
	return xfs.DecodeString(share.App.Icons.Get("png").(string))
}

// processFileContent 返回文件内容
func processFileContent(process *gou.Process) interface{} {
	process.ValidateArgNums(2)
	filename := process.ArgsString(0)
	encode := process.ArgsBool(1, true)
	content := xfs.Stor.MustReadFile(filename)
	if encode {
		return xfs.Encode(content)
	}
	return string(content)
}

// processAppFileContent 返回应用文件内容
func processAppFileContent(process *gou.Process) interface{} {
	process.ValidateArgNums(2)
	fs := xfs.New(filepath.Join(config.Conf.Root, "data"))
	filename := process.ArgsString(0)
	encode := process.ArgsBool(1, true)
	content := fs.MustReadFile(filename)
	if encode {
		return xfs.Encode(content)
	}
	return string(content)
}
