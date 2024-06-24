package guml

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Guml struct {
	entities      []string
	relationships []string
}

func New() *Guml {
	return &Guml{}
}

func (g *Guml) checkDotInstalled() error {
	cmd := exec.Command("dot", "-V")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("dot command not found, please install Guml")
	}
	return nil
}

/*
Generate the output diagram as a PNG file.
Example:

	g.New().Entity(StructA).Entity(StructB).OneToOne(StructA,StructB).Gen("output")
*/
func (g *Guml) Gen(filename string) error {
	if err := g.checkDotInstalled(); err != nil {
		return err
	}

	var buffer bytes.Buffer
	buffer.WriteString("digraph G {\nnode [shape=record];\n")
	for _, entity := range g.entities {
		buffer.WriteString(entity)
	}
	for _, relationship := range g.relationships {
		buffer.WriteString(relationship)
	}
	buffer.WriteString("}\n")

	cmd := exec.Command("dot", "-Tpng", "-o", fmt.Sprintf("%s.png", filename))
	cmd.Stdin = &buffer
	return cmd.Run()
}
