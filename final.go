package main
import(
	"context"
	"fmt"
	"os/exec"
	"html/template"
	"net/http"
	"runtime"
	"strings"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)
type listdata struct {
	Imgid          string
	Imgname        string
	Command        string
	Status         string
	Container_Name string
}

func Listt() []listdata {
	context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All:true})
	if err != nil {
		panic(err)
	}

	list := make([]listdata, 0)
	for _, container := range containers {
		container_name := container.Names
		cname := fmt.Sprint(container_name)
		temp := listdata{container.ID[:10], container.Image, container.Command, container.Status, cname[2 : len(cname)-1]}
		list = append(list, temp)
	}
	return list
}
func main(){
	
	http.HandleFunc("/",index_page)
	fmt.Println("Starting server for testing HTTP POST...\n")
	http.ListenAndServe(":7070", nil)
	
}
func index_page(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("sample.html"))
	cmd:="sudo"
	if runtime.GOOS=="ubuntu"{
	fmt.Println("Cannot Execute this command\n")
	}else{ 
	commandtorun := r.FormValue("commandtorun")
	imagename := r.FormValue("imagename")
	conname := r.FormValue("conname")
	if imagename != "" {
      	commands := strings.Split(commandtorun, " ")
     	if len(commands)==1{
       		output, err := exec.Command(cmd ,"docker","run",  "-d", "--name", conname, imagename, commandtorun).CombinedOutput()
                if err != nil {
  			fmt.Printf("%s",err)
  		}else{
			fmt.Println("\nCommand Succesfully Executed: ")
			fmt.Printf(string(output))}
	}else if len(commands)==2{
       		output, err := exec.Command(cmd ,"docker","run",  "-d", "--name", conname, imagename, commands[0], commands[1]).CombinedOutput()
                if err != nil {
  			fmt.Printf("%s",err)
  		}else{
			fmt.Println("\nCommand Succesfully Executed: ")
			fmt.Printf(string(output))}
	}else if len(commands)==3{
       	output, err := exec.Command(cmd ,"docker","run",  "-d", "--name", conname, imagename, commands[0], commands[1], commands[2]).CombinedOutput()
                if err != nil {
  			fmt.Printf("%s",err)
  		}else{
			
			fmt.Println("\nCommand Succesfully Executed: ")
			fmt.Printf(string(output))}
	}else{
        	output, err := exec.Command(cmd, "docker","run",  "-d", "--name", conname, imagename).CombinedOutput()
        	if err != nil {
  			fmt.Printf("%s",err)
  		}
  		fmt.Printf("Command Successfully Executed\n")
		fmt.Printf(string(output))}
	   }
	
       }
       List := Listt()
       if err := t.ExecuteTemplate(w, "sample.html", List); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	constopname := r.FormValue("constopname")
   	if constopname != "" {
		output, err := exec.Command(cmd , "docker", "stop", constopname).CombinedOutput()
                if err != nil {
  			fmt.Printf("%s",err)
  		}else{
			fmt.Println("\nContainer Succesfully Stopped: ")
			fmt.Printf(string(output))}
	}
	constartname := r.FormValue("constartname")
   	if constartname != "" {
		output, err := exec.Command(cmd , "docker", "start", constartname).CombinedOutput()
                if err != nil {
  			fmt.Printf("%s",err)
  		}else{
			fmt.Println("\nContainer Succesfully Started: ")
			fmt.Printf(string(output))}
	}

}

