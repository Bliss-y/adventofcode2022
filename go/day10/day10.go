package main

import (
    "os"
    "strings"
    "strconv"
    "fmt"
)

func main() {
    content,_ := os.ReadFile("input.txt");
    lines := strings.Split(string(content) , "\r\n" );
    cyclenum := 0;
    register_val := 1;
    next_check_index := 40
    render := [6*40]byte{}
    for _, line := range lines {
        if line == "" {
            break;
        }
        cyclenum++;
        commands := strings.Split(line, " ")
        val := 0
        current_position := (cyclenum -1) % 40
        fmt.Println(current_position, cyclenum, register_val)
        if register_val + 1 >= current_position && current_position >= register_val - 1 {
            render[cyclenum -1] = '#'
        } else {
            render[cyclenum -1] = '.'
        }
        if commands[0] == "addx" {
            val ,_ = strconv.Atoi(commands[1]);
            cyclenum++;
            current_position := (cyclenum-1) % 40
            if register_val + 1 >= current_position && current_position >= register_val - 1 {
                render[cyclenum-1] = '#'
            } else {
                render[cyclenum -1] = '.'
            }
        }
        if cyclenum >= next_check_index {
            next_check_index += 40
        }
        register_val += val
    }
    for i, e := range render {
        if i % 40 == 0 {
            fmt.Printf("\n")
        }
        fmt.Printf("%c ",rune(e))
    }

}
