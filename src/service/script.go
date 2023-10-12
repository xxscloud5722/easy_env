package service

import (
	"github.com/nuwa/server.v3/bean"
	"github.com/nuwa/server.v3/db"
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
