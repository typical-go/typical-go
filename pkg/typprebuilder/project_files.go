package typprebuilder

import "io/ioutil"

func scanProject(root string) (files []string, err error) {
	err = scanDir(root, &files)
	return
}

func scanDir(root string, files *[]string) (err error) {
	fileInfos, err := ioutil.ReadDir(root)
	if err != nil {
		return
	}
	for _, f := range fileInfos {
		if f.IsDir() {
			dirPath := root + "/" + f.Name()
			scanDir(dirPath, files)
		} else {
			*files = append(*files, root+"/"+f.Name())
		}
	}
	return
}
