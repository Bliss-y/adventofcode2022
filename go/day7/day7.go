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

type dirTree struct {
    dirname string
    filesize int
    prev_dir *dirTree
    subdirs []*dirTree
    subdirs_size int
    calculated bool
}

func (d *dirTree) addDir(newdir *dirTree) {
    d.subdirs = append(d.subdirs, newdir)
}

func (d *dirTree)calculatesize() int{
    if(d.calculated){
        return d.filesize + d.subdirs_size
    }
    for _,c := range d.subdirs{
        d.subdirs_size += c.calculatesize() 
    }
    d.calculated = true;
    return d.filesize + d.subdirs_size
}

func main(){
    inputfile := "input.txt"
    content, err := os.ReadFile(inputfile)
    if err != nil{
        fmt.Println("err reading file");
        return;
    }
    reader := LineReader(content);
    dirHead := dirTree{
        "/",
        0,
        nil,
        make([]*dirTree, 0, 10),
        0,
        false,
    }
    currentdir := &dirHead;
    for s, i := reader(); i < len(content); s, i = reader(){
        if(len(s) <1) {
            break;
        }
        outputs := strings.Split(string(s), " ")
        if(outputs[0] == "$"){
            
            if(outputs[1] == "cd"){
                // push directory to the tree
                if(outputs[2] == "/"){
                    currentdir = &dirHead;
                }else if outputs[2]== ".."{
                   currentdir = currentdir.prev_dir 
            }else {
                    newdir := dirTree{
                        outputs[2],
                        0,
                        currentdir,
                        make([]*dirTree, 0, 10),
                        0,
                        false,
                    }
                    currentdir.subdirs = append(currentdir.subdirs, &newdir);
                    currentdir = &newdir
                }
            }
            if(outputs[1] == "ls"){
                continue;
            }

        } else {
            if(outputs[0] == "dir"){
                fmt.Println("continuing because dir")
                continue
            }
            filesize, _ := strconv.Atoi(outputs[0]);
            fmt.Println("Adding to directory: ", currentdir.dirname, currentdir.subdirs_size, filesize)
            currentdir.subdirs_size += filesize;
        }

    }
    totaldirsize := 0
    dirHead.calculatesize();
    dirQueue := make([]*dirTree, 100, 100);
    dirQueue[0] = &dirHead
    dirs := 1;
    current := 0;
    for dirs>0 {
        dir := dirQueue[current]
        dirsize := dir.calculatesize();
        if(dirsize < 100000) {
        totaldirsize+=dirsize;
        }
        fmt.Println("len of ",dir.dirname,len(dir.subdirs), dirsize)
        dirs--;
        current--;
        for _, d:= range dir.subdirs {
            current++;
            fmt.Println("adding to queue:  ",d.dirname, current)
            dirQueue[current] = d;
            dirs++;
        }
    }

    println(totaldirsize)


}
