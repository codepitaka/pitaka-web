: '
  -(1) go 삭제 -
    go는 /usr/local에 설치, 해당 위치의 go 삭제
	purge로 파일명에 golang이 포함된 파일 찾아 삭제
	gopath에 해당하는 경로 삭제
	macos의 경우, apt-get대신 다른 명령어 필요할 듯 싶음.
'

echo "removing existing go..."
sudo rm -rf /usr/local/go
sudo apt-get purge golang*
sudo rm -rf ~/.go
echo "remove complete."

: ' 
  -(2) go 설치 -
    프로세서 변경: go.1.12.16.[프로세서].src.tar.gz
    설치위치 변경: -O flag로 설치위치를 설정변경 가능
'

echo "downloading go..."
wget https://dl.google.com/go/go1.12.16.linux-amd64.tar.gz -O /go1.12.16.linux-amd64.tar.gz
echo "download complete."

: ' 
  -(3) go 압축해제 -
    프로세서 변경: go.1.12.16.[프로세서].src.tar.gz
'

echo "unzipping go..."
sudo tar -C /usr/local -xzf /go1.12.16.linux-amd64.tar.gz
echo "unzip complete."

: ' 
  -(4) GOPATH 경로 설정 -
    설명: gomodules 밖에서 패키지 설치 시, GOPATH에 담긴다. 
'

echo "create dir for GOPATH..."
mkdir ~/.go
echo "GOPATH dir creation complete."

: ' 
  -(5) 환경변수 설정 - 
    설정변수 3가지
	  1. GOROOT
	  2. GOPATH
	  3. PATH
'

echo "setting environment variables for go..."
export GOROOT=/usr/local/go
export GOPATH=~/.go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
echo "setting environment variables for go completed"

: ' 
  -(6) go command 설정 -
'

echo "setting go command..."
sudo update-alternatives --install "/usr/bin/go" "go" "/usr/local/go/bin/go" 0
sudo update-alternatives --set go /usr/local/go/bin/go
echo "setting go command completed"
: ' 
  -(7) go version 확인 - 
'
echo "current version of go is,"
go version

: ' 
  -(8-1) 필요한 패키지 설치 -
    reflex: 파일 변경시, 자동으로 서버 재시작하는 auto reload를 위한 패키지.
	golangci-lint: 린터 패키지.
'
echo "installing reflex, golangci-lint..."
go get github.com/cespare/reflex
go get github.com/golangci/golangci-lint/cmd/golangci-lint
echo "installing reflex, golangci-lint completed."

: '
  -(8-2) gopherjs 패키지 설치 -
'
echo "installing gopherjs..."
go get -u github.com/gopherjs/gopherjs
export GOPHERJS_GOROOT="$(go env GOROOT)"  # Also add this line to your .profile or equivalent.
npm install --global source-map-support
echo "installing gopherjs completed."