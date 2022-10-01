package main

import "fmt"

func main() {
	poet1 := "죽는 날까지 하늘을 우러러\n한 점 부끄럼이 없기를,\n잎새에 이는 바람에도\n나는 괴로워했다.\n"
	// 문자열을 저장할 때, ""와 ``를 사용할 수 있다(``를 사용하면 줄바꿈을 하여 문자열을 입력 가능하다)
	// `` 사용 시, Tab과 특수 문자 사용을 삽입하는 경우 그대로 출력되니 주의해야 한다
	poet2 := `죽는 날까지 하늘을 우러러
	한 점 부끄럼이 없기를,\n
잎새에 이는 바람에도
나는 괴로워했다.`

	fmt.Println(poet1)
	fmt.Println(poet2)
}
