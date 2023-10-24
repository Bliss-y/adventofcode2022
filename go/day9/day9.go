package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

type Position struct {
    x int;
    y int;
}

func main(){
    content,_:= os.ReadFile("input.txt")
    commands := strings.Split(string(content), "\r\n")
    highest_x := 0;
    highest_y := 0;
    for _, i := range commands {
        if(len(i) <3){
            break;
        }
        commandp := strings.Split(i, " ");
        num, _ :=strconv.Atoi(commandp[1]);
        switch x := commandp[0]; x {
            case "R":
                highest_x += num
            case "D":
                highest_y += num
            case "U":

                highest_y += num
            case "L":
                highest_x += num
        }
        
    }
    if highest_y>highest_x{
        highest_x = highest_y
    }
    length := 800
    visited := make([]bool, length * length)
    headPosition:= Position {
        length/2,
        length/2,
    } 
    tailPosition:= Position {
        length/2,
        length/2,
    } 
    record(&visited, &tailPosition, length)
    for _,i:= range commands {
        if(len(i) <3){
            break;
        }
        commandp := strings.Split(i, " ");
        num, _ :=strconv.Atoi(commandp[1]);
        move(commandp[0], num,&visited, &headPosition,length, &tailPosition)
    }

    total :=0;
    for _, x:= range visited {
        if x {total++}
    }
    fmt.Println(total)
}

func record(visited *[]bool, position *Position, length int){
    p := position.y * length + position.x;
    if(position.y > length-1 || position.x > length -1 || position.x < 0 || position.y < 0){
       panic("WTF???")
    }
    if(*visited)[p] == true {
        println("Already visited!");
    }
    (*visited)[p] = true;
}

func move(command string, amount int,visited *[]bool, headPosition *Position, length int ,tailPosition *Position){
        fmt.Println(command, amount);
        switch x := command; x {
            case "R":
                if amount + headPosition.x > tailPosition.x+1 {
                    fmt.Println("Moving right: ", tailPosition, headPosition);
                    for i:= tailPosition.x + 1; i < headPosition.x + amount; i++ {
                        tailPosition.y = headPosition.y;
                        tailPosition.x  = i;
                        fmt.Println(*tailPosition)
                        record(visited, tailPosition, length);
                        //   fmt.Println("Moving right: ", tailPosition);
                    }
                    if(headPosition.x + amount != tailPosition.x + 1){
                        panic("Head PAnic!!!!")
                    }
                }
                headPosition.x += amount;
            case "D":
                if  headPosition.y - amount < tailPosition.y-1 {
                    fmt.Println("Moving down: ", tailPosition, headPosition);
                    for i:= tailPosition.y - 1; i > headPosition.y - amount; i-- {
                        tailPosition.x = headPosition.x;
                        tailPosition.y  = i;
                        fmt.Println(*tailPosition)
                        record(visited, tailPosition, length);
                  //      fmt.Println("Moving down: ", tailPosition);
                    }
                if(headPosition.y - amount != tailPosition.y - 1){
                    panic("Head PAnic!!!!")
                }
                }
                headPosition.y -= amount;
            case "U":
                if amount + headPosition.y > tailPosition.y+1 {
                    fmt.Println("Moving Up: ", tailPosition, headPosition);
                    tailPosition.y = headPosition.y;
                    for i:= tailPosition.y + 1; i < headPosition.y + amount; i++ {
                        tailPosition.x = headPosition.x;
                        tailPosition.y  = i;
                        fmt.Println(*tailPosition)
                        record(visited, tailPosition, length);
                      //  fmt.Println("Moving Up: ", tailPosition);
                    }
                    if(headPosition.y + amount != tailPosition.y + 1){
                        panic("Head PAnic!!!!")
                    }
                }
                headPosition.y += amount;
            case "L":
                if  headPosition.x - amount < tailPosition.x-1 {
                    fmt.Println("Moving left: ", tailPosition, headPosition);
                    for i:= tailPosition.x - 1; i > headPosition.x - amount; i-- {
                        tailPosition.y = headPosition.y;
                        tailPosition.x  = i;
                        fmt.Println(*tailPosition)
                        record(visited, tailPosition, length);
                      //  fmt.Println("Moving left: ", tailPosition);
                    }
                if(headPosition.x - amount != tailPosition.x - 1){
                    panic("Head PAnic!!!!")
                }
                }
                headPosition.x -= amount;
        }
}
