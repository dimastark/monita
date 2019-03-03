package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

// Credentials prompts for user login and password
func Credentials(c *cli.Context) (login string, password string) {
	login = Login(c)
	password = Password(c)

	return
}

// Login prompts for user login
func Login(c *cli.Context) string {
	var err error

	reader := bufio.NewReader(os.Stdin)

	login := c.String("login")

	for err != nil || strings.TrimSpace(login) == "" {
		fmt.Print("Enter login: ")

		login, err = reader.ReadString('\n')
	}

	return strings.TrimSpace(login)
}

// Password prompts for user password
func Password(c *cli.Context) string {
	var err error
	var bytePassword []byte

	password := c.String("password")

	for err != nil || strings.TrimSpace(password) == "" {
		fmt.Print("Enter password:")

		bytePassword, err = terminal.ReadPassword(int(syscall.Stdin))

		password = string(bytePassword)
	}

	return strings.TrimSpace(password)
}
