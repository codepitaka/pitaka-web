package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pitaka-web/src/config"
)

func SetRouter(engineRouter *gin.Engine, configurations *config.Configuration) {	
	engineRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Welcome to Pitaka!",
			"contents": []string {
				"메인홈 화면에는 추천 글들이 담길 예정입니다.",
				"추천 code snippet들이 담길 예정입니다.",
				"깃헙, gist와 연동되어도 좋을 것 같구요.",
			},
			"configurations": configurations,
		})
	})

	engineRouter.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Please Login!",
			"contents": []string {
				"로그인 페이지구요.",
				"계정관련작업은 여기서 합니다.",
				"비밀번호 찾기라든가, 계정 생성이라든가 등등.",
			},
			"configurations": configurations,
		})
	})
	
	engineRouter.GET("/edit", func(c *gin.Context) {
		c.HTML(http.StatusOK, "edit.html", gin.H{
			"title": "You can edit!",
			"contents": []string {
				"글을 에디팅하는 페이지구요.",
				"WYSIWYG 스타일로 만들 예정입니다.",
				"코드 스니펫 넣는 자리도 필요할 것이구요.",
				"기본적으로 수행해야하는 작업들을 적어둔 파일도 넣을 수 있도록 하면 좋을 듯 싶구요.",
			},
			"configurations": configurations,
		})
	})
	
	engineRouter.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", gin.H{
			"title": "You can create!",
			"contents": []string {
				"글을 생성하는 페이지구요.",
				"이 버튼을 눌렀으니, 처음에는 라우트가 /create일 거에요.",
				"하지만, 에디터에 뭐라도 한 글자 쓰면, /posts/:id/edit으로 갈 것이에요.",
				"왜냐하면, 에디터에 뭐라도 한 글자 쓰면, 자동으로 저장이 되고, post의 id를 부여받기 때문이에요.",
			},
			"configurations": configurations,
		})
	})
	
	engineRouter.GET("/vecty", func(c *gin.Context) {
		c.HTML(http.StatusOK, "vecty.html", gin.H{
			"title": "You can use vecty!",
			"contents": []string {
				"vecty를 사용할 수 있어요.",
			},
			"configurations": configurations,
		})
	})
	
	// router.GET("/blog/:userid/status", func(c *gin.Context) {
	// userid := c.Param("userid") 
	// message := "userid is " + userid
	// c.String(http.StatusOK, message)
	// fmt.Println(message)
	// })
}