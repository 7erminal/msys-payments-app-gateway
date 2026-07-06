package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnsToClientTables_20260706_124222 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnsToClientTables_20260706_124222{}
	m.Created = "20260706_124222"

	migration.Register("AddColumnsToClientTables_20260706_124222", m)
}

// Run the migrations
func (m *AddColumnsToClientTables_20260706_124222) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE clients ADD COLUMN has_pos INT(11) DEFAULT 0")
	m.SQL("ALTER TABLE clients ADD COLUMN has_app INT(11) DEFAULT 0")
}

// Reverse the migrations
func (m *AddColumnsToClientTables_20260706_124222) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE clients DROP COLUMN has_pos")
	m.SQL("ALTER TABLE clients DROP COLUMN has_app")
}
