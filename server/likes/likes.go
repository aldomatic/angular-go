package likes

import(
	"fmt"
)


func init(){
	fmt.Println("Likes package ready")
}

func HowManyLikes() string{
	return "Hello from likes package"
}