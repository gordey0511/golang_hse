package main

import (
	"fmt"
	"golang_hse/src/bookstore"
	"golang_hse/src/idgenerator"
	"golang_hse/src/model"
	"golang_hse/src/store"
)

func main() {
	books := []model.Book{
		{Title: "Война и мир"},
		{Title: "Преступление и наказание"},
		{Title: "Анна Каренина"},
		{Title: "Мастер и Маргарита"},
		{Title: "1984"},
		{Title: "Убить пересмешника"},
		{Title: "Гордость и предубеждение"},
		{Title: "Старик и море"},
		{Title: "Над пропастью во ржи"},
		{Title: "Великий Гэтсби"},
		{Title: "Тихий Дон"},
		{Title: "Унесённые ветром"},
		{Title: "Доктор Живаго"},
		{Title: "Собачье сердце"},
		{Title: "Золото пылающих скал"},
	}

	// Создаем библиотеку
	library := bookstore.NewLibrary(store.NewMapStore(), idgenerator.NewFnvGenerator())
	library.AddBook(books[0])
	library.AddBook(books[1])
	library.AddBook(books[2])

	// Добавляли книгу, проверяем что она есть
	book, ok := library.FindBook(books[0].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Добавляли книгу, проверяем что она есть
	book, ok = library.FindBook(books[2].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Книгу не добавляли, проверяем что ее нет
	book, ok = library.FindBook(books[7].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Удалили книгу, проверяем что ее нет
	library.RemoveBook(books[0].Title)
	book, ok = library.FindBook(books[0].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Меняем генератор id
	library.SetGenerator(idgenerator.NewAdlerGenerator())

	// Проверяем, что книги которые были добавлены находятся
	book, ok = library.FindBook(books[1].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}
	book, ok = library.FindBook(books[2].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Проверяем, что книга которой не было добавлена не находится
	book, ok = library.FindBook(books[5].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Меняем хранилище
	library.SetStore(store.NewSliceStore())
	library.AddBook(books[7])
	library.AddBook(books[8])
	library.AddBook(books[9])

	// Книгу не добавляли, проверяем что ее нет
	book, ok = library.FindBook(books[0].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Книгу добавляли, проверяем что она есть
	book, ok = library.FindBook(books[8].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}

	// Удаляем книгу и проверяем, что ее больше нет
	library.RemoveBook(books[8].Title)
	book, ok = library.FindBook(books[8].Title)
	if ok {
		fmt.Println(book)
	} else {
		fmt.Println("no book")
	}
}
