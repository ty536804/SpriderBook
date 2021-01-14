package Console

import (
	"Book/Services"
	"bufio"
	"strings"
)

func ActTable(reader *bufio.Reader) {
	stdinCon, _ := reader.ReadString('\n')
	stdinCon = strings.TrimSpace(stdinCon)
	stdinSlice := len(stdinCon)

	if stdinSlice < 1 {
		return
	}

	if strings.ToLower(stdinCon) == "init" {
		Services.InitTable()
		Services.InitBookListTable()
	}

	if strings.ToLower(stdinCon) == "drop" {
		Services.DropTable()
		Services.DropBookListTable()
	}

	if strings.ToLower(stdinCon) == "migrate" {
		Services.MigrateTable()
		Services.MigrateBookListTable()
	}
}
