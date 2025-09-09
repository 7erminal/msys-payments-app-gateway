package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToCustomerCorporativesTable_20250907_054304 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToCustomerCorporativesTable_20250907_054304{}
	m.Created = "20250907_054304"

	migration.Register("AddColumnToCustomerCorporativesTable_20250907_054304", m)
}

// Run the migrations
func (m *AddColumnToCustomerCorporativesTable_20250907_054304) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE customer_corporatives ADD COLUMN isActive SMALLINT DEFAULT 0 AFTER isDefault;")
}

// Reverse the migrations
func (m *AddColumnToCustomerCorporativesTable_20250907_054304) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
