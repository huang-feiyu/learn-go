package cyoa

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Paras   []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Chapter

func ReadJSONFile(filename string) (story Story) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&story); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	return
}

type handler struct {
	s Story
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(htmlTemplate))

	path := strings.TrimSpace(r.URL.Path[1:])
	if chapter, ok := h.s[path]; ok {
		tpl.Execute(w, chapter)
		return
	}

	tpl.Execute(w, h.s["intro"])
}

var htmlTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Choose Your own Adventure</title>
	</head>
	<body>
		<section class="page">
			<h1>{{.Title}}</h1>
				{{range .Paras}}
				<p>{{.}}</p>
				{{end}}
			<ul>
				{{range .Options}}
				<li><a href="/{{.Arc}}">{{.Text}}</a></li>
				{{end}}
			</ul>
		</section>
		<style>
			body {
				font-family: verdana, helvetica, arial;
			}
			h1 {
				text-align:center;
				position:relative;
			}
			.page {
				width: 80%;
				max-width: 500px;
				margin: auto;
				margin-top: 40px;
				margin-bottom: 40px;
				padding: 80px;
				background: #FFFCF6;
				border: 1px solid #eee;
				box-shadow: 0 10px 6px -6px #777;
			}
			ul {
				border-top: 1px dotted #ccc;
				padding: 10px 0 0 0;
				-webkit-padding-start: 0;
			}
			li {
				padding-top: 10px;
			}
			a,
			a:visited {
				text-decoration: none;
				color: #6295b5;
			}
			a:active,
			a:hover {
				color: #7792a2;
			}
			p {
				text-indent: 1em;
			}
		</style>
	</body>
</html>
`
