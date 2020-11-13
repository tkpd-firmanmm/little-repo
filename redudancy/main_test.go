package redudancy

import (
	"context"
	"database/sql"
	"faulty/module/book"
	"log"
	"testing"

	"go.uber.org/fx"
)

func Runner(findBook *book.FindBook) {
	log.Println(findBook.Find(2))
	log.Println(findBook.Find(3))
	log.Println(findBook.Find(4))
}

func TestRedudancy(t *testing.T) {

	suppliedData := fx.Supply(
		&sql.DB{}, //DUmmy purpose only, not used
	)

	app := fx.New(book.LoadBook(), fx.Invoke(Runner), suppliedData, book.LoadRedundantBookParam())
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		t.Fatal(err)
	}
	if err := app.Stop(ctx); err != nil {
		t.Fatal(err)
	}

}
