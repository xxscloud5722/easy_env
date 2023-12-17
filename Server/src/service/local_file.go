package service

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/xxscloud5722/easy_env/server/src/bean"
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

func (localFile *LocalFile) ListLocalFile(localPath string) ([]*bean.DirInfo, []*bean.FileInfo, error) {
	var filePath = path.Join(localFile.Root, localPath)
	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, nil, err
	}
	return lo.Map(lo.Filter(files, func(item os.DirEntry, index int) bool {
			return item.IsDir()
		}), func(item os.DirEntry, index int) *bean.DirInfo {
			return &bean.DirInfo{
				Name: item.Name(),
			}
		}), lo.Filter(lo.Map(lo.Filter(files, func(item os.DirEntry, index int) bool {
			return !item.IsDir()
		}), func(item os.DirEntry, index int) *bean.FileInfo {
			info, err := item.Info()
			if err != nil {
				return nil
			}
			return &bean.FileInfo{
				Name: item.Name(),
				Last: info.ModTime().Format("2006-01-02 15:04:05"),
				Size: formatStorageSize(info.Size()),
			}
		}), func(item *bean.FileInfo, index int) bool {
			return item != nil
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

func formatStorageSize(kilobytes int64) string {
	const (
		KB = 1.0
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
		PB = TB * 1024
	)

	// Select Storage unit
	switch {
	case kilobytes < MB:
		return fmt.Sprintf("%.2f KB", float64(kilobytes)/KB)
	case kilobytes < GB:
		return fmt.Sprintf("%.2f MB", float64(kilobytes)/MB)
	case kilobytes < TB:
		return fmt.Sprintf("%.2f GB", float64(kilobytes)/GB)
	case kilobytes < PB:
		return fmt.Sprintf("%.2f TB", float64(kilobytes)/TB)
	default:
		return fmt.Sprintf("%.2f PB", float64(kilobytes)/PB)
	}
}
