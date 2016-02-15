package cli

import (
	"log"

	"github.com/albert-wang/rawr-website-go/routes"
)

func Dispatch(args []string, context *routes.Context) {
	var err error

	switch args[0] {
	case "import-post":
		err = importPost(args[1:], context)

	case "import-image":
		err = importImage(args[1:], context)

	case "clear-cache":
		err = clearCache(args[1:], context)

	default:
		log.Fatal("Unknown commandline option ", args[0])
	}

	if err != nil {
		log.Fatal(err)
	}
}
