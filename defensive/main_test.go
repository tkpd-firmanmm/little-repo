package defensive

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	ferr "faulty/errors"
	"faulty/module/book"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"go.uber.org/fx"
)

type FindBookHttpHandler struct {
	findBook *book.FindBook
}

func (f *FindBookHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	rawID := query.Get("id") //What if ID is empty or invalid
	id, err := strconv.Atoi(rawID)
	if err != nil || id < 0 {
		w.Write([]byte("Parameter id is wrong"))
		return
	}
	book, err := f.findBook.Find(uint(id)) //What is the possible error inside this function
	if err != nil {
		w.Write([]byte(handleError(err)))
		return
	}
	body, err := json.Marshal(book)
	if err != nil {
		w.Write([]byte(handleError(err)))
		return
	}
	w.Write(body)
}

func NewFindBookHTTPHandler(findBook *book.FindBook) *FindBookHttpHandler {
	return &FindBookHttpHandler{
		findBook: findBook,
	}
}

func handleError(err error) string {
	var clientErr *ferr.ClientError
	if errors.As(err, &clientErr) {
		return clientErr.Error()
	}
	log.Println("This is logged internally,\n===================\n " + err.Error()) //Never return internal error to user
	return "Internal Server Error"
}

func TestDefensive(t *testing.T) {

	var findBookHttpHandler *FindBookHttpHandler
	suppliedData := fx.Supply(
		&sql.DB{}, //DUmmy purpose only, not used
	)

	app := fx.New(
		book.LoadBook(),
		suppliedData,
		book.LoadRedundantBookParam(),
		fx.Provide(NewFindBookHTTPHandler),
		fx.Populate(&findBookHttpHandler),
	)
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		t.Fatal(err)
	}

	writer := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/book", nil)
	urlQuery := &url.Values{}
	urlQuery.Add("id", "5")
	req.URL.RawQuery = urlQuery.Encode()
	findBookHttpHandler.Handle(writer, req)

	if err := app.Stop(ctx); err != nil {
		t.Fatal(err)
	}

	log.Println("This is return \n\n" + string(writer.Body.Bytes()))
}
