// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	a := [...]string{".........#..##.##..............", "#...#.#..#.....................", ".#...#..#...................#..", "........##..#...#..............", ".#.#.....#..#..##......#.......", "....#..#...#..##........#..##..", "...#....##........#.......#.#..", "....#................#...###..#", "...#...#.#..#....#.......####.#", ".....#...#..........#...#..#.#.", "....#..#............#.#.#.#..#.", "..#....#.###..#............#...", ".....#.............#.#.........", ".#.##............##.........#..", "...##...#..#....#.##..#.....#..", "..............#.#.........#.##.", "...........#.....##....##......", ".......#............#...#......", "............#.#....#....#..#..#", "....#................####......", "...#.........................##", "..........#........#.#.........", "....#.#....#...........#......#", "..#.#..##......##..##..#..#.#..", "...#.....##......#.#.#.........", ".........#.#....#...#.#.#......", ".......#.......###.#.......#...", "..#............##..#.......#...", "...#....#......#...#...#...#...", "......#..#.#.....##......#.....", "...........##...##...#....#.##.", "#.##..#.....##..#.#............", ".#.#.....##......#.##........#.", "..#...#...#...#..........#...##", "...##.........................#", ".....#......#.....##....#.....#", "..#........#...................", "#......#..#.#..#..#.#..#...#...", "...............#..........#....", ".....#...........#......#....#.", "........#..#...............#...", ".........#...#.......#.#..#...#", "..#..#......#.##..........#....", ".#...#....#.....#.............#", ".##.....#.........#......#..#..", "........#..##.......#......#...", ".......#.....###......#..#.....", ".......#.#.......#.............", "...#................##.#.......", "..##..#...#.#...#.#..#.#.#.##..", ".......#.#............#...#....", "#...#.....#......#..........##.", ".#.......#......#.......#.#.#..", ".#.##.#.#...........#..........", ".......#.....#....#.#.##......#", ".###..#...#.............##.....", "......#......#.................", "##...#.#...##...#.#.....#....#.", "#.............#....##...#....#.", "#.#...#....#........#.###..##..", "......#.........#......#.#.#.#.", "..#.#.#.....#........#..#...#..", "#.##....#.#......#...........#.", ".#.#.####.........#..#.##....##", "......##...............#......#", ".......#.........#........#.#..", "....#....#..#.##.........#..#..", ".#..........#.##....#.#..#.....", "#..#.#............#..#....#.#.#", "..................#..#.........", "##..##.#....#.................#", "..................#........#..#", ".....#.........#.......##......", ".....................#.#..#...#", ".....#.........#..........#.#..", "...#.#..#..#.#.#.......#.......", ".....#.....#.#.........#.....##", ".............##....#....##.#...", ".#......#........##..#...###...", "........#.......##.##.#......#.", "..#....................#.##..#.", "......#.......#..#....##.#.....", "...#....#.......##...#.......#.", ".#.#..#.#..........##..........", "....#.......##...........#.....", "###....#.......#..#...#.....##.", "...#......#.........#..#.#..#.#", "#.........#..##.#..............", ".#.....#..##.#..#..###.....##.#", "..........#..#....##.......#...", ".#..#.#...#...##.#..#.##.#.....", "#....#...#........#......##....", "..#.####....#.#........#....#..", "#......#............#.#........", "...#..#.......##...........#...", ".........#..#.#..#.###.#...#..#", "..#....##.......#.............#", "............#..#......#........", "........#......#..............#", "..#.#.#...........#............", "...........#......##.#.#.......", ".#..........#...........#..#...", ".....#...#..#.........##...#...", ".......#....##....#.#.........#", "..#.#......#.......#...##.#....", ".....#..........#........#.....", "#.......#.......#............#.", "..##.#.....#.##.#.#.#..#.......", "..#...#.......#.###............", ".#...#..#....#...#...#..#....##", ".....#.....#...................", ".......................#......#", "......#...##.........#...#..#..", ".....#..#.....#..............#.", ".#.##..#..#....................", "....#..#...#....#.............#", "..###..#...#......#.....#......", "..#......###....#.....#.....###", "...#.##..#...#.....#........#..", ".#.#...........##....#...#.##..", ".......#...##......#..#..#.....", "#.............#..#...##.#..#..#", "..........#......#.......#.....", "...............#.#.#....#...#..", "#.......#.#..#.....#........#..", ".#.#.#.......#..#.........##...", "......#.....#.#....#...........", "..#.....##.#........##.#......#", "...###...#..#.........#........", "....#...................#..#...", ".##........#...................", "....#..#...........#.#.........", ".....#.......#...#....#.#......", ".........#...#.......#.#...#...", ".......#.#..#....#....#.......#", "..#.............#..............", ".#...#..#.#.#..#............#..", "...#.##.##..#..#...........##..", "...........#...#..#.#........#.", "....#...#.....#...#.#....#...#.", ".......#.#...##..#.............", ".......................#....#..", "..#..#.....#...........#....#..", ".#..#...#.##........##....#....", "#.....##.#.#.......#.....#...#.", ".#....#.......................#", "#..##..###...#.........#.......", "..##...#...#..........#....#...", "......#..##......##.#.........#", "................#........#..#..", ".....#.#..#.....#.......#..#...", "..#..#.....#.......#..#..#...#.", ".#....#...#...#......##.....#..", "....#........#...#......##....#", "..#..........##......#......#..", "#.#.....#.....#.......#........", "...#...#......#....##.#..#...##", "...#....#...#.#...........##...", "#....##...#...#....#...........", "...#.#..#...#..............##..", "#..#..........##.#.#.....#.....", "..#...#.........#.#..........#.", "....#.....#..........#.........", "........................#......", ".#.....#.#...###...#....#......", "....##....#....#..#.##........#", "..#........#.........#.......#.", ".....#.#......#...#...#........", "........#..#.....#....###....#.", "...........#..#.#....#.#....##.", ".......#.....##.#............#.", "...............#........##.##..", ".............#...##......#...#.", "#...##.#.......#......###.....#", "..........#...#........#..#....", "....#....................#...#.", ".#......#...#.......#....#.#...", "....#.......................#..", "#...#...#...#.##....##.........", "..........#.#...##.#...#.......", "..#...............#....#..#...#", "#..#..#...#..#.........#...#...", ".....#..#..........#.##.#..##..", "........#......##.....#........", ".#....#.#.........#...#..#.#...", "....#..............##..........", "#...............#..............", "..###.........#....##.........#", ".........#.#....##........#...#", "....#.#..#......#...#..........", "...#.#.....#....#..#....#.#..#.", "............#..#.....#...##....", "...........#....#.#.#...#......", "#...............#....###.......", ".........#.....##.#..#..#......", "...#...##...###...............#", ".#......#.#.#.................#", ".........##..#............#....", "..#..#.....#.....#.#...........", ".#......##............#.#....#.", ".#.##..##.##..#.........#.....#", "...##.##......##.##......#.....", "##.....#.#...#...#...#..#......", "....................#......#...", ".....#.................#...###.", "...........#..#.........#.#....", "...#........#..#........#....#.", "#................#......###...#", ".............##.#.....#.#......", "...#...#.##..#........##.......", "#..#.##...#....#.#.............", ".#.........#.#..#...........#..", "....#...#.....#.#..........#...", ".#.#....###.......##.....#.##..", ".##....##......#......#.#....#.", "..#...#.#........#...#..##.....", "..............###..........##..", "#....#..##.....#.....#.....#...", "...#...#....................#..", ".#....#.#.....#.#..#..##.......", "#...##..###......#.............", "..........#.#....##.#........##", "..#..#.....#...#....#.#.#......", "#.....#........#..##.#.........", "....#.....#..........##......#.", "......#..#.....#........#.....#", ".....#..#....#...........#.##..", ".#....................#....#..#", "........#..#...........#.......", "#....#.#.......#........#.#..#.", "........#.....#...#............", "..#........#........#....#...#.", ".....##.......#..#..........#..", "......#.#......###...#....##..#", ".#..#.............#.#..........", "#.....##.#.#.#.#.#.....#.....#.", ".#..#.....#.......#.#.....#....", "###....##...#.#...#..#......###", ".#................#.....#.##...", "....##....#.#........###.#.#...", "#.#....#........#.....#.......#", "..........#..........#.##...#..", "....#....#..##......#..#.......", ".....#..........#.##...........", "##......#.#......#.##..........", "##..........##.......##........", "..#.....#....#.##..#..#..#.....", "......###...#...........#...###", "#..#.............##............", "...#.###.....#..#.........#.#..", "......#...............#...#.#..", ".....#...##.#...#.....#.#..#...", "..#..#.#....#.#................", "...............##.....#........", "......#.#.....#...#.........#..", "........#..#...#.#...#......#..", "#...........#.......#...##...#.", "........#.#...#..##..#.#...#...", "..#....#...#......#..........##", "..#..............##...##.#.....", "...#....#..#....##.........#.#.", ".#.#....#..........#.......#...", "...##....#.#....#....#.#...#...", "..............#..##........#..#", "..........#.#...##......#..#.#.", "#...##..#......................", ".......#........##.#.#.#.......", ".........##..#.#.......####....", "..#.............#..#........##.", "##..#..#...#....#.....#...#..#.", "..#.#...#.#.....#..............", "..#....#....#..##...#.#........", "##.....#..#...#................", "#....#.....................#...", "..............###.....#.#.#....", "..#......##.#....#.#...##......", "#...#.#......#...#.#......#....", "....#...................##.#...", ".........##......#.....#.####..", "##..#........#.....#......##..#", "...#..#...#...#.............#..", "#..#..#.#......###...#.........", ".......#.#..#........#....##..#", "............#..##.....#.#.#....", "#..#.....#.....#....##........#", "......#..........##............", ".....#...#...........#.........", "...........#....#...#....#.#...", "....#.........##.##.......#....", "......#....#...........#.##...#", ".##.#.#..##...#.....##.#...#...", ".......#.#....#...#...#....#...", ".#...##.#.#.....#..#....#......", ".#....###..#.......#......#...#", "..#.#.........#.........#.....#", ".......#.#.##..#.#.......##..#.", ".##............#.........#....#", ".#...##.###..#........##.#..#..", "..#........#.#.....##..##.#....", "...........#...........#.....#.", ".#...######..##...#.....#......", ".#.##.#.......#......#......#..", ".#.....#.....#........#........", "...#..#...#.##...#...........#.", ".......#.....#.......#.........", "............#...###...........#", "...#.......#.......##....#..#..", "##.......#....#....####........", ".......#.#......#..........#..#", "#.....##..#..#.....#....#...#..", "#............#........##.......", ".#.#...#.............#..##.....", ".#....#..#.#......#.##.......##", "...................##...##..###", "..#.....#...#................#.", "..#...#....#...#.#.#...#.....#.", ".....#............#....#...#..#", ".#.....#....#..#......#.#.....#", "............#.#.....####.##....", "....#......###....#...#....#...", "#.....#..#.....#..#...#.......#", "..#.#...#....#....##..#...##...", ".##..#..#..##....##...#........"}
	//a := [...]string{"..##.......", "#...#...#..", ".#....#..#.", "..#.#...#.#", ".#...##..#.", "..#.##.....", ".#.#.#....#", ".#........#", "#.##...#...", "#...##....#", ".#..#...#.#"}
	var x = 0
	var s1 = 0 
	var s2 = 0
	var s3 = 0
	var s4 = 0 
	
	var treesHit = 0
	var treesHitS1 = 0
	var treesHitS2 = 0
	var treesHitS3 = 0
	var treesHitS4 = 0

	for i, s := range a {
		s = strings.Repeat(s, 100)

		if i == 0 {
			fmt.Println(i, "Line 0")
			s4 = s4 + 1
		} else {
			var isTree = string(s[x]) == "#"
			if isTree {
				treesHit++			
			}
			
			if string(s[s1]) == "#" { treesHitS1++ } 
			if string(s[s2]) == "#" { treesHitS2++ } 
			if string(s[s3]) == "#" { treesHitS3++ } 
			
			if i%2==0 { 		        	
				if string(s[s4]) == "#" {
					treesHitS4++
				//	s = string(s[0:s4-1]) + "X" + string(s[s4:])
				} else {
				//	s = string(s[0:s4-1]) + "O" + string(s[s4:])
				}
				s4 = s4 + 1
			}
		}

		//fmt.Println(i, s)
		x = x + 3
		s1 = s1 + 1
		s2 = s2 + 5
		s3 = s3 + 7
		
	}

	fmt.Println("s1 Trees Hit: ", treesHitS1)
	fmt.Println("x Trees Hit: ", treesHit)	
	fmt.Println("s2 Trees Hit: ", treesHitS2)
	fmt.Println("s3 Trees Hit: ", treesHitS3)
	fmt.Println("s4 Trees Hit: ", treesHitS4)
	
	var answer = treesHit * treesHitS1 * treesHitS2 * treesHitS3 * treesHitS4
	fmt.Println("Answer is ", answer)
}
