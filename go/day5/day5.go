package main
import (
	"fmt"
	"os"
    "strings"
    "strconv"
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

type Stack struct{
    s [50]byte
    length int
}

func (s *Stack) Put(b byte){
    s.s[s.length] = b;
    s.length++;
}

func (s *Stack) Pop() byte{
    if s.length == 0 {return 0}
    s.length--;
    return s.s[s.length]
}

func (s *Stack) reverse() {
    for i:=0; i < s.length/2; i++ {
        tmp := s.s[s.length- i -1 ];
        s.s[s.length - i - 1] = s.s[i];
        s.s[i] = tmp;
    }
}

func (s Stack) initialize() {
    s.s = [50]byte{} 
}

func main() {
    inputfile := "input.txt"
    content, err := os.ReadFile(inputfile)
    if err != nil{
        fmt.Println("err reading file");
        return;
    }
    reader := LineReader(content);
    stacks := [9]Stack{}
    for i:=0; i<len(stacks); i++{
        stacks[0].initialize();
    }
    // loop to create stacks
    for i:=0; i < len(stacks); i++ {
        line, _ := reader();
        for j:=0; j < len(stacks); j+=1 {
            if line[j*4] == ' '{
                continue;
            }
            stacks[j].Put(line[j*4 + 1])
        }
    }
    for i:=0; i<len(stacks); i++{
        stacks[i].reverse()
    }
    //remove 2 useless
    _, i := reader();
    if(i > len(content)){
        return;
    }
    
    n :=0
    for s, i := reader(); i < len(content); s, i = reader(){
        if(len(s) <1) {
            break;
        }
        pairs := strings.Split(string(s), " ")
        from,_ := strconv.Atoi(pairs[3])
        to,_ := strconv.Atoi(pairs[5])
        amount,_ := strconv.Atoi(pairs[1])
        n++;
        fmt.Println(n, pairs, amount, from, to);
        for x:=0; x<amount; x++ {
            if(stacks[from-1].length <1){
                return;
            }
            val := stacks[from-1].Pop();
            stacks[to-1].Put(val)
        }

    }
    for i:= 0; i < len(stacks); i++ {
        fmt.Printf("%c", stacks[i].Pop());
    }

}
