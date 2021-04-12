package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/seanhamiltonkim/protogo/proto"
	"google.golang.org/protobuf/proto"
)

func writeOutPerson(person *pb.Person, w io.Writer) error {
	_, err := fmt.Fprintf(w, "Name: %s\n", person.Name)
	if err != nil {
		log.Fatalln("Error writing person", err)
		return err
	}
	if person.Email != "" {
		fmt.Fprintf(w, "Email: %s\n", person.Email)
	}
	for _, phone := range person.Phones {
		fmt.Fprintf(w, "Phone: [%s] %s\n", phone.Type, phone.Number)
	}
	ts := person.LastUpdated.AsTime()
	fmt.Fprintf(w, "Updated: %s\n", ts.String())

	return err
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Printf("[Result] The addressbook file %s contains %d person record(s)\n", fname, len(book.People))

	fmt.Println("Done")
}
