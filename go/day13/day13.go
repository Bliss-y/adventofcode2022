package main

import (
    "os"
    "strings"
    "fmt"
    "strconv"
)


func main() {
    content,_ :=  os.ReadFile("input.txt");
    grid := strings.Split(string(content), "\r\n")
    sum := 0
    index := 0
    for i:=0; i< len(grid)-1; i+=3{
        index++
        list1:= grid[i]
        list2:= grid[i+1]
        add := true
        lists := 0
        y := 0
        //fmt.Println(list1, list2)
        for x := 0;x< len(list1); x++{
            e:= list1[x]
            if y >= len(list2){
                //fmt.Println("more content in x", x)
                add = false
                break
            }
            for list2[y] == ',' || list2[y] == ' ' {
                y++
                //fmt.Println("Next number at", y)
            }
            if e ==' ' || e == ','{continue}
            if e == ']' && list2[y] == ']' {
                //fmt.Println("Both ending list")
               y++;
               continue
            }
            if e == ']' && list2[y] != e{
                //fmt.Println("list ended first in left", x)
                add = true
                break
            }
            if list2[y] == ']' && e != ']' {
                //fmt.Println("list ended first in right. ", x)
                add = false
                break;
            }
            //fmt.Printf("%c %c\n", e, list2[y])
            if e == '[' {
                if list2[y] != '['{
                    // act as 1 element list
                    right, yI := getNumber(&list2, y);

                    for list1[x+1] == '['{
                        x++;
                    }
                    if list1[x+1] == ']' {
                        add = true
                        break
                    }
                    left, xI := getNumber(&list1, x+1)
                    //fmt.Println(left, right)
                    if right > left {
                        fmt.Println("right greater than left")
                        add = true;
                        break
                    } else if left > right {
                        fmt.Println("left greater than left")
                        add = false
                        break
                    } else {
                        x += xI+1
                        y += yI
                        if list1[x] != ']' {
                            fmt.Println("longer list in left")
                            add = false 
                            break
                        }
                    }
                    continue

                }
                lists++;
                y++;
                continue
            } else if list2[y] == '[' {
                fmt.Println("starting list in y")
                for list2[y+1] == '['{
                    y++;
                }
                if list2[y+1] == ']' {
                    fmt.Println("list ended first in y comparing empty string")
                    add = false 
                    break
                }
                right, yI := getNumber(&list2, y+1);
                //fmt.Println("y:", right, yI)
                 
                left, xI := getNumber(&list1, x)
                //fmt.Println("x:", left, xI)
               /**
                [[[]]] [1] => 
               **/
                if left > right {
                    fmt.Println("left > right")
                    add = false
                    break
                } else if left < right{
                    fmt.Println("left < right")
                    add= true;
                    break
                }else if list1[y+yI+1] != ']' {
                    fmt.Println("equals but list y continues")
                    add = true
                    break
                } else {
                    y = yI+1
                    x += xI
                   if list2[y] != ']' {
                       add = true
                       break
                   }
                }
                continue

            }
            left,xI:= getNumber(&list1, x); 
            right,yI  := getNumber(&list2, y);
            //fmt.Println("x:", left, xI)
            //fmt.Println("y:", right, yI)

            if  left <right {
                add = true
                break
            }else if right < left {
                add = false
                break
            }
            //println("equal")
            x += xI;
            y += yI+1;
            //fmt.Println("going to", x, y)

        }
        if add {
            sum += index
        }
        fmt.Println()
    }
    fmt.Println(sum)
}

func getNumber(s *string, index int) (int, int){
    str := make([]byte, 0,len(*s) - index)
    ss := *s
    for x := index; x < len(ss); x++{
        if ss[x] == ',' || ss[x] == ']' || ss[x] == '['{
            break
        }
        str = append(str, ss[x]);
    }
    v,err := strconv.Atoi(string(str))
    if (err != nil){
        fmt.Printf("got: %v %v\n", string(ss[index:]), string(str))
        panic("NaN")
    }
    return v, len(str)

}
