package dir

import (
	"io/ioutil"
	"strings"
)

func Ls(dir string) map[string]struct{} {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	cmds := make(map[string]struct{})

	for _, file := range files {
		cmd := strings.Replace(file.Name(), ".go", "", 1)
		if cmd == "root" {
			continue
		}

		cmds[cmd] = struct{}{}
	}

	return cmds
}
