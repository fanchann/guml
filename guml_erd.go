package guml

import (
	"bytes"
	"fmt"
	"reflect"
)

func (g *Guml) entitiesVal() []string {
	return g.entities
}

func (g *Guml) relationshipVal() []string {
	return g.relationships
}

/*
Declare an entity for diagram generation.
Example:

	type StructA struct{}
	g := New().Entity(StructA)
*/
func (g *Guml) Entity(entity any) *Guml {
	entityType := reflect.TypeOf(entity)
	entityName := entityType.Name()

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s [label=\"%s|", entityName, entityName))

	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		gumlTag := field.Tag.Get("guml")

		if gumlTag != "" {
			buffer.WriteString(fmt.Sprintf("%s\\l", gumlTag))
		}
	}

	buffer.WriteString("\"];\n")
	g.entities = append(g.entities, buffer.String())
	return g
}

/*
Define a one-to-one relationship between two entities.
Example:

	g.OneToOne(StructA, StructB)
*/
func (g *Guml) OneToOne(entity1, entity2 any) *Guml {
	g.addRelationship(entity1, entity2, "1:1")
	return g
}

/*
Define a one-to-many relationship between two entities.
Example:

	g.OneToMany(StructA, StructB)
*/
func (g *Guml) OneToMany(entity1, entity2 any) *Guml {
	g.addRelationship(entity1, entity2, "1:*")
	return g
}

/*
Define a many-to-many relationship between two entities.
Example:

	g.ManyToMany(StructA, StructB)
*/
func (g *Guml) ManyToOne(entity1, entity2 any) *Guml {
	g.addRelationship(entity1, entity2, "*:1")
	return g
}

/*
Define a many-to-one relationship between two entities.
Example:

	g.ManyToOne(StructA, StructB)
*/
func (g *Guml) ManyToMany(entity1, entity2 any) *Guml {
	g.addRelationship(entity1, entity2, "*:*")
	return g
}

func (g *Guml) addRelationship(entity1, entity2 any, relation string) {
	entity1Name := reflect.TypeOf(entity1).Name()
	entity2Name := reflect.TypeOf(entity2).Name()
	g.relationships = append(g.relationships, fmt.Sprintf("%s -> %s [label=\"%s\"];\n", entity1Name, entity2Name, relation))
}
