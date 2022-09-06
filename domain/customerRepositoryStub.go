package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Vasa", "New York", "1213", "2000-01-01", "1"},
		{"2", "Nik", "Mexico", "876", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
