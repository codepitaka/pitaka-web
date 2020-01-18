package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := SetupRouter()
	if err := router.Run(); err != nil {
    	// exit with a message and a code status 1 on errors
	// todo : things should be done for non stopping web. Need routing page for each error code.
    	log.Fatalf("error running server: %vn", err)
    }
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	// load all templates at once
	r.LoadHTMLGlob("src/static/templates/*")

	// routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", gin.H{
			"title": "Welcome to Pitaka!",
			"contents": []string {
				"메인홈 화면에는 추천 글들이 담길 예정입니다.",
				"추천 code snippet들이 담길 예정입니다.",
				"깃헙, gist와 연동되어도 좋을 것 같구요.",
			},
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "Please Login!",
			"contents": []string {
				"로그인 페이지구요.",
				"계정관련작업은 여기서 합니다.",
				"비밀번호 찾기라든가, 계정 생성이라든가 등등.",
			},
		})
	})

	r.GET("/view", func(c *gin.Context) {
		c.HTML(http.StatusOK, "view", gin.H{
			"title": "You can view!",
			"contents": []string {
				"글을 보는 페이지구요.",
				"WYSIWYG 스타일로 만들 예정입니다.",
				"독자가 특정 버튼을 누르면, 글과 글 사이에서 코딩 창이 딱 튀어나오면 좋겠어요.",
				"깃헙, gist와 연동되어도 좋을 것 같구요.",
			},
		})
	})

	r.GET("/edit", func(c *gin.Context) {
		c.HTML(http.StatusOK, "edit", gin.H{
			"title": "You can edit!",
			"contents": []string {
				"글을 에디팅하는 페이지구요.",
				"WYSIWYG 스타일로 만들 예정입니다.",
				"코드 스니펫 넣는 자리도 필요할 것이구요.",
				"기본적으로 수행해야하는 작업들을 적어둔 파일도 넣을 수 있도록 하면 좋을 듯 싶구요.",
			},
		})
	})

	return r
}
