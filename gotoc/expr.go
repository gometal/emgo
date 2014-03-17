package gotoc

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"strconv"

	"code.google.com/p/go.tools/go/exact"
	"code.google.com/p/go.tools/go/types"
)

func writeInt(w *bytes.Buffer, ev exact.Value, k types.BasicKind) {
	if k == types.Uintptr {
		u, _ := exact.Uint64Val(ev)
		w.WriteString("0x")
		w.WriteString(strconv.FormatUint(u, 16))
		return
	}

	w.WriteString(ev.String())
	switch k {
	case types.Int32:
		w.WriteByte('L')

	case types.Uint32:
		w.WriteString("UL")

	case types.Int64:
		w.WriteString("LL")

	case types.Uint64:
		w.WriteString("ULL")
	}
}

func writeFloat(w *bytes.Buffer, ev exact.Value, k types.BasicKind) {
	v, _ := exact.Int64Val(exact.Num(ev))
	w.WriteString(strconv.FormatInt(v, 10))
	v, _ = exact.Int64Val(exact.Denom(ev))
	if v != 1 {
		w.WriteByte('/')
		w.WriteString(strconv.FormatInt(v, 10))
	}
	w.WriteByte('.')
	if k == types.Float32 {
		w.WriteByte('F')
	}
}

func (cdd *CDD) Value(w *bytes.Buffer, ev exact.Value, t types.Type) {
	k := t.Underlying().(*types.Basic).Kind()

	// TODO: use t instead ev.Kind() in following switch
	switch ev.Kind() {
	case exact.Int:
		writeInt(w, ev, k)

	case exact.Float:
		writeFloat(w, ev, k)

	case exact.Complex:
		switch k {
		case types.Complex64:
			k = types.Float32
		case types.Complex128:
			k = types.Float64
		default:
			k = types.UntypedFloat
		}
		writeFloat(w, exact.Real(ev), k)
		im := exact.Imag(ev)
		if exact.Sign(im) != -1 {
			w.WriteByte('+')
		}
		writeFloat(w, im, k)
		w.WriteByte('i')

	case exact.String:
		w.WriteString("__EGSTR(")
		w.WriteString(ev.String())
		w.WriteByte(')')

	default:
		w.WriteString(ev.String())
	}
}

func (cdd *CDD) Name(w *bytes.Buffer, obj types.Object, direct bool) {
	if obj == nil {
		w.WriteByte('_')
		return
	}
	switch o := obj.(type) {
	case *types.PkgName:
		// Imported package name in SelectorExpr: pkgname.Name
		w.WriteString(upath(o.Pkg().Path()))
		return

	case *types.Func:
		s := o.Type().(*types.Signature)
		if r := s.Recv(); r != nil {
			t := r.Type()
			if p, ok := t.(*types.Pointer); ok {
				t = p.Elem()
				direct = false
			}
			cdd.Type(w, t)
			w.WriteByte('_')
			w.WriteString(o.Name())
			if !cdd.gtc.isLocal(t.(*types.Named).Obj()) {
				cdd.addObject(o, direct)
			}
			return
		}
	}

	if p := obj.Pkg(); p != nil && !cdd.gtc.isLocal(obj) {
		cdd.addObject(obj, direct)
		w.WriteString(upath(obj.Pkg().Path()))
		w.WriteByte('_')
	}
	name := obj.Name()
	switch name {
	case "_":
		w.WriteString("__unused")
		w.WriteString(cdd.gtc.uniqueId())
		return

	case "init":
		name = cdd.gtc.uniqueId() + name
	}
	w.WriteString(name)
}

func (cdd *CDD) NameStr(o types.Object, direct bool) string {
	buf := new(bytes.Buffer)
	cdd.Name(buf, o, direct)
	return buf.String()
}

func (cdd *CDD) SelectorExpr(w *bytes.Buffer, e *ast.SelectorExpr) (recv ast.Expr) {
	sel := cdd.gtc.ti.Selections[e]
	switch sel.Kind() {
	case types.FieldVal:
		cdd.Expr(w, e.X, nil)
		if sel.Indirect() {
			w.WriteString("->")
		} else {
			w.WriteByte('.')
		}
		w.WriteString(e.Sel.Name)

	case types.MethodVal:
		cdd.Name(w, sel.Obj(), true)
		if _, ok := sel.Recv().(*types.Pointer); !ok {
			// Method has non-pointer receiver so there is guaranteed
			// that e.X isn't a pointer.
			recv = e.X
		}
		// Method has pointer receiver.
		if sel.Indirect() {
			// e.X is pointer.
			recv = e.X
		} else {
			// e.X isn't pointer but pointer receiver is need.
			recv = &ast.UnaryExpr{Op: token.AND, X: e.X}
		}

	case types.PackageObj:
		cdd.Name(w, sel.Obj(), true)

	default: // types.MethodExpr
		notImplemented(e)
	}
	return
}

func (cdd *CDD) Expr(w *bytes.Buffer, expr ast.Expr, nilT types.Type) {
	cdd.Complexity++

	if t := cdd.gtc.ti.Types[expr]; t.Value != nil {
		// Constant expression
		cdd.Value(w, t.Value, t.Type)
		return
	}

	switch e := expr.(type) {
	case *ast.BinaryExpr:
		op := e.Op.String()
		ltyp := cdd.gtc.ti.Types[e.X].Type
		rtyp := cdd.gtc.ti.Types[e.Y].Type

		lhs := cdd.ExprStr(e.X, ltyp)
		rhs := cdd.ExprStr(e.Y, rtyp)
		if op == "==" || op == "!=" {
			eq(w, lhs, op, rhs, ltyp, rtyp)
			break
		}
		// BUG: string type
		hi := (op == "*" || op == "/" || op == "%")
		if !hi {
			w.WriteByte('(')
		}
		if op == "&^" {
			op = "&~"
		}
		w.WriteString(lhs + op + rhs)

		if !hi {
			w.WriteByte(')')
		}

	case *ast.UnaryExpr:
		op := e.Op.String()
		if op == "^" {
			op = "~"
		}
		w.WriteString(op)
		cdd.Expr(w, e.X, nil)

	case *ast.CallExpr:
		var recv ast.Expr

		switch t := cdd.gtc.ti.Types[e.Fun].Type.(type) {
		case *types.Signature:
			switch f := e.Fun.(type) {
			case *ast.SelectorExpr:
				recv = cdd.SelectorExpr(w, f)

			case *ast.Ident:
				switch o := cdd.object(f).(type) {
				case *types.Builtin:
					cdd.builtin(w, o, e.Args)
					return

				default:
					cdd.Name(w, o, true)
				}

			default:
				cdd.Expr(w, f, nil)
			}

			w.WriteByte('(')
			if recv != nil {
				cdd.Expr(w, recv, nil)
				if len(e.Args) > 0 {
					w.WriteString(", ")
				}
			}
			tup := t.Params()
			for i, a := range e.Args {
				if i != 0 {
					w.WriteString(", ")
				}
				cdd.Expr(w, a, tup.At(i).Type())
			}
			w.WriteByte(')')

		default:
			arg := e.Args[0]
			switch typ := cdd.gtc.ti.Types[e.Fun].Type.(type) {
			case *types.Slice:
				switch cdd.gtc.ti.Types[arg].Type.(type) {
				case *types.Basic: // string
					w.WriteString("__NEWSTR(")
					cdd.Expr(w, arg, typ)
					w.WriteByte(')')

				default: // slice
					w.WriteByte('(')
					cdd.Expr(w, arg, typ)
					w.WriteByte(')')
				}

			default:
				w.WriteString("((")
				dim, _ := cdd.Type(w, typ)
				w.WriteString(dimFuncPtr("", dim))
				w.WriteString(")(")
				cdd.Expr(w, arg, typ)
				w.WriteString("))")
			}
		}

	case *ast.Ident:
		if e.Name == "nil" {
			cdd.Nil(w, nilT)
		} else {
			cdd.Name(w, cdd.object(e), true)
		}

	case *ast.IndexExpr:
		typ := cdd.gtc.ti.Types[e.X].Type

		pt, isPtr := typ.(*types.Pointer)
		if isPtr {
			w.WriteString("(*")
			typ = pt.Elem()
		}

		var indT types.Type

		switch t := typ.(type) {
		case *types.Basic: // string
			cdd.Expr(w, e.X, nil)
			w.WriteString(".str")

		case *types.Slice:
			w.WriteString("((")
			dim, _ := cdd.Type(w, t.Elem())
			dim = append([]string{"*"}, dim...)
			w.WriteString(dimFuncPtr("", dim))
			w.WriteByte(')')
			cdd.Expr(w, e.X, nil)
			w.WriteString(".arr)")

		case *types.Array:
			cdd.Expr(w, e.X, nil)

		case *types.Map:
			indT = t.Key()
			notImplemented(e)

		default:
			panic(t)
		}

		if isPtr {
			w.WriteByte(')')
		}

		w.WriteByte('[')
		cdd.Expr(w, e.Index, indT)
		w.WriteByte(']')

	case *ast.KeyValueExpr:
		w.WriteByte('.')
		kt := cdd.gtc.ti.Types[e.Key].Type
		if i, ok := e.Key.(*ast.Ident); ok && kt == nil {
			// e.Key is field name
			w.WriteString(i.Name)
		} else {
			cdd.Expr(w, e.Key, kt)
		}
		w.WriteString(" = ")
		cdd.Expr(w, e.Value, nilT)

	case *ast.ParenExpr:
		w.WriteByte('(')
		cdd.Expr(w, e.X, nilT)
		w.WriteByte(')')

	case *ast.SelectorExpr:
		cdd.SelectorExpr(w, e)

	case *ast.SliceExpr:
		cdd.SliceExpr(w, e)

	case *ast.StarExpr:
		w.WriteByte('*')
		cdd.Expr(w, e.X, nil)

	case *ast.TypeAssertExpr:
		notImplemented(e)

	case *ast.CompositeLit:
		typ := cdd.gtc.ti.Types[e].Type

		switch t := typ.(type) {
		case *types.Array:
			w.WriteByte('{')
			nilT = t.Elem()

		case *types.Slice:
			w.WriteString("(__slice){(")
			dim, _ := cdd.Type(w, t.Elem())
			dim = append([]string{"[]"}, dim...)
			w.WriteString(dimFuncPtr("", dim))
			w.WriteString("){")
			nilT = t.Elem()

		default:
			w.WriteByte('(')
			cdd.Type(w, t)
			w.WriteString("){")
			nilT = nil
		}

		for i, el := range e.Elts {
			if i > 0 {
				w.WriteString(", ")
			}
			cdd.Expr(w, el, nilT)
		}

		switch typ.(type) {
		case *types.Slice:
			w.WriteByte('}')
			plen := ", " + strconv.Itoa(len(e.Elts))
			w.WriteString(plen)
			w.WriteString(plen)
			w.WriteByte('}')

		default:
			w.WriteByte('}')
		}

	default:
		fmt.Fprintf(w, "!%v<%T>!", e, e)
	}
}

func (cdd *CDD) SliceExpr(w *bytes.Buffer, e *ast.SliceExpr) {
	sx := cdd.ExprStr(e.X, nil)

	typ := cdd.gtc.ti.Types[e.X].Type
	pt, isPtr := typ.(*types.Pointer)
	if isPtr {
		typ = pt.Elem()
		sx = "(*" + sx + ")"
	}

	switch t := typ.(type) {
	case *types.Slice:
		if e.Low == nil && e.High == nil && e.Max == nil {
			w.WriteString(sx)
			break
		}

		if e.Low != nil {
			switch {
			case e.High == nil && e.Max == nil:
				w.WriteString("__SLICEL(")

			case e.High != nil && e.Max == nil:
				w.WriteString("__SLICELH(")

			case e.High == nil && e.Max != nil:
				w.WriteString("__SLICEM(")

			default:
				w.WriteString("__SLICELHM(")
			}
			w.WriteString(sx)
			w.WriteString(", ")
			dim, _ := cdd.Type(w, t.Elem())
			dim = append([]string{"*"}, dim...)
			w.WriteString(dimFuncPtr("", dim))
			w.WriteString(", ")
			cdd.Expr(w, e.Low, nil)
		} else {
			switch {
			case e.High != nil && e.Max == nil:
				w.WriteString("__SLICEH(")

			case e.High == nil && e.Max != nil:
				w.WriteString("__SLICEM(")

			default:
				w.WriteString("__SLICEHM(")
			}
			w.WriteString(sx)
		}

		if e.High != nil {
			w.WriteString(", ")
			cdd.Expr(w, e.High, nil)
		}
		if e.Max != nil {
			w.WriteString(", ")
			cdd.Expr(w, e.Max, nil)
		}

		w.WriteByte(')')

	case *types.Array:
		if e.Low != nil {
			switch {
			case e.High == nil && e.Max == nil:
				w.WriteString("__ASLICEL(")

			case e.High != nil && e.Max == nil:
				w.WriteString("__ASLICELH(")

			case e.High == nil && e.Max != nil:
				w.WriteString("__ASLICEM(")

			default:
				w.WriteString("__ASLICELHM(")
			}
			w.WriteString(sx)
			w.WriteString(", ")
			cdd.Expr(w, e.Low, nil)
		} else {
			switch {
			case e.High == nil && e.Max == nil:
				w.WriteString("__ASLICE(")

			case e.High != nil && e.Max == nil:
				w.WriteString("__ASLICEH(")

			case e.High == nil && e.Max != nil:
				w.WriteString("__ASLICEM(")

			default:
				w.WriteString("__ASLICEHM(")
			}
			w.WriteString(sx)
		}

		if e.High != nil {
			w.WriteString(", ")
			cdd.Expr(w, e.High, nil)
		}
		if e.Max != nil {
			w.WriteString(", ")
			cdd.Expr(w, e.Max, nil)
		}

		w.WriteByte(')')

	case *types.Basic: // string
		if e.Low == nil && e.High == nil {
			w.WriteString(sx)
			break
		}
		switch {
		case e.Low == nil:
			w.WriteString("__SSLICEH(")

		case e.High == nil:
			w.WriteString("__SSLICEL(")

		default:
			w.WriteString("__SSLICELH(")
		}

		w.WriteString(sx)

		if e.Low != nil {
			w.WriteString(", ")
			cdd.Expr(w, e.Low, nil)
		}
		if e.High != nil {
			w.WriteString(", ")
			cdd.Expr(w, e.High, nil)
		}

		w.WriteByte(')')

	default:
		panic(e)
	}
}

func (cdd *CDD) builtin(w *bytes.Buffer, b *types.Builtin, args []ast.Expr) {
	etm := cdd.gtc.ti.Types

	switch name := b.Name(); name {
	case "len":
		switch t := etm[args[0]].Type.(type) {
		case *types.Slice, *types.Map, *types.Basic: // Basic == String
			w.WriteString("len(")

		case *types.Array:
			w.WriteString("sizeof(")

		default:
			panic(t)
		}

	case "copy":
		switch t := etm[args[1]].Type.(type) {
		case *types.Basic: // string
			w.WriteString("__STRCPY(")

		case *types.Slice:
			w.WriteString("__SLICPY(")
			dim, _ := cdd.Type(w, t.Elem())
			w.WriteString(dimFuncPtr("", dim))
			w.WriteString(", ")

		default:
			panic(t)
		}

	default:
		w.WriteString(name + "(")
	}

	for i, a := range args {
		if i != 0 {
			w.WriteString(", ")
		}
		cdd.Expr(w, a, nil)
	}
	w.WriteByte(')')

}

func (cdd *CDD) ExprStr(expr ast.Expr, nilT types.Type) string {
	buf := new(bytes.Buffer)
	cdd.Expr(buf, expr, nilT)
	return buf.String()
}

func (cdd *CDD) Nil(w *bytes.Buffer, t types.Type) {
	switch t.(type) {
	case *types.Slice:
		w.WriteString("__NILSLICE")

	case *types.Map:
		w.WriteString("__NILMAP")

	case *types.Pointer:
		w.WriteString("nil")

	default:
		w.WriteString("{0}")
	}
}

func eq(w *bytes.Buffer, lhs, op, rhs string, ltyp, rtyp types.Type) {
	typ := ltyp
	if typ == types.Typ[types.UntypedNil] {
		typ = rtyp
	}

	switch typ.(type) {
	case *types.Slice:
		if rtyp == types.Typ[types.UntypedNil] {
			lhs += ".arr"
			rhs = "nil"
		} else {
			lhs = "nil"
			rhs += ".arr"
		}
	}
	w.WriteString(lhs + " " + op + " " + rhs)
}
