package main
import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

type Node struct {
	Id string `json:"-"`
	ParentId string `json:"-"`
	Name string `json:"name"`
	Value string `json:"Value,omitempty"`
	Children []*Node `json:"children,omitempty"`
}

func (this *Node) Size() int {
	var size int = len(this.Children)
	for _, c := range this.Children {
		size += c.Size()
	}
	return size
}

func (this *Node) Add(nodes... *Node) bool {
	var size = this.Size();
	for _, n := range nodes {
		if n.ParentId == this.Id {
			this.Children = append(this.Children, n)
		} else {	
			for _, c := range this.Children {
				if c.Add(n) {
					break
				}
			}
		}
	}
	return this.Size() == size + len(nodes)
}


func Save(dataset []byte){
configFile:="Config.json"
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Println(err,file)
	}
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		_, err := os.Create(configFile)
		if err==nil{
			fmt.Printf("%s file created ... \n", configFile)
		}else{
			fmt.Printf("file cannot create please check file location ")
		}
	}
	file1, err := os.OpenFile(configFile,os.O_WRONLY,0600)
	if err != nil {
		// panic(err)
		fmt.Printf("Appended into file not success please check again \n")
	}
	defer file.Close()
	
	if _, err = file1.WriteString(string(dataset)); err != nil {
		panic(err)
	}
}


func Load(jsonfile string){
	file, e := ioutil.ReadFile(jsonfile)
    if e != nil {
        if _, err := os.Stat(jsonfile); os.IsNotExist(err) {
			_, err := os.Create(jsonfile)
			if err==nil{
				fmt.Printf("%s file created ... \n", jsonfile)
			}else{
				fmt.Printf("file cannot create please check file location ")
			}
		}
	}
    
    fmt.Printf("%s\n", string(file))
    var jsontype Node
    json.Unmarshal(file, &jsontype)
    fmt.Printf("Results:%v\n", jsontype)
}



func main() {
	var root *Node = &Node{"001", "", "DbName", "MySql", nil}
	data := []*Node{
				&Node{"002", "001", "Db","DuoV6", nil},
				&Node{"003", "002", "Username","Duov6", nil},
				&Node{"004", "002", "Password","123", nil},
				&Node{"005", "004", "tables","Auth", nil},
				&Node{"006", "004", "tables","Config", nil},
				&Node{"007", "004", "tables","derective", nil},
				&Node{"008", "004", "tables","Users", nil},
				&Node{"009", "004", "tables","Canves", nil},
				&Node{"010", "004", "tables","RabbitMQ", nil},
				&Node{"011", "004", "tables","test table 2 ", nil},
				&Node{"012", "004", "tables","test table 3", nil},
			}

	fmt.Println(root.Add(data...), root.Size())
	bytes, _ := json.MarshalIndent(root, "", "\t") //formated output
	//bytes, _ := json.Marshal(root)
	fmt.Println(string(bytes))
	

	Load("Config.json")
	Save(bytes)


}
