package routes

import (
    "testing"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
    "github.com/stretchr/testify/assert"
	"github.com/codepitaka/pitaka-web/src/utils"
	"log"
)

func setUp() *gin.Engine{
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine = SetRouter(engine)
	
	var allTemplates []string = utils.FilesUnder("../static/templates")
	engine.LoadHTMLFiles(allTemplates...)

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
		"/view",
		"/edit",
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
		"/view",
		"/edit",
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
		{"/view", "You can view!"},
		{"/edit", "You can edit!"},
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
			doc.Find("h1").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the title, remove newlines & tabs
				title := strings.TrimSpace(s.Text())

				// Assert we encoded correctly,
				// the request gives a 200
				assert.Equal(t, test.title, title)
			})
		})
	}
}