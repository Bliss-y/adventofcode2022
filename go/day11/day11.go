package main
import (
    "fmt"
    "strings"
    "strconv"
    "os"
)

type Monkey struct {
    items []uint64;
    divisor uint64;
    operation byte;
    operator uint64;
    pass [2]uint64;
}

func (m *Monkey) operate(old uint64) uint64 {
    value := uint64(0)
    switch m.operation {
        case '+':
            if(m.operator == 0){
            value = old + old
            } else{
            value = old + m.operator 
        }
        case '*':
            if(m.operator == 0){
            value = old * old
            } else{
            value = old * m.operator 
        }
    }
    return value
}

func main(){
    content, _ := os.ReadFile("input.txt")
    lines := strings.Split(string(content), "\r\n")
    monkeys := make([]*Monkey, 0, 10);
    for i:=0; i < len(lines); i+=7 {
        if lines[i] == "" {
            break;
        }
        monkey := Monkey{}
        items := strings.Split((lines[i+1])[18:], ", ")
        for _,e := range items{ 
            num,_ := strconv.Atoi(e);
            monkey.items = append(monkey.items,uint64(num) )
        }
        operator,_ := strconv.Atoi(lines[i+2][25:])
        monkey.operator = uint64(operator)
        monkey.operation = lines[i+2][23]
        monkey.divisor, _ = strconv.ParseUint(lines[i+3][21:], 10, 0)
        monkey.pass[0],_ = strconv.ParseUint(lines[i+4][29:], 10, 0)
        monkey.pass[1],_ = strconv.ParseUint(lines[i+5][30:], 10,0)
        fmt.Println(monkey)
        monkeys = append(monkeys, &monkey)
        
    }
    limit := uint64(1);
    for _,c:= range monkeys{
       limit*= c.divisor 
    }
    inspection_counts := make([]int, len(monkeys))
    for c:=0; c<10000;c++  {
        for x:=0; x < len(monkeys); x++{
            m := monkeys[x]
            items := m.items
            m.items = make([]uint64, 0, 10)
            inspection_counts[x]+= len(items)
            for _,e := range items {
                e := m.operate(e);
                i :=0
                //e = e/1;
                if (e % m.divisor != 0){i = 1}
                e = e % limit;
                monkeys[m.pass[i]].items = append(monkeys[m.pass[i]].items, e); 
                         }
        }
        //for x:=0; x < len(monkeys); x++{
        //    fmt.Println(monkeys[x])
        //}
        //fmt.Println("end round")

    }
        for x:=0; x < len(monkeys); x++{
       fmt.Println(monkeys[x])
       }
    highest := 0;
    lowest := 0;
    for _,e := range inspection_counts {
        if highest < e {
            lowest = highest;
            highest = e;
        } else if lowest < e {
            lowest = e; 
        }
    }

    fmt.Println(inspection_counts)
    fmt.Println(highest*lowest)
}
