package cli

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/albert-wang/rawr-website-go/routes"
	"github.com/albert-wang/tracederror"
	"github.com/atotto/clipboard"

	"github.com/mitchellh/goamz/s3"
)

func importPost(args []string, context *routes.Context) error {
	//Args in the format
	// <n>.mkd hero title category time
	return nil
}

func importImage(args []string, context *routes.Context) error {
	//Args in the format
	// gallery <image>
	if len(args) < 2 {
		return fmt.Errorf("usage: <gallery> <image>")
	}

	args[0] = strings.TrimSpace(args[0])
	log.Print(args)

	// Step 1: create a sha-256 hash of the file contents.
	f, err := ioutil.ReadFile(args[1])
	if err != nil {
		return tracederror.New(err)
	}

	sum := sha256.Sum256(f)
	key := base64.URLEncoding.EncodeToString(sum[:])

	ext := filepath.Ext(args[1])
	m := mime.TypeByExtension(ext)

	// Step 2: Upload the original file to orig-<sha256>.jpg
	log.Printf("Uploading %s to %s/%s", args[1], args[0], key)
	err = context.Bucket.Put(fmt.Sprintf("%s/orig-%s", args[0], key), f, m, s3.PublicRead)
	if err != nil {
		return tracederror.New(err)
	}

	convbinary := "convert"
	if runtime.GOOS == "windows" {
		convbinary = "imgconvert"
	}

	// Step 3: Generate thumbnails/hero by cutting it through the center.
	os.Mkdir("temp", 0777)
	target := fmt.Sprintf("temp/%s%s", key, ext)
	cmd := exec.Command(convbinary, args[1],
		"-gravity", "center",
		"-resize", "25%",
		target)

	_, err = cmd.CombinedOutput()
	if err != nil {
		return tracederror.New(err)
	}

	f, err = ioutil.ReadFile(target)
	if err != nil {
		return tracederror.New(err)
	}

	log.Printf("Uploading thumbnail...")
	err = context.Bucket.Put(fmt.Sprintf("%s/thumb-%s", args[0], key), f, m, s3.PublicRead)
	if err != nil {
		return tracederror.New(err)
	}

	cmd = exec.Command(convbinary, args[1],
		"-gravity", "center",
		"-crop", "1200x400+0+0",
		target)

	_, err = cmd.CombinedOutput()
	if err != nil {
		return tracederror.New(err)
	}

	f, err = ioutil.ReadFile(target)
	if err != nil {
		return tracederror.New(err)
	}

	log.Printf("Uploading hero...")
	err = context.Bucket.Put(fmt.Sprintf("%s/hero-%s", args[0], key), f, m, s3.PublicRead)
	if err != nil {
		return tracederror.New(err)
	}

	os.Remove(target)

	log.Printf("Cleaning cache")
	clearCache(args, context)

	log.Printf("Saving Original Image URL to Clipboard...")

	url := context.Bucket.URL(fmt.Sprintf("/%s/orig-%s", args[0], key))
	clipboard.WriteAll(url)

	return nil
}

func clearCache(args []string, context *routes.Context) error {
	log.Printf("Cleaning cache")
	if context.Pool != nil {
		conn := context.Pool.Get()
		defer conn.Close()

		conn.Do("DEL", fmt.Sprintf("galleryimages.%s", args[0]))
	}
	return nil
}
