package controllers

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "github.com/Tamiru-Alemnew/project-phase-backend-tasks/Task-three/library_management/models"
    "github.com/Tamiru-Alemnew/project-phase-backend-tasks/Task-three/library_management/services"
)

// LibraryController handles library operations through the LibraryService interface
type LibraryController struct {
    library *services.Library
}

// NewLibraryController creates a new LibraryController
func NewLibraryController(library *services.Library) *LibraryController {
    return &LibraryController{library: library}
}

func (lc *LibraryController) RunLibrarySystem() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("\nLibrary Management System")
        fmt.Println("1. Add a new book")
        fmt.Println("2. Remove an existing book")
        fmt.Println("3. Borrow a book")
        fmt.Println("4. Return a book")
        fmt.Println("5. List all available books")
        fmt.Println("6. List all borrowed books by a member")
        fmt.Println("7. Exit")
        fmt.Print("Enter your choice: ")

        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            lc.addBook(scanner)
        case "2":
            lc.removeBook(scanner)
        case "3":
            lc.borrowBook(scanner)
        case "4":
            lc.returnBook(scanner)
        case "5":
            lc.listAvailableBooks()
        case "6":
            lc.listBorrowedBooks(scanner)
        case "7":
            return
        default:
            fmt.Println("Invalid choice, please try again.")
        }
    }
}

func (lc *LibraryController) addBook(scanner *bufio.Scanner) {
    fmt.Print("Enter book ID: ")
    scanner.Scan()
    id, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter book title: ")
    scanner.Scan()
    title := scanner.Text()

    fmt.Print("Enter book author: ")
    scanner.Scan()
    author := scanner.Text()

    book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
    lc.library.AddBook(book)
    fmt.Println("Book added successfully!")
}

func (lc *LibraryController) removeBook(scanner *bufio.Scanner) {
    fmt.Print("Enter book ID: ")
    scanner.Scan()
    id, _ := strconv.Atoi(scanner.Text())

    lc.library.RemoveBook(id)
    fmt.Println("Book removed successfully!")
}

func (lc *LibraryController) borrowBook(scanner *bufio.Scanner) {
    fmt.Print("Enter book ID to borrow: ")
    scanner.Scan()
    bookID, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    err := lc.library.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book borrowed successfully!")
    }
}

func (lc *LibraryController) returnBook(scanner *bufio.Scanner) {
    fmt.Print("Enter book ID to return: ")
    scanner.Scan()
    bookID, _ := strconv.Atoi(scanner.Text())

    fmt.Print("Enter member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    err := lc.library.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book returned successfully!")
    }
}

func (lc *LibraryController) listAvailableBooks() {
    books := lc.library.ListAvailableBooks()
    if len(books) == 0 {
        fmt.Println("No available books.")
    } else {
        fmt.Println("Available books:")
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
    }
}

func (lc *LibraryController) listBorrowedBooks(scanner *bufio.Scanner) {
    fmt.Print("Enter member ID: ")
    scanner.Scan()
    memberID, _ := strconv.Atoi(scanner.Text())

    books := lc.library.ListBorrowedBooks(memberID)
    if len(books) == 0 {
        fmt.Println("No borrowed books for this member.")
    } else {
        fmt.Println("Borrowed books:")
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
    }
}
