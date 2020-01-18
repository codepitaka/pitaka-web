# pitaka-web

### A Web Blog Service For Code Education  

This web blog service provides a lightening-fast pipeline to create readable and runnable tech-blog posts for elementary code tutors, and a comprehensible and do-able code blog posts for newbie code learners.  

Etymologically, pitaka (पिटक) originated from buddhism, which means "basket or box made from bamboo or wood". It refers to "scripture or book" in plain english. There were 3 pitaka(s). Vinaya Piṭaka (“Basket of Discipline”), Sūtra Piṭaka (“Basket of Discourse”), and Abhidharma Piṭaka (“Basket of Special[or Further] Doctrine”). As so, this web blog service aims at providing readable, runnable and do-able code scriptures.

### File Structure
```
pitaka-web
+-- src/         : web server 소스파일을 모아둔 폴더
+-- main.go      : entry 소스파일
+-- main_test.go : entry 소스 test 파일
+-- go.mod       : package manager 파일 [go mod init 시 자동생성]
+-- go.sum       : package manager validation 파일 [go.mod에 write시, 자동생성]
+-- reflex       : 개발 중 웹페이지 auto reload를 위한 설정파일
+-- go.sum       : package manager validation 파일 [go.mod에 write시, 자동생성]
+-- .gitignore   : GOPATH로 설정된 .go 폴더의 패키지들을 ignore
+-- travis.yml   : travis 설정파일
+-- Makefile     : heroku 설정파일
+-- Procfile     : heroku 설정파일
```


### Download Go ver 1.13.6 
- [GO archive](https://golang.org/dl/)
- [how to download](https://www.quora.com/Whats-the-easiest-way-to-update-Go-programming-language-to-the-latest-version-in-Linux)

### Package Manager: go modules
- 
