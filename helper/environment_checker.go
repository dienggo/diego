package helper

import "os"

type resultCheckIsExistAllEnvKeys struct {
	Status       bool
	NotExistKeys []string
	ExistKeys    []string
}

func IsExistAllEnvKeys(strings ...string) resultCheckIsExistAllEnvKeys {
	var notExistKeys []string
	var existKeys []string
	for _, v := range strings {
		if os.Getenv(v) == "" {
			notExistKeys = append(notExistKeys, v)
		} else {
			existKeys = append(existKeys, v)
		}
	}
	return resultCheckIsExistAllEnvKeys{
		Status:       len(notExistKeys) > 0,
		NotExistKeys: notExistKeys,
		ExistKeys:    existKeys,
	}
}
