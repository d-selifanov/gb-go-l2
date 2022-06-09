package main

import "github.com/google/uuid"

func (s *Product) FromMap(m map[string]interface{}) {
	s.Code = m["code"].(uuid.UUID)
	s.Name = m["name"].(string)
	s.Price = m["price"].(float64)
	s.Count = m["count"].(int64)

}
