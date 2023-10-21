package main


import (
	"fmt"
	"os"
)

func LineReader(x []byte) func() ([]byte,int){
    content := x
    currentIndex := 0;
    return func()([]byte, int){
    if(currentIndex >= len(content)){
        return nil, currentIndex
    }
    current_string := make([]byte, 0, 20) 
    for ;content[currentIndex] != '\r' && currentIndex != len(content); currentIndex++{
       current_string = append(current_string, content[currentIndex]);
    }
    currentIndex+=2
    return current_string, currentIndex;
    }
}




func main() {
    inputfile := "input.txt"
    content, err := os.ReadFile(inputfile)
    if err != nil{
        fmt.Println("err reading file");
        return;
    }
    reader := LineReader(content);
    wins := []int{2,0,1}
    opp_score := 0;
    my_score := 0;
    for s, i := reader(); i < len(content); s,i = reader() {
        if(len(s)==3){
        opp_move := s[0] - 'A'
        my_move := s[2] - 'X'
        my_move = byte(wins[(opp_move + my_move)%3])
        opp_score += int(opp_move) + 1
        my_score += int(my_move) +1
          if my_move == opp_move {
              opp_score += 3;
              my_score += 3;
          } else if wins[my_move] == int(opp_move){
              my_score += 6;
          } else {
              opp_score += 6;
          }
        }
    }
    fmt.Print("opp, my: ",opp_score, my_score)
}

