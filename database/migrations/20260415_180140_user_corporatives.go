package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserCorporatives_20260415_180140 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserCorporatives_20260415_180140{}
	m.Created = "20260415_180140"

	migration.Register("UserCorporatives_20260415_180140", m)
}

// Run the migrations
func (m *UserCorporatives_20260415_180140) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user_corporatives(`user_corporative_id` int(11) NOT NULL AUTO_INCREMENT,`user_id` varchar(100) NOT NULL,`corp_id` int(11) DEFAULT NULL,`active` int(11) DEFAULT NULL,`default` int(11) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`user_corporative_id`), FOREIGN KEY (corp_id) REFERENCES clients(client_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *UserCorporatives_20260415_180140) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user_corporatives`")
}
