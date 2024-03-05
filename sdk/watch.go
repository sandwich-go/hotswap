package sdk

import (
	"context"
	"io/fs"
	"os"
	"path"
	"sort"

	"github.com/sandwich-go/boost/module"
	"github.com/sandwich-go/logbus"
)

const EnvCommitId = "current_revision"
const EnvServiceName = "sys_cd_service"
const EnvCdEnv = "sys_cd_env"
const EnvStage = "sys_stage"
const FlagFileName = "version.txt"

func initWatchDir(spec *PluginSpec) (loadDir string) {
	if !spec.GetHotReload() {
		return spec.GetInternalDir()
	}
	commitId := os.Getenv(EnvCommitId)
	serviceName := os.Getenv(EnvServiceName)
	cdEnv := os.Getenv(EnvCdEnv)
	stage := os.Getenv(EnvStage)

	var watchDir = spec.GetInternalDir()
	if commitId != "" { // in k8s
		// /mount/data/test/gate/online/c2das4/bin/plugin
		watchDir = path.Join(spec.GetMountDir(), cdEnv, serviceName, stage, commitId, spec.GetInternalDir())
	}
	// 检查目录是否存在
	_, err := os.Stat(watchDir)
	if os.IsNotExist(err) {
		// 目录不存在，创建它
		errDir := os.MkdirAll(watchDir, 0755)
		if errDir != nil {
			panic(errDir)
		}
		// /mount/data/test/gate/online
		cleanDir(path.Join(spec.GetMountDir(), cdEnv, serviceName, stage), spec.GetDirsToKeep())
	} else if err != nil {
		panic(err)
	}

	flagFile := path.Join(watchDir, FlagFileName)
	_, err = os.Stat(flagFile)
	if err != nil {
		// 使用内部携带的so加载
		loadDir = spec.GetInternalDir()
	} else { // 使用外部的so加载
		loadDir = watchDir
	}

	loader := NewLocalLoader()
	loader.MustWatch(watchDir, module.ProcessShutdownNotify(),
		func(ctx context.Context, key string, data []byte) error {
			GetManager().ResetPluginDir(watchDir)
			_, err := GetManager().Reload(nil)
			if err != nil {
				logbus.Error("hotswap reload plugin", logbus.String("pluginDir", watchDir), logbus.ErrorField(err))
			} else {
				logbus.Info("hotswap reload success", logbus.String("pluginDir", watchDir))
			}
			return nil
		},
		FlagFileName,
	)
	logbus.Info("hotswap start watching", logbus.String("watchDir", watchDir))
	return
}

func cleanDir(dir string, dirsToKeep int) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err1 := entry.Info()
		if err1 != nil {
			panic(err1)
		}
		infos = append(infos, info)
	}
	// Sort entries by modification time
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].ModTime().After(infos[j].ModTime())
	})

	// Keep the latest n directories
	for i, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if i >= dirsToKeep {
			err1 := os.RemoveAll(dir + "/" + entry.Name())
			if err1 != nil {
				logbus.Error("remove directory failed", logbus.String("dir", entry.Name()), logbus.ErrorField(err1))
			} else {
				logbus.Info("remove directory", logbus.String("dir", entry.Name()))
			}
		}
	}
}
