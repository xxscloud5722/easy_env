package service

import (
	"github.com/samber/lo"
	"os"
	"path"
)

type LocalFile struct {
	Root string
}

func (localFile *LocalFile) Init() error {
	if !localFile.IsExist("") {
		var filePath = path.Join(localFile.Root)
		err := os.MkdirAll(filePath, 644)
		if err != nil {
			return err
		}
	}
	return nil
}

func (localFile *LocalFile) ListLocalFile(localPath string) ([]string, []string, error) {
	files, err := os.ReadDir(localPath)
	if err != nil {
		return nil, nil, err
	}
	return lo.Map(lo.Filter(files, func(item os.DirEntry, index int) bool {
			return item.IsDir()
		}), func(item os.DirEntry, index int) string {
			return item.Name()
		}), lo.Map(lo.Filter(files, func(item os.DirEntry, index int) bool {
			return !item.IsDir()
		}), func(item os.DirEntry, index int) string {
			return item.Name()
		}), nil
}

func (localFile *LocalFile) IsExist(localPath string) bool {
	var filePath = path.Join(localFile.Root, localPath)
	if _, err := os.Stat(filePath); err == nil || os.IsExist(err) {
		return true
	}
	return false
}

func (localFile *LocalFile) IsDir(localPath string) (*bool, error) {
	var filePath = path.Join(localFile.Root, localPath)
	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	var result = fi.IsDir()
	return &result, err
}

func (localFile *LocalFile) GetFilePath(localPath string) string {
	return path.Join(localFile.Root, localPath)
}
