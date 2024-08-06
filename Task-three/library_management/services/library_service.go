package services
import "errors"
import "github.com/Tamiru-Alemnew/project-phase-backend-tasks/Task-three/library_management/models"

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() LibraryManager {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

func (l *Library) AddBook(book models.Book) {
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exists := l.members[memberID]
	if !exists {
		member = models.Member{ID: memberID}
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member
	l.books[bookID] = book

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	bookIndex := -1
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			bookIndex = i
			break
		}
	}
	if bookIndex == -1 {
		return errors.New("book not borrowed by this member")
	}

	book := member.BorrowedBooks[bookIndex]
	book.Status = "Available"
	member.BorrowedBooks = append(member.BorrowedBooks[:bookIndex], member.BorrowedBooks[bookIndex+1:]...)
	l.members[memberID] = member
	l.books[bookID] = book

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	availableBooks := []models.Book{}
	for _, book := range l.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}
