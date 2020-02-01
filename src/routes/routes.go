package routes

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/codepitaka/pitaka-web/src/config"
)

func SetRouter(engineRouter *gin.Engine) *gin.Engine {	
	engineRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Welcome to Pitaka!",
			"contents": []string {
				"메인홈 화면에는 추천 글들이 담길 예정입니다.",
				"추천 code snippet들이 담길 예정입니다.",
				"깃헙, gist와 연동되어도 좋을 것 같구요.",
			},
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
		})
	})

	engineRouter.GET("/view", func(c *gin.Context) {
		res, err := http.Get(config.DevServerURL + "/posts")
		if err != nil {
			panic(err.Error())
		}
		if res.StatusCode != http.StatusOK{
			c.Status(http.StatusServiceUnavailable)
			return
		}
		
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}
		
		var d map[string]interface{}
		err = json.Unmarshal(body, &d)
		if err != nil {
			panic(err.Error())
		}
		
		posts, ok := d["data"].([]interface{})
		if !ok {
			log.Fatal("not ok")
		}
		
		c.HTML(http.StatusOK, "view.html", gin.H{
			"title": "You can view!",
			"contents": []string {
				"글을 보는 페이지구요.",
				"WYSIWYG 스타일로 만들 예정입니다.",
				"독자가 특정 버튼을 누르면, 글과 글 사이에서 코딩 창이 딱 튀어나오면 좋겠어요.",
				"깃헙, gist와 연동되어도 좋을 것 같구요.",
			},
			"posts": posts,
		})
	})
	
	return engineRouter
}