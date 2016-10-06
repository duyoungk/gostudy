package main

import (
	"bufio"
	_ "encoding/json"
	_ "flag"
	"fmt"
	_ "io/ioutil"
	_ "log"
	_ "net"
	_ "net/rpc"
	_ "os"
	"os/exec"
	_ "strconv"
	_ "strings"

	"github.com/mitchellh/go-ps"
)

type ServerInfo struct {
	Ip   string
	Port int
}

type ConfigObject struct {
	Server_list []ServerInfo
}

type Calc int
type Args struct {
	A, B int
}
type Reply struct {
	C int
}

func (c *Calc) Sum(args *Args, r *Reply) error {
	r.C = args.A + args.B
	return nil
}

type Commander int
type CmdArgs struct {
	CommandString string
	Params        string
}
type CmdReply struct {
	ReplyString string
}

func (c *Commander) Execute(args *CmdArgs, r *CmdReply) error {
	cmd := exec.Command(args.CommandString)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Start()
	ret := bufio.NewReader(out)
	line, _, err := ret.ReadLine()
	r.ReplyString = string(line)
	return nil
}

var mode string
var ip string
var port int
var command string

//func init() {
//	flag.StringVar(&mode, "mode", "client", "run mode [client|server]")
//	flag.StringVar(&ip, "ip", "", "")
//	flag.IntVar(&port, "port", 0, "")
//	flag.StringVar(&command, "cmd", "", "")
//	flag.Parse()

//	mode = strings.ToLower(mode)

//	if mode == "client" && ip == "" {
//		fmt.Println("missing ip")
//		os.Exit(1)
//	}
//	if mode == "client" && port == 0 {
//		fmt.Println("missing port")
//		os.Exit(1)
//	}
//	if mode == "server" && port == 0 {
//		fmt.Println("missing port")
//		os.Exit(1)
//	}
//}
func main() {

	//  json test
	// file, e := ioutil.ReadFile("./config.json")
	// if e != nil {
	//     fmt.Printf("File Open Error: %v\n", e)
	//     os.Exit(1)
	// }

	// var configObject ConfigObject
	// json.Unmarshal(file, &configObject)

	// for _, obj := range configObject.Server_list {
	//     fmt.Printf("ip : %s port : %d\n", obj.Ip, obj.Port)
	// }

	// rpc test
	//	switch mode {
	//	case "client":
	//		fmt.Println("client mode")
	//		addr := ip + ":" + strconv.Itoa(port)

	//		client, err := rpc.Dial("tcp", addr)
	//		if err != nil {
	//			log.Fatal("dialing: ", err)
	//			return
	//		}
	//		defer client.Close()
	//		fmt.Println("Connected!!")

	//		if command != "" {
	//			args := &CmdArgs{command, ""}
	//			reply := new(CmdReply)
	//			err = client.Call("Commander.Execute", args, reply)
	//			if err != nil {
	//				fmt.Println("cannot execute command")
	//				return
	//			}
	//			fmt.Println(reply.ReplyString)
	//		}

	//		//		args := &Args{1, 2}
	//		//		reply := new(Reply)
	//		//		err = client.Call("Calc.Sum", args, reply)
	//		//		if err != nil {
	//		//			log.Fatal("call error:", err)
	//		//			return
	//		//		}
	//		//		fmt.Printf("%d + %d = %d\n", args.A, args.B, reply.C)

	//		//		args2 := &CmdArgs{"dir", ""}
	//		//		reply2 := new(CmdReply)
	//		//		err = client.Call("Commander.Execute", args2, reply2)
	//		//		if err != nil {
	//		//			log.Fatal("call error:", err)
	//		//			return
	//		//		}
	//		//		fmt.Printf("result %s\n", reply2.ReplyString)

	//	case "server":
	//		fmt.Println("server mode")
	//		calc := new(Calc)
	//		commander := new(Commander)
	//		rpc.Register(calc)
	//		rpc.Register(commander)
	//		l, e := net.Listen("tcp", ":"+strconv.Itoa(port))
	//		if e != nil {
	//			log.Fatal("listen error:", e)
	//		}
	//		defer l.Close()

	//		for {
	//			fmt.Println("Waiting...")
	//			conn, err := l.Accept()
	//			if err != nil {
	//				fmt.Println("Accept error:", err)
	//				continue
	//			}
	//			defer conn.Close()
	//			fmt.Println("Accepted!!")
	//			go rpc.ServeConn(conn)
	//		}

	//	default:
	//		fmt.Println("End")
	//	}

	p, err := ps.Processes()
	if err != nil {
		fmt.Println(err.Error)
		return
	}

	for _, p1 := range p {
		fmt.Println(p1.Executable())
	}

}
