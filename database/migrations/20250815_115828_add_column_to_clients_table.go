package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToClientsTable_20250815_115828 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToClientsTable_20250815_115828{}
	m.Created = "20250815_115828"

	migration.Register("AddColumnToClientsTable_20250815_115828", m)
}

// Run the migrations
func (m *AddColumnToClientsTable_20250815_115828) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE clients ADD COLUMN client_corp_id VARCHAR(100) DEFAULT NULL AFTER client_name")
}

// Reverse the migrations
func (m *AddColumnToClientsTable_20250815_115828) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
