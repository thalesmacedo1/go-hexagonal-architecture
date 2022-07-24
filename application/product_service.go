package application

type ProductService struct {
	Persistence ProductPersistenceI
}

func (s *ProductService) Get(id string) (ProductI, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
