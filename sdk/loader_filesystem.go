package sdk

import (
	"context"
	"crypto/md5"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/sandwich-go/boost/xpanic"
	"github.com/sandwich-go/logbus"

	"github.com/sandwich-go/xconf-providers/pkg/filenotify"
)

type filesystemLoader struct {
	sync.Mutex
	fileMap map[string]string
}

func newLocalLoader() *filesystemLoader {
	return &filesystemLoader{fileMap: make(map[string]string)}
}

func (p *filesystemLoader) isChange(name string, data []byte) bool {
	p.Lock()
	defer p.Unlock()
	hash := md5.New()
	hash.Write(data)
	md5Str := string(hash.Sum(nil))
	if v, ok := p.fileMap[name]; ok && v == md5Str {
		return false
	}
	p.fileMap[name] = md5Str
	return true
}
func (p *filesystemLoader) Name() string                     { return "file-system" }
func (p *filesystemLoader) Read(path string) ([]byte, error) { return os.ReadFile(path) }

type LoaderFunc = func(ctx context.Context, key string, data []byte) error

func (p *filesystemLoader) MustWatch(
	path string,
	exitChan chan struct{},
	loaderFunc LoaderFunc) {

	watcher := filenotify.NewPollingWatcher()
	if err := watcher.Add(path); err != nil {
		panic("filesystemLoader MustWatch got error: " + err.Error())
	}
	var wg sync.WaitGroup
	wg.Add(1)
	// 防止异常退出
	go xpanic.AutoRecover("filesystemLoader", func() {
		wg.Done()
		defer func() {
			_ = watcher.Close()
		}()
		for {
			select {
			case event := <-watcher.Events():
				if (event.Op & fsnotify.Write) == fsnotify.Write {
					if b, err := p.Read(path); err == nil && len(b) != 0 {
						if p.isChange(path, b) {
							err := loaderFunc(context.Background(), path, b)
							if err != nil {
								logbus.Error(fmt.Sprintf("plugin local watcher error %s", err.Error()))
							} else {
								logbus.Debug("plugin local watcher reload succ", logbus.String("path", path))
							}
						}
					}
				}
			case err := <-watcher.Errors():
				logbus.Error("plugin filesystemLoader watcher error", logbus.ErrorField(err))
			case <-exitChan:
				return
			}
		}
	}, xpanic.WithAutoRecoverOptionDelayTime(time.Second))

	wg.Wait()
}
