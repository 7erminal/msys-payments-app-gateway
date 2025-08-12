package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ApiRequests_20250720_163922 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ApiRequests_20250720_163922{}
	m.Created = "20250720_163922"

	migration.Register("ApiRequests_20250720_163922", m)
}

// Run the migrations
func (m *ApiRequests_20250720_163922) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE api_requests(`api_request_id` int(11) NOT NULL AUTO_INCREMENT,`request` varchar(255) NOT NULL,`request_type` varchar(100) DEFAULT NULL,`request_response` varchar(255) DEFAULT NULL,`request_date` datetime DEFAULT CURRENT_TIMESTAMP,`response_date` datetime DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`api_request_id`))")
}

// Reverse the migrations
func (m *ApiRequests_20250720_163922) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `api_requests`")
}
