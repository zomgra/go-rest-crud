package repository

import "l2/mux/entities"

var Books []entities.Book

func InitialTestData() {
	Books = append(Books,
		entities.Book{
			Id:    "1",
			Title: "T1",
			Author: &entities.Author{
				FirstName: "F1", LastName: "L1"},
			Isnb: "fu738273728",
		},
		entities.Book{
			Id:    "2",
			Title: "T2",
			Author: &entities.Author{
				FirstName: "F2", LastName: "L2"},
			Isnb: "jekd3298289748932",
		},
		entities.Book{
			Id:    "3",
			Title: "T3",
			Author: &entities.Author{
				FirstName: "F3", LastName: "L3"},
			Isnb: "lsls28291",
		})
}
