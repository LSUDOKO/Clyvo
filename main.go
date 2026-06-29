package main
import(
	"github.com/LSUDOKO/Clyvo/routes"
	"github.com/gin-gonic/gin"
	"github.com/LSUDOKO/Clyvo/controller"
	"github.com/LSUDOKO/Clyvo/middlewares"
	"github.com/LSUDOKO/Clyvo/databases"
	
	
)
func main(){
	port:=os.Getenv("PORT")
	if (port==""){
		port="8000"
		
	}
}