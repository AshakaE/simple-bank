package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (r CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return r.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "001", Name: "Ashaka", City: "Lagos", Zipcode: "102102", DateofBirth: "2005-05-21", Status: "1"},
		{Id: "002", Name: "Ahastha", City: "Mumbai", Zipcode: "104105", DateofBirth: "2075-11-23", Status: "2"},
	}

	return CustomerRepositoryStub{customers}
}