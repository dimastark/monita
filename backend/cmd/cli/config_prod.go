// +build production

package main

import (
	"os"
	"path"
)

var (
	TokenFile  = path.Join(os.Getenv("HOME"), ".monita_token")
	APIBaseURL = "https://monita-backend.now.sh"
)
