package routes

import (
    "testing"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
    "github.com/stretchr/testify/assert"
	"pitaka-web/src/utils"
	"pitaka-web/src/config"
	"log"
)

func setUp() *gin.Engine{	
	configurations := config.New()
	engine := gin.New()
	SetRouter(engine, configurations)
	var templatePaths []string = utils.HTMLTemplatePathsUnder("../static/templates")
	engine.LoadHTMLFiles(templatePaths...)

	return engine
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
   req, _ := http.NewRequest(method, path, nil)
   res := httptest.NewRecorder()
   r.ServeHTTP(res, req)
   return res
}

func Test_Routes(t *testing.T) {
    // Setup router
	router := setUp()
	
	// data for test
	tests := []string {
		"/",
		"/login",
		"/edit",
		"/create",
	}
	
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			// Perform a GET request with that handler.
			res := performRequest(router, "GET", test)
			// Assert we encoded correctly, the request gives a 200
			assert.Equal(t, http.StatusOK, res.Code)
		})
	}
}

func Test_All_Routes_Status_200(t *testing.T) {
    // Setup router
	router := setUp()
	
	// data for test
	tests := []string {
		"/",
		"/login",
		"/edit",
		"/create",
	}
	
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			// Perform a GET request with that handler.
			res := performRequest(router, "GET", test)
			// Assert we encoded correctly, the request gives a 200
			assert.Equal(t, http.StatusOK, res.Code)
		})
	}
}

func Test_HTML_Title(t *testing.T) {
    // Setup router
	router := setUp()
	
	// data for test
	tests := []struct {
		route  string
		title string
	}{
		{"/", "Welcome to Pitaka!"},
		{"/login", "Please Login!"},
		{"/edit", "You can edit!"},
		{"/create", "You can create!"},
	}
	
	for _, test := range tests {
		t.Run(test.route, func(t *testing.T) {
			// Perform a GET request with that handler.
			res := performRequest(router, "GET", test.route)
			// get doc from goQuery
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
		    }
			
			// find h1 element text data and validate
			title := strings.TrimSpace(doc.Find("h1").First().Text())
			// Assert we entered title correctly,
			assert.Equal(t, test.title, title)
		})
	}
}
