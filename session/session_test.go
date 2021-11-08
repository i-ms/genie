package session

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"reflect"
	"testing"
)

func TestSession_InitSession(t *testing.T) {
	// setting dummy values for testing session
	s := &Session{
		CookieLifetime: "100",
		CookieName:     "genie",
		CookiePersist:  "true",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	// initializing session manager
	var sm *scs.SessionManager
	ses := s.InitSession()

	var sessKind reflect.Kind
	var sessType reflect.Type

	// rv is return value that we get from session manager
	rv := reflect.ValueOf(ses)

	// looping through all variables in rv
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("For Loop: ", rv.Kind(), rv.Type(), rv)
		sessKind = rv.Kind()
		sessType = rv.Type()

		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Error("Invalid type or kind; kind: ", rv.Kind(), " type: ", rv.Type())
	}

	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Error("Wrong kind returned testing cookie session. Expected ", reflect.ValueOf(sm).Kind(), " and got ", sessKind)
	}

	if sessType != reflect.ValueOf(sm).Type() {
		t.Error("Wrong type returned testing cookie session. Expected ", reflect.ValueOf(sm).Type(), " and got ", sessType)
	}

}
