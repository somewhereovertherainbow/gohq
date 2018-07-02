package gohq

import (
	"testing"
	"fmt"
	"log"
)

func TestHQ(t *testing.T) {
	account, err := New("0C29Gs1MSx06fakNyvsYQps9SqLid0Tn3tvSpC4mYYZ90UoGmB4hHX8p3dx53ZUe")
	if err != nil {
		log.Fatal(err)
	}

	me, err := account.Me()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Lives: " + me.Lives)

	users, err := account.SearchUser("Discoli")
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range users.Data {
		fmt.Println(u.Username)
	}
}
