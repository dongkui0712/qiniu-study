package third
import (
	"fmt"
	"os"
)

func usage(){
	fmt.Printf("%s %s %s\n",os.Args[0],"filename","newFilename")
}
func CpFile(n int) {
	if len(os.Args) !=3{
		usage()//Args hold the command-line arguments, starting with the program name.
		return
	}
	//open file
	filename_from := os.Args[1]
	ff,err := os.Open(filename_from)
	if err != nil{
		panic(err)
	}
	defer ff.Close()

	//create new file
	filename_to := os.Args[2]
	ft,err := os.Create(filename_to)
	if err != nil{
		panic(err)
	}
	defer ft.Close()

	//read from file_from,write to file_to
	var buff = make([]byte, n)
	for{
		n,err := ff.Read(buff)
		if err != nil {
			panic(err)
		}
		if n == 0{
			break
		}

		if _,err := ft.Write(buff[:n]);err != nil{
			panic(err)
		}
	}




}
