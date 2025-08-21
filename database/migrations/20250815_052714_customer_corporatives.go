package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CustomerCorporatives_20250815_052714 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CustomerCorporatives_20250815_052714{}
	m.Created = "20250815_052714"

	migration.Register("CustomerCorporatives_20250815_052714", m)
}

// Run the migrations
func (m *CustomerCorporatives_20250815_052714) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE customer_corporatives(`customer_corporative_id` int(11) NOT NULL AUTO_INCREMENT,`customer_number` varchar(255) NOT NULL,`corp_id` int(11) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`customer_corporative_id`), FOREIGN KEY (corp_id) REFERENCES clients(client_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *CustomerCorporatives_20250815_052714) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `customer_corporatives`")
}
