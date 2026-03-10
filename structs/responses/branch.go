package responses

import "time"

type Currencies struct {
	CurrencyId   int64     `orm:"auto;omitempty"`
	Symbol       string    `orm:"size(20)"`
	Currency     string    `orm:"size(50)"`
	Active       int       `orm:"omitempty"`
	DateCreated  time.Time `orm:"type(datetime);omitempty"`
	DateModified time.Time `orm:"type(datetime);omitempty"`
	CreatedBy    int       `orm:"omitempty"`
	ModifiedBy   int       `orm:"omitempty"`
}

type Countries struct {
	CountryId       int64       `orm:"auto"`
	Country         string      `orm:"size(255)"`
	Description     string      `orm:"size(500)"`
	CountryCode     string      `orm:"size(20)"`
	DefaultCurrency *Currencies `orm:"rel(fk);column(default_currency)"`
	DateCreated     time.Time   `orm:"type(datetime)"`
	DateModified    time.Time   `orm:"type(datetime)"`
	CreatedBy       int
	ModifiedBy      int
}

type Branches struct {
	BranchId      int64      `orm:"auto"`
	Branch        string     `orm:"size(80);unique"`
	Country       *Countries `orm:"rel(fk);column(country_id)"`
	Location      string     `orm:"column(location)"`
	PhoneNumber   string     `orm:"column(phone_number)"`
	Active        int        `orm:"omitempty"`
	DateCreated   time.Time  `orm:"type(datetime);omitempty"`
	DateModified  time.Time  `orm:"type(datetime);omitempty"`
	CreatedBy     int        `orm:"omitempty"`
	ModifiedBy    int        `orm:"omitempty"`
	BranchManager *Users     `orm:"rel(fk);column(branch_manager);null"`
}

type BranchesResponseDTO struct {
	StatusCode int
	Branches   *[]interface{}
	StatusDesc string
}

type BranchApiResponseDTO struct {
	StatusCode int
	Branch     *Branches
	StatusDesc string
}
