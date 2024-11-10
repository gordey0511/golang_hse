package bookstore

import (
	"golang_hse/src/idgenerator"
	"golang_hse/src/model"
)

type BookStore interface {
	Find(id uint32) (model.Book, bool)
	Add(id uint32, b model.Book)
	Remove(id uint32)
	Regenerate(generator idgenerator.BookIdGenerator)
}

type Library struct {
	store     BookStore
	generator idgenerator.BookIdGenerator
}

func NewLibrary(store BookStore, generator idgenerator.BookIdGenerator) *Library {
	return &Library{
		store:     store,
		generator: generator,
	}
}

func (library *Library) AddBook(book model.Book) {
	library.store.Add(library.generator.GeneratorId(book.Title), book)
}

func (library *Library) FindBook(title string) (model.Book, bool) {
	return library.store.Find(library.generator.GeneratorId(title))
}

func (library *Library) RemoveBook(title string) {
	library.store.Remove(library.generator.GeneratorId(title))
}

func (library *Library) SetStore(store BookStore) {
	library.store = store
}

func (library *Library) SetGenerator(generator idgenerator.BookIdGenerator) {
	library.generator = generator
	library.store.Regenerate(generator)
}
