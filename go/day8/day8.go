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
    visibles := make([]bool, length * length);
    column_biggest := make([]rune, length)
    fmt.Println("LengthX, lengthY total_edge", length, len(grid[0]), length *4 -3)
    total := 0
    position_count := 0

    for y,y_e := range grid {
        if(y >= length){break;}
        curr_biggest := '/'
        for x,x_e := range y_e {
            position_count ++
            position := (length) * y + x;
        if(x >= length) {break;}
           if(x == 0 || x_e > curr_biggest){
               visibles[position] = true;
               curr_biggest = x_e
           }
           if(y ==0 || x_e > column_biggest[x]){
                column_biggest[x] = x_e;
                visibles[position] = true;
           }
             
        }
    }
    for y:=length-1; y >=0; y--{
        y_e:= grid[y]
        curr_biggest := '/'
        for x:=length-1; x >=0; x--{
            position := (length) * y + x;
            x_e := y_e[x]
           if(x == length-1 || rune(x_e) > curr_biggest){
               visibles[position] = true;
               curr_biggest = rune(x_e)
           }
           if(y == length-1 || rune(x_e) > column_biggest[x]){
                column_biggest[x] = rune(x_e)
                visibles[position] = true;
           }
        }
    }
    for _,v := range visibles {
        if(v){
            total++;
        }
    }
    fmt.Println(position_count, len(visibles))
    fmt.Println(total)
}

