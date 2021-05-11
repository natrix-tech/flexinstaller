package main

import (
	"strings"
)

func (p Program) parseArgs() {
	if p.Args != nil {
		for i, v := range p.Args {
			nv := strings.Replace(v, "${BASEPATH}", BASEPATH, -1)
			nv = strings.Replace(nv, "${RESPATH}", BASEPATH+p.Name+"\\", -1)
			p.Args[i] = nv
		}
	}
}
