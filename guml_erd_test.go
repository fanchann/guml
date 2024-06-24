package guml

import (
	"fmt"
	"os"
	"testing"
)

type Address struct {
	ID       string `guml:"id"`
	City     string `guml:"city"`
	District string `guml:"district"`
	Country  string `guml:"country"`
	UserID   string `guml:"UserID"`
}

type User struct {
	ID        string `guml:"id"`
	Username  string `guml:"username"`
	Password  string `guml:"password"`
	AddressID string `guml:"address_id"`
}

type Employee struct {
	ID           string `guml:"id"`
	Name         string `guml:"name"`
	DepartmentID string `guml:"departement_id"`
}

type Department struct {
	ID          string `guml:"id"`
	Name        string `guml:"name"`
	Description string `guml:"description"`
}

func TestGuml(t *testing.T) {
	// entity
	g := New()
	g.Entity(User{})
	g.Entity(Address{})

	if len(g.entitiesVal()) != 2 {
		t.Errorf("expected 2 entities, got %d", len(g.entitiesVal()))
	}

	// one to one
	g.OneToOne(User{}, Address{})
	if len(g.relationshipVal()) != 1 {
		t.Errorf("expected 1 relationship, got %d", len(g.relationshipVal()))
	}

	// one to many
	g.OneToMany(User{}, Address{})
	if len(g.relationshipVal()) != 2 {
		t.Errorf("expected 2.relationshipLen(), got %d", len(g.relationshipVal()))
	}

	// many to many
	g.ManyToMany(User{}, Address{})
	if len(g.relationshipVal()) != 3 {
		t.Errorf("expected 3.relationshipLen(), got %d", len(g.relationshipVal()))
	}

	// many to one
	g.ManyToOne(Employee{}, Department{})
	if len(g.relationshipVal()) != 4 {
		t.Errorf("expected 4.relationshipLen(), got %d", len(g.relationshipVal()))
	}

	// 
	fmt.Printf("g.entitiesVal(): %v\n", g.entitiesVal())
	fmt.Printf("g.relationshipVal(): %v\n", g.relationshipVal())

	// check dot installed?
	err := g.checkDotInstalled()
	if err != nil {
		t.Errorf("expected dot to be installed, but got error: %s", err)
	}

	// generate
	filename := "test_graph"
	err = g.Gen(filename)
	if err != nil {
		t.Errorf("expected Gen to succeed, but got error: %s", err)
	}

	// check if the generated file exists
	if _, err := os.Stat(fmt.Sprintf("%s.png", filename)); os.IsNotExist(err) {
		t.Errorf("expected %s.png to be generated, but it does not exist", filename)
	} else {
		// remove generated file
		os.Remove(fmt.Sprintf("%s.png", filename))
	}
}

func TestCheckDotInstalled(t *testing.T) {
	g := New()
	err := g.checkDotInstalled()
	if err != nil {
		t.Fatalf("dot command not found : %v", err)
	}
}

func TestAddRelationship(t *testing.T) {
	type Entity1 struct{}
	type Entity2 struct{}

	g := New()
	g.addRelationship(Entity1{}, Entity2{}, "1:1")
	expected := "Entity1 -> Entity2 [label=\"1:1\"];\n"

	if len(g.relationshipVal()) != 1 {
		t.Fatalf("expected 1 relationship, got %d", len(g.relationshipVal()))
	}
	if g.relationshipVal()[0] != expected {
		t.Fatalf("expected relationship definition: %s, got: %s", expected, g.relationshipVal()[0])
	}
}
