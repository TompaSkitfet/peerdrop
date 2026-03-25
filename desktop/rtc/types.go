package rtc

import "os"

type IncomingFile struct {
	File *os.File
	Name string
}
