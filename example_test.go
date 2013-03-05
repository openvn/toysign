package toysign_test

import (
	"fmt"
	"github.com/openvn/toys/view"
	"github.com/openvn/toysign"
	"labix.org/v2/mgo"
	"net/http"
)

func ExampleHandler() {
	//database session
	dbsess, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer dbsess.Close()

	//template for cms
	tmpl := view.NewView("template")
	tmpl.Resource = "//localhost:8080/statics"
	tmpl.Parse("default")
	if err != nil {
		fmt.Println(err.Error())
	}

	http.Handle("/user/", toysign.Handler("/user/", dbsess, tmpl))
	http.ListenAndServe("localhost:8080", nil)
	fmt.Println("done")
	// Output: done
}
