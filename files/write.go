package files

import "io/ioutil"

func SaveFile(name, data string) {
	ioutil.WriteFile(name, []byte(data), 0644)
}
