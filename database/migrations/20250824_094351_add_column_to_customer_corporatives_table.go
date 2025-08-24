package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToCustomerCorporativesTable_20250824_094351 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToCustomerCorporativesTable_20250824_094351{}
	m.Created = "20250824_094351"

	migration.Register("AddColumnToCustomerCorporativesTable_20250824_094351", m)
}

// Run the migrations
func (m *AddColumnToCustomerCorporativesTable_20250824_094351) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE customer_corporatives ADD COLUMN isDefault SMALLINT DEFAULT 0 AFTER corp_id;")
}

// Reverse the migrations
func (m *AddColumnToCustomerCorporativesTable_20250824_094351) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
