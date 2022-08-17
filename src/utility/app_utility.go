package utility

import (
	"crypto/sha512"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type stringUtility struct {
	ToSha512 func(data string) string
}

var AttachmentPath = "./attachment"
var ConfigPath = "./config.toml"

func NewUuid() string {
	uuid := uuid.New()
	return strings.Replace(uuid.String(), "-", "", -1)
}

var String = stringUtility{
	ToSha512: func(data string) string {
		sum := sha512.Sum512([]byte(data))
		return fmt.Sprintf("%x", sum)
	},
}
