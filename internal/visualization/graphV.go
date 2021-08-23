package visualization

import (
	"Compiler_Objc/internal/parser"
	"bytes"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"log"
	"strings"
)

var arrDeepFunc = make(map[string][]string, 1)
func ParseGraph(globalFunc, allFuncArr map[int]parser.InfoType) {
	// clear Scope (func4.if.for.func4) -> (func4)
	for key, k := range allFuncArr {
		if strings.Contains(k.Scope, ".") {
			words := strings.Split(k.Scope, ".")
			allFuncArr[key] = parser.InfoType{Name: k.Name, DataType: k.DataType, Scope: words[0]}
		}
	}
	sortArrDeepFunc(globalFunc, allFuncArr)

	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		err := g.Close()
		if err != nil {
			return
		}
	}()

	var arrFunc = make(map[string]*cgraph.Node)

	// подставляем вершины что откуда
	for _, t := range globalFunc {
		arrFunc[t.Name], err = graph.CreateNode(t.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	for key, _ := range arrDeepFunc {
		for _, x := range arrDeepFunc[key] {
			if arrFunc[x] == nil {
				continue
			}
			_, err = graph.CreateEdge("e", arrFunc[key], arrFunc[x])
			if err != nil {
				log.Fatal(err)
			}
		}
	}


	// create your graph
	// 1. write encoded PNG data to buffer
	var buf bytes.Buffer
	if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
		log.Fatal(err)
	}

	// 2. get as image.Image instance
	_, err = g.RenderImage(graph)
	if err != nil {
		log.Fatal(err)
	}

	// 3. write to file directly
	if err := g.RenderFilename(graph, graphviz.PNG, "./result/graph.png"); err != nil {
		log.Fatal(err)
	}
}

func sortArrDeepFunc(global, arr map[int]parser.InfoType)  {
	for gk, g := range global {
		for ak, a := range arr {
			if global[gk].Name == arr[ak].Scope {
				arrDeepFunc[g.Name] = append(arrDeepFunc[g.Name], a.Name)
			}
		}
	}
}
