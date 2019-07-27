package main

import (
	"os"
	"path"
)

var (
	// TokenFile file with API token for CLI
	TokenFile = path.Join(os.Getenv("HOME"), ".monita_token")
	// APIBaseURL base url for CLI
	APIBaseURL = "http://91.240.84.219"
)
