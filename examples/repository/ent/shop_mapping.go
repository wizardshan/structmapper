package ent

import "examples/domain"

func (s *Shop) Mapping() *domain.Shop {
	/**************** mapping start ****************/
	dom := new(domain.Shop)
	dom.ID = s.ID
	dom.Name = s.Name
	return dom

	/**************** mapping end  ****************/
}
