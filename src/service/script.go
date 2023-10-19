package service

import (
	"errors"
	"github.com/nuwa/server.v3/bean"
	"github.com/nuwa/server.v3/db"
	"strings"
)

type Script struct {
}

func (p *Script) Get(id string) (*bean.Script, error) {
	return db.SQLCli.GetScript(id)
}

func (p *Script) List(name, path string) ([]bean.Script, error) {
	return db.SQLCli.ListScript(name, path)
}

func (p *Script) Save(script bean.Script) error {
	return db.SQLCli.PutScript(&script.Id, script.Name, script.Path, script.Value, script.Description)
}

func (p *Script) Remove(id int) error {
	return db.SQLCli.RemoveScript(id)
}

func (p *Script) GetByPath(path string) (*bean.Script, error) {
	return db.SQLCli.GetScriptByPath(path)
}

func (p *Script) ListDirectory() ([]bean.ScriptDirectory, error) {
	return db.SQLCli.ListDirectory(nil)
}
func (p *Script) CreateDirectory(parent int, name string) error {
	return db.SQLCli.CreateDirectory(parent, name)
}
func (p *Script) RemoveDirectory(id int) error {
	// 检查目录是否存在
	directory, err := db.SQLCli.GetDirectory(id)
	if err != nil {
		return err
	}
	var paths = []string{directory.Name}
	var parentId = directory.ParentId
	// 读取到顶层路径
	for i := 0; i < 20; i++ {
		var parent *bean.ScriptDirectory
		parent, err = db.SQLCli.GetDirectory(parentId)
		if err != nil {
			return err
		}
		if parent == nil {
			break
		}
		paths = append([]string{parent.Name}, paths...)
		parentId = parent.ParentId
	}
	// 检查目录是否为空
	directoryItems, err := db.SQLCli.ListDirectory(&id)
	if err != nil {
		return err
	}
	if len(directoryItems) > 0 {
		return errors.New("请先删除子目录")
	}
	// 检查目录脚本
	var path = strings.Join(paths, "/")
	if strings.HasPrefix(path, "ROOT/") {
		path = path[5:]
	}
	script, err := db.SQLCli.ListScript("", path)
	if err != nil {
		return err
	}
	if len(script) > 0 {
		return errors.New("请先删除脚本")
	}
	return db.SQLCli.RemoveDirectory(id)
}

func (p *Script) RenameDirectory(id int, name string) error {
	// 检查目录是否存在
	directory, err := db.SQLCli.GetDirectory(id)
	if err != nil {
		return err
	}
	var paths = []string{directory.Name}
	var parentId = directory.ParentId
	// 读取到顶层路径
	for i := 0; i < 20; i++ {
		var parent *bean.ScriptDirectory
		parent, err = db.SQLCli.GetDirectory(parentId)
		if err != nil {
			return err
		}
		if parent == nil {
			break
		}
		paths = append([]string{parent.Name}, paths...)
		parentId = parent.ParentId
	}
	// 检查目录是否为空
	directoryItems, err := db.SQLCli.ListDirectory(&id)
	if err != nil {
		return err
	}
	if len(directoryItems) > 0 {
		return errors.New("请先删除子目录")
	}
	// 检查目录脚本
	var path = strings.Join(paths, "/")
	if strings.HasPrefix(path, "ROOT/") {
		path = path[5:]
	}
	script, err := db.SQLCli.ListScript("", path)
	if err != nil {
		return err
	}
	if len(script) > 0 {
		return errors.New("请先删除脚本")
	}
	return db.SQLCli.RenameDirectory(id, name)
}
