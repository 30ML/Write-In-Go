# Go

### Go 개발 환경
#### 1. workspace 지정
Go 작업 파일(소스 코드, 실행 파일, 패키지 파일 등)이 위치할 경로를 GOPATH로 지정함
```bash
# OS X, Linux
export PATH=$GOPATH/bin:$PATH
export GOPATH=/Users/mjdn0011/mjdn0011/Go_MJDN
```

#### Go 환경 설정 확인
go env

```go
package main

import "fmt"

func main () {
  fmt.Println("Hello, World)
}
```

##### package
go의 package는 코드를 구조화하고 재사용하기 위한 단위
다른 언어의 모듈이나 라이브러리와 유사한 개념
모든 Go 코드는 패키지에 포함되어야 하고, 
Go 코드는 항상 패키지 선언으로 시작함
* 특별히 main 패키지는 프로그램의 시작점 역할
모든 실행 가능한 Go 프로그램에는 main 패키지가 있고,
프로그램을 구동할 때 main 패키지의 main 함수를 찾아 실행함

> Go 프로그램의 두 가지 타입
> 1. 실행 가능한 프로그램: 명령 프롬프트에서 명령을 내려 실행할 수 있는 프로그램
> 2. 라이브러리: 다른 프로그램에서 호출하여 사용할 수 있게 만든 코드의 묶음

##### import 
외부 패키지를 사용할 때는 import 문으로 해당 패키지를 임포트

##### func main() {...}


* Go에는 `while` 문이 없음
* 