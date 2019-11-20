package image

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	rsLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = rsLetters[rand.Intn(len(rsLetters))]
	}
	return string(b)
}

func joinPath(path, filename string) string {
	return strings.TrimRight(path, "/") + "/" + filename
}

func initDir(dirname string) error {
	if fInfo, err := os.Stat(dirname); err != nil {
		if os.IsExist(err) {
			//Something wrong
			return err
		}
		// Directory does not exist
		return os.Mkdir(dirname, 0777)
	} else {
		if !fInfo.IsDir() {
			return fmt.Errorf("file already exists")
		}
	}
	return nil
}
