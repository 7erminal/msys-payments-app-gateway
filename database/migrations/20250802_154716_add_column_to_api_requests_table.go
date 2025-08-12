package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToApiRequestsTable_20250802_154716 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToApiRequestsTable_20250802_154716{}
	m.Created = "20250802_154716"

	migration.Register("AddColumnToApiRequestsTable_20250802_154716", m)
}

// Run the migrations
func (m *AddColumnToApiRequestsTable_20250802_154716) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE `api_requests` ADD COLUMN `phone_number` VARCHAR(100) DEFAULT NULL COMMENT 'Phone number' AFTER `request_type`")
}

// Reverse the migrations
func (m *AddColumnToApiRequestsTable_20250802_154716) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
