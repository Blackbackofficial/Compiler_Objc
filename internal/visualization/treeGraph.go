package visualization

import (
	"Compiler_Objc/internal/parser"
	"bytes"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"log"
	"strconv"
)

func Graph(tree *parser.BaseObjCListener) {
	var queueGraph []*cgraph.Node
	var queue []*opts.TreeData
	var count int

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

	// Step: 0
	start, err := graph.CreateNode(tree.Root.Name)
	if err != nil {
		log.Fatal(err)
	}
	queue = append(queue, tree.Root)
	queueGraph = append(queueGraph, start)
	for len(queue) > 0 {
		queue = append(queue, queue[0].Children...)
		for _, v := range queue[0].Children {
			name, err := graph.CreateNode(strconv.Itoa(count))
			count++
			if err != nil {
				log.Fatal(err)
			}
			name.SetLabel(v.Name)
			queueGraph = append(queueGraph, name)
			e, err := graph.CreateEdge(strconv.Itoa(count+1), queueGraph[0], name)
			if err != nil {
				log.Fatal(err)
			}
			e.SetDecorate(false)
		}
		queueGraph = queueGraph[1:]
		//fmt.Println(queue[0].Name)
		if len(queue) != 0 {
			queue = queue[1:]
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
	if err := g.RenderFilename(graph, graphviz.PNG, "./result/test.png"); err != nil {
		log.Fatal(err)
	}
}
