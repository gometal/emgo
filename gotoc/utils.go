package gotoc

import (
	"bytes"
	"code.google.com/p/go.tools/go/types"
	"fmt"
	"go/ast"
	"io"
	"os"
	"strconv"
	"strings"
)

func notImplemented(n ast.Node) {
	fmt.Fprintf(os.Stderr, "not implemented: %s <%T>", n, n)
	os.Exit(1)
}

func upath(path string) string {
	return strings.Replace(path, "/", "_", -1)
}

func tmpname(w *bytes.Buffer) string {
	return "__" + strconv.Itoa(w.Len())
}

func uniqueId(o types.Object) string {
	return strconv.Itoa(int(o.Pos()))
}

func write(s string, ws ...io.Writer) error {
	for _, w := range ws {
		if _, err := io.WriteString(w, s); err != nil {
			return err
		}
	}
	return nil
}