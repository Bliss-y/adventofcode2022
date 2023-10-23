package main

import (
    "os"
    "fmt"
    "strings"
)

func main(){
    content, _ := os.ReadFile("input.txt") 
    grid := strings.Split(string(content), "\r\n")
    length := len(grid)-1
    visibles := make([]int, length * length);
    fmt.Println("LengthX, lengthY total_edge", length, len(grid[0]), length *4 -3)
    total := 0

    for y,y_e := range grid {
        if(y >= length){break;}
        for x := range y_e {
            fmt.Println(x,y);
            position := (length) * y + x;
            visibles[position] = getStreakX(&grid, x,y, 1) * getStreakX(&grid, x,y, -1) * getStreakY(&grid, x,y, 1) * getStreakY(&grid, x,y, -1) 
            println("finished......", visibles[position])
       }
    }
    fmt.Println(visibles)
    for _,v := range visibles {
        if(v > total){
            total = v;
        }
    }
    fmt.Println(total)
}

func getStreakY( grid *[]string, x int, y int, direction int) int{ 
    length := len(*grid) - 1
    if(y ==0 || y >= length -1){ return 0;}
    streak := 0
    for i:= y + 1*direction; i < length && i >=0; i += 1*direction {
        fmt.Print(i);
        streak++; 
        if((*grid)[y][x] <= (*grid)[i][x]){
            break;
        }
    } 
    return streak
}

func getStreakX( grid *[]string, x int, y int, direction int) int{ 
    length := len(*grid) - 1
    if(x == 0 || x >= length-1){ return 0;}
    streak := 0
    for i:= x + 1*direction; i < length && i >=0 ; i += 1*direction{
        fmt.Print(i);
        streak++; 
        if((*grid)[y][x] <= (*grid)[y][i]){
            break;
        }

    } 
    fmt.Print(" ",streak, " ");
    println()
    return streak
}
