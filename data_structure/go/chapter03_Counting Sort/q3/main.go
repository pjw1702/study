package main

import "fmt"

// Q3: 한 반의 학생들 중 키가 특정 범위의 학생들만 출력하시오.
// 범위 값은 여러번 입력받을 수 있다.
// 키는 소수점 1자리까지 주어진다.
// 입력 예:
// 140 170
// 180 190
// 200 210
// 160 180
func main() {
	students := []struct {
		Name   string
		Height float64
	}{
		{Name: "Kyle", Height: 173.4},
		{Name: "Ken", Height: 164.5},
		{Name: "Ryu", Height: 178.8},
		{Name: "Honda", Height: 154.2},
		{Name: "Hwarang", Height: 188.8},
		{Name: "Lebron", Height: 209.8},
		{Name: "Hodong", Height: 197.7},
		{Name: "Tom", Height: 164.8},
		{Name: "Kevin", Height: 164.8},
	}

	// 가장 단순한 방법
	// 가장 빨리 짜는 방법이긴 하지만, 입력이 여러 번 되는 경우에는 문제가 발생한다
	// 모든 경우에 대해 if문을 작성해 놓을 순 없다
	// for i := 0; i < len(students); i++ {
	// 	if students[i].Height >= 170.0 && students[i].Height > 180.0 {
	// 		fmt.Println(students[i].Name, students[i].Height)
	// 	}
	// }

	// 사람의 키가 최대 3m까지 존재한다고 가정
	// 3m를 소숫점 한 자리까지 환산했을 때 300.0cm: 데이터 수는 3000
	var heightMap [3000][]string
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		heightMap[idx] = append(heightMap[idx], students[i].Name)
	}

	// 140 ~ 170
	for i := 1400; i < 1700; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name: ", name, ",", "height: ", float64(i)/10)
		}
	}

	fmt.Println("")

	// 180 ~ 210
	for i := 1800; i < 2100; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name: ", name, ",", "height: ", float64(i)/10)
		}
	}
}
