package typbuildtool

import "io/ioutil"

func projectFiles(root string) (dirs, files []string, err error) {
	dirs = []string{root}
	err = scanDir(root, &dirs, &files)
	return
}

func scanDir(root string, dirs, files *[]string) (err error) {
	fileInfos, err := ioutil.ReadDir(root)
	if err != nil {
		return
	}
	for _, f := range fileInfos {
		if f.IsDir() {
			dirPath := root + "/" + f.Name()
			*dirs = append(*dirs, dirPath)
			scanDir(dirPath, dirs, files)
		} else {
			*files = append(*files, root+"/"+f.Name())
		}
	}
	return
}
