package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// 2차원 슬라이스 생성
	result := make([][]uint8, dy)

	// 2차원 슬라이스의 각 요소에 대해
	for i := range result {
		// 1차원 슬라이스 생성
		result[i] = make([]uint8, dx)

		// 1차원 슬라이스의 각 요소에 대해
		for j := range result[i] {
			// 1차원 슬라이스의 각 요소에 값을 할당
			result[i][j] = uint8(i * j)
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
