package chapter7

import (
	"fmt"
	"log"
	"net/mail"
)

func ExampleParseAddressList() {
	const list = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
	emails, err := mail.ParseAddressList(list)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range emails {
		fmt.Println(v.Name, v.Address)
	}

	// Output:
	// Alice alice@example.com
	// Bob bob@example.com
	// Eve eve@example.com
}
