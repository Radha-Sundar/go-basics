package godep_app // main.go
import (
	"fmt"
	"github.com/Radha-Sundar/godep-app/pkg/greeter"
)

func main() {
	message := greeter.GetGreeting()
	fmt.Println(message)
}
