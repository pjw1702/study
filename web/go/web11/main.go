package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

// 함수 또한 템플릿에 쓰일 수 있다
// 템플릿 파일에서, '-' 을 입력하면, if문 등의 조건문이나 반복문 사용시에 발생하는 조건절에서의 공백 및 개행을 없앨 수 있다(조건문 앞에 쓰면 앞의 공백을 없앨 수 있고 조건문 뒤에 쓰면 뒤의 공백을 없앨 수 있다)
func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "jwpark", Email: "pjw1702@gmail.com", Age: 29}
	user2 := User{Name: "jwpark2", Email: "pjw1702@naver.com", Age: 31}
	users := []User{user, user2}
	//tmpl, err := template.New("Tmpl1").Parse("Name: {{.Name}}\nEmail: {{.Email}}\nAge: {{.Age}}")
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		// panic(err)
		fmt.Println("Cannot creating instance of templates: Parsing Error!")
		return
	}
	// text/template 패키지를 import하고 ExecuteTemplate()함수를 실행하면, '@' 등의 특수문자 탈락을 방지할 수 있다
	// html/template 패키지를 import하고 ExecuteTemplate()함수를 실행하면, '@' 등의 특수문자 탈락시킨다
	// 탬플렛 내에서 template 키워드를 이용하여 다른 템플릿을 include하여 inner 템플릿을 적용(다른 템플릿을 해당 템플릿에 포함)시킬 수 있다
	// inner 템플릿 적용 시, 뒤에 '.'을 입력하면 포함 시키려는 다른 템플릿의 인스턴스에 대한 데이터 모두를 포함할 수 있다
	// 리스트를 생성하여 한 번에 여러 데이터를 저장 시, {{range}}를 이용하여 모든 항목에 대한 데이터를 출력할 수 있다 (range 뒤에 쓰이는 '.'은 리스트의 각 항목(인스턴스)으로, 인스턴스의 데이터를 의미하는 template 키워드에서의 '.'과 다르다)
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user)
	// fmt.Println("")
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user2)
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
