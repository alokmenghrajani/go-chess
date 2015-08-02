package logger

import (
	"fmt"
)

type Logger struct {
	enabled bool
	offset  int
}

func New(enabled bool) *Logger {
	l := new(Logger)
	l.enabled = enabled
	l.offset = 0
	return l
}

func (logger *Logger) Push() *Logger {
	l := *logger
	l.offset++
	return &l
}

func (logger *Logger) Log(msg string) {
	if !logger.enabled {
		return
	}
	for i := 0; i < logger.offset; i++ {
		fmt.Printf("  ")
	}
	for _, c := range msg {
		fmt.Printf("%c", c)
		if c == '\n' {
			for i := 0; i < logger.offset; i++ {
				fmt.Printf("  ")
			}
		}
	}
	fmt.Printf("\n")
}
