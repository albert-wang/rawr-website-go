package cli

import (
	"log"

	"github.com/albert-wang/rawr-website-go/routes"
)

func Dispatch(args []string, context *routes.Context) {
	switch args[0] {
	case "import-post":
		importPost(args, context)

	default:
		log.Fatal("Unknown commandline option ", args[0])
	}
}
