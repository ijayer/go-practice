package main

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	// TmplTest1()
	// TmplTest2()
	// TmplTest3()
	// TmplTest4()
	EmbeddedTmpl()
}

// Test range and with
type Person struct {
	Username string
	Emails   []string
	Friends  []*Friend
	CreateAt time.Time
}

type Friend struct {
	Name string
}

var tmpl = `Hello {{.Username}}
{{range .Emails}}
	an email {{.}}
{{end}}
{{range .Friends}}
	my friend name is {{.Name}}
{{end}}
`

func TmplTest1() {
	f1 := Friend{Name: "Lie"}
	f2 := Friend{Name: "Shu"}

	t := template.New("FieldName example")
	t, _ = t.Parse(tmpl)

	p := Person{
		Username: "zhe",
		Emails:   []string{"1@1.com", "2@2.com"},
		Friends:  []*Friend{&f1, &f2},
	}
	t.Execute(os.Stdout, p)
}

// Test func that define by myself
type Para struct {
	X int
	Y int
}

var tmpl2 = `Result is {{sum .X .Y}}`

func TmplTest2() {
	funcMap := template.FuncMap{
		"sum": add,
	}

	t := template.New("Call func").Funcs(funcMap)
	t, _ = t.Parse(tmpl2)

	p := Para{
		X: 1,
		Y: 2,
	}

	t.Execute(os.Stdout, p)
}

func add(x, y int) int {
	return x + y
}

// Test range slices
type Slices struct {
	Names []string
	Out   string
}

var tmpl3 = `
{{range $i, $v := .Names}}
	the index is {{$i}} and the value is {{$v}}
	outside var is {{$.Out}}
{{end}}
`

func TmplTest3() {
	s := Slices{
		Names: []string{"1", "2", "3", "4", "5", "6", "7"},
		Out:   "ooooooo",
	}

	t := template.New("range test")
	t, _ = t.Parse(tmpl3)
	t.Execute(os.Stdout, s)
}

func TmplTest4() {
	tOk := template.New("first")
	template.Must(tOk.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed OK.")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("The second one parsed OK.")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }}"))
}

// Embedded template
func EmbeddedTmpl() {
	path := "/home/work/code/Go_Path/src/instance.golang.com/tmpl/tmpl"
	t, err := template.ParseFiles(path+"/header.tmpl", path+"/content.tmpl", path+"/footer.tmpl")
	if err != nil {
		logrus.Error(err)
	}
	println(t.Name())

	t.ExecuteTemplate(os.Stdout, "header", nil) // 渲染指定模板 header
	fmt.Println()

	t.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()

	t.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()

}
