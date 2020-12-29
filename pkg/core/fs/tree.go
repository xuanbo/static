package fs

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Tree 树
type Tree struct {
	Title    string  `json:"title"`
	Path     string  `json:"path"`
	Href     string  `json:"href"`
	Children []*Tree `json:"children"`
}

// Walk 遍历文件
func Walk(path string, tree *Tree, level, maxLevel int) error {
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	if len(fs) == 0 {
		return nil
	}

	children := make([]*Tree, 0, len(fs))
	for _, f := range fs {
		name := f.Name()
		if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "__MACOSX") {
			continue
		}
		t := &Tree{
			Title: name,
			Path:  filepath.Join(path, name),
			Href:  filepath.Join("/", path, name),
		}
		children = append(children, t)
		if f.IsDir() && level < maxLevel {
			t.Href = ""
			Walk(filepath.Join(path, name), t, level+1, maxLevel)
		}
	}
	tree.Children = children
	return nil
}
