package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/albert-wang/tracederror"
	"github.com/russross/blackfriday"
)

type MarkdownRenderer struct {
	*blackfriday.Html

	Marker int
}

//Options
var flags = blackfriday.HTML_USE_XHTML | blackfriday.HTML_USE_SMARTYPANTS | blackfriday.HTML_SMARTYPANTS_FRACTIONS | blackfriday.HTML_SMARTYPANTS_DASHES | blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
var extensions = 0 |
	blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
	blackfriday.EXTENSION_TABLES |
	blackfriday.EXTENSION_FENCED_CODE |
	blackfriday.EXTENSION_AUTOLINK |
	blackfriday.EXTENSION_STRIKETHROUGH |
	blackfriday.EXTENSION_SPACE_HEADERS |
	blackfriday.EXTENSION_HEADER_IDS |
	blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
	blackfriday.EXTENSION_DEFINITION_LISTS

func escapeSingleChar(char byte) (string, bool) {
	if char == '"' {
		return "&quot;", true
	}
	if char == '&' {
		return "&amp;", true
	}
	if char == '<' {
		return "&lt;", true
	}
	if char == '>' {
		return "&gt;", true
	}
	return "", false
}

func attrEscape(out *bytes.Buffer, src []byte) {
	org := 0
	for i, ch := range src {
		if entity, ok := escapeSingleChar(ch); ok {
			if i > org {
				//copy all the normal characters since the last escape
				out.Write(src[org:i])
			}
			org = i + 1
			out.WriteString(entity)
		}
	}
	if org < len(src) {
		out.Write(src[org:])
	}
}

//Overwrite img to add a <p> caption below it, encase the image in a div for positioning and sizing purposes.
func (m *MarkdownRenderer) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	out.WriteString("<div class=\"img-container ")
	attrEscape(out, alt)
	out.WriteByte('"')
	out.WriteByte('>')

	out.WriteString("<div class=\"img\">")
	m.Html.Image(out, link, title, alt)
	out.WriteString("<div>")
	out.WriteString("<p class=\"caption\">")
	attrEscape(out, title)
	out.WriteString("</p></div></div></div>")
}

func templateBlackfriday(mkd string) template.HTML {
	//Set up the HTML renderer
	renderer := blackfriday.HtmlRenderer(flags, "", "").(*blackfriday.Html)
	wrappedRenderer := &MarkdownRenderer{
		Html: renderer,
	}

	res := blackfriday.MarkdownOptions([]byte(mkd), wrappedRenderer, blackfriday.Options{
		Extensions: extensions,
	})

	return template.HTML(res)
}

func s3img(path string) string {
	return "//s3.amazonaws.com/img.rawrrawr.com/" + path
}

func add(a, b int) int {
	return a + b
}

func slug(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	fields := strings.Fields(s)

	res := strings.Join(fields, "-")
	return url.QueryEscape(res)
}

func timef(f string) string {
	return time.Now().Format(f)
}

func LoadTemplates(str ...string) (*template.Template, error) {
	pathed := []string{}
	for _, v := range str {
		pathed = append(pathed, fmt.Sprintf("assets/templates/%s", v))
	}

	res := template.New(str[0])
	// Load default functions.
	mapping := template.FuncMap{
		"blackfriday": templateBlackfriday,
		"s3img":       s3img,
		"add":         add,
		"slug":        slug,
		"timef":       timef,
	}

	res = res.Funcs(mapping)

	res, err := res.ParseFiles(pathed...)
	if err != nil {
		log.Print(tracederror.New(err))
		return nil, tracederror.New(err)
	}

	return res, tracederror.New(err)
}

func RenderTemplateWithData(w io.Writer, r *http.Request, tpl *template.Template, data interface{}) error {
	err := tpl.Execute(w, data)
	if err != nil {
		return MakeHttpError(err, http.StatusInternalServerError, r)
	}

	return nil
}
