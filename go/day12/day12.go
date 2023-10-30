package main

import (
    "os"
    "strings"
    "fmt"
)


func main() {
    content,_ :=  os.ReadFile("input.txt");
    grid := strings.Split(string(content), "\r\n")
    visited := make([]int, len(grid)*len(grid[1]));
    stack:= make([][3]int, len(grid)*len(grid[0]));
    end_p := [2]int{}
    cur := 0
    for y,e := range grid {
        for x,s := range e {
            if s == 'S' || s == 'a' {
                stack[cur] = [3]int{x,y,1}
                cur++;
            }
            if s == 'E' {
                end_p[0] = x 
                end_p[1] = y 
            }
        }
    }
    x_len := len(grid[0])
    for cur >=0 {
        curr := stack[cur] 
        cur--;
        //fmt.Println(curr);
        cur_elevation := grid[curr[1]][curr[0]] 
        if cur_elevation == 'S' {
            cur_elevation = 'a'
        }
        visited[curr[1]*x_len + curr[0]] = curr[2]
        if (cur_elevation == 'E'){
            //fmt.Println("Found at: ", curr)
            continue;
        }
        next_val := curr[2] + 1
        xes := [4]int{1,0,0,-1}
        yes := [4]int{0,1,-1,0}
        for i,e := range xes{
            nx := curr[0] + e;
            ny := curr[1] + yes[i]
            visited_position:= nx + x_len *ny;
            if nx <0 || ny < 0{continue}
            if nx >= x_len || ny >= len(grid)-1 || (visited[visited_position] != 0 && visited[visited_position] <= next_val) || cur_elevation +1 < grid[ny][nx] {
                continue;
            }
            if grid[ny][nx] == 'E' && cur_elevation < 'y' {
               continue 
            }
            cur++;
            if(cur >= len(stack)){
               stack = append(stack,[3]int{nx, ny, next_val})
               continue
            }
            stack[cur] = [3]int{nx, ny, next_val}
        } 
    }
    println(end_p[0])
    println(end_p[1])
    
    fmt.Println()
    println(visited[end_p[0] + x_len * end_p[1]])
}
/*
        if nx:=curr[0] -1; nx  >= 0 && (visited[nx + x_len* curr[1]] == 0 || visited[nx +x_len* curr[1]] > next_val) && cur_elevation+1 >= grid[curr[1]][nx]{
            println("Going left")
           cur++;
           stack[cur] = [3]int{nx, curr[1], next_val};
        }
        if ny := curr[1] -1; ny >= 0 && (visited[ny * x_len + curr[0]] == 0 || visited[ny * x_len + curr[0]] > next_val)&& cur_elevation+1 >= grid[ny][curr[0]]{
            println("Going up")
           cur++;
           stack[cur] = [3]int{curr[0], ny, next_val};
        }
        if ny := curr[1] +1; ny < len(grid) && (visited[ny * x_len+curr[0]] == 0 || visited[ny * x_len + curr[0]] > next_val)&& cur_elevation+1 >= grid[ny][curr[0]]{
            println("Going down")
            cur++;
           stack[cur] = [3]int{curr[0], ny, next_val};
        }
        if nx:=curr[0] +1; nx< len(grid[0]) && (visited[nx +  x_len * curr[1]] == 0 || visited[nx + x_len * curr[1]] > next_val)&& cur_elevation+1 >= grid[curr[1]][nx]{
            println("Going right")
            cur++;
           stack[cur] = [3]int{nx, curr[1], next_val};
       }
    */
