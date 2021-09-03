package skrib

import (
	"fmt"
	"log"
)

const (
	SUCCESS int = iota
	WARNING
	ERROR
	INFO
)

type Log struct {
	LogType int
	Message string
}

type FGCode string

const (
	GREEN         = "32"
	YELLOW        = "33"
	RED    FGCode = "31"
	CYAN          = "36"
)

type Pattern struct {
	Prefix string
	Color  FGCode
}

func (i *Log) Coloring() Pattern {
	var p Pattern
	switch i.LogType {
	case 0:
		p = Pattern{Color: GREEN, Prefix: "SUCCESS\t"}
	case 1:
		p = Pattern{Color: YELLOW, Prefix: "WARNING\t"}
	case 2:
		p = Pattern{Color: RED, Prefix: "ERROR\t"}
	default:
		p = Pattern{Color: CYAN, Prefix: "INFO\t"}
	}
	return p
}

func (i *Log) Write(fatal bool) {
	p := i.Coloring()
	head := p.formatPrefix()
	log.SetPrefix(head)
	if fatal {
		log.Fatal(i.Message)
	} else {
		log.Println(i.Message)
	}
}

func (p *Pattern) formatPrefix() string {
	return fmt.Sprintf("\033[1;%sm%s\033[0m", p.Color, p.Prefix)
}
