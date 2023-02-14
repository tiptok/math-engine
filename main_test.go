package main

import (
	"fmt"
	"github.com/dengsgo/math-engine/engine"
	"testing"
)

func TestExecA(t *testing.T) {
	exp := "1+2"
	exec(exp)
}

func TestExecB(t *testing.T) {
	exp := "1+2-4"
	exec(exp)
}

func TestExecC(t *testing.T) {
	exp := "1+2-4*3-8"
	exec(exp)
}

func TestExecD(t *testing.T) {
	exp := "1+2-(4*3-8)"
	exec(exp)
}

func TestExecE(t *testing.T) {
	exp := "1+2-(4*3+(1-8))"
	exec(exp)
}

func TestExecF(t *testing.T) {
	exp := "1+(2-(4*3+(1-8)))"
	exec(exp)
}

func TestExecG(t *testing.T) {
	exp := "((1-2)*(3-8))*((((9+2222))))"
	exec(exp)
}

func TestExecH(t *testing.T) {
	exp := "0.8888-0.1 * 444         -0.2"
	exec(exp)
}

func TestExecI(t *testing.T) {
	exp := "0.8888-0.1 * (444         -0.2)"
	exec(exp)
}

func TestExecJ(t *testing.T) {
	exp := "1_234_567*2-3"
	exec(exp)
}

func TestExecK(t *testing.T) {
	exp := "2.3e4*4/3"
	exec(exp)
}

func TestExecL(t *testing.T) {
	exp := "-1+9-88"
	exec(exp)
}

func TestExecM(t *testing.T) {
	exp := "-1+9-88+(88)"
	exec(exp)
}

func TestExecN(t *testing.T) {
	exp := "-1+9-88+(-88)*666-1"
	exec(exp)
}

func TestExecO(t *testing.T) {
	exp := "-(1)+(3)-(-3)*7-((-3))"
	exec(exp)
}

func TestExecP(t *testing.T) {
	exp := "-(-9+3)"
	exec(exp)
}

func TestExecQ(t *testing.T) {
	exp := "2e-3*2+2e2+1"
	exec(exp)
}

func TestExecR(t *testing.T) {
	exp := "3.8 - 56 / (1-1) - 4"
	exec(exp)
}

func TestExecS(t *testing.T) {
	exp := "noerr(3.8 - 56 / (1-1) - 4)"
	exec(exp)
}

func TestFunCaller(t *testing.T) {
	funs := []struct {
		Name string
		Argc int
		Fun  func(expr ...engine.ExprAST) float64
		Exp  string
		R    float64
	}{
		//{
		//	"double",
		//	1,
		//	func(expr ...engine.ExprAST) float64 {
		//		return engine.ExprASTResult(expr[0]) * 2
		//	},
		//	"double(6)",
		//	12,
		//},
		{
			"sum",
			-1,
			nil,
			"sum(if(100+10,table.a,20))",
			10,
		},
		{
			"sum",
			-1,
			nil,
			"sum(if(100<10,table.a,20))",
			10,
		},
		{
			"sum",
			-1,
			nil,
			"sum(if(100<10,table.a,20+30))",
			10,
		},
		{
			"sum",
			-1,
			nil,
			"sum(if(table.a<table.b,table.a,20+30))",
			10,
		},
		{
			"sum",
			-1,
			nil,
			"sum(if(table.a<=table.b,table.a,20+30))",
			10,
		},
	}
	for _, f := range funs {
		if f.Fun != nil {
			_ = engine.RegFunction(f.Name, f.Argc, f.Fun)
		}
		r, err := Parse(f.Exp)
		if err != nil {

		}
		if r != 0 {

		}
	}
}

func TestFunCaller2(t *testing.T) {
	funs := []struct {
		Name string
		Exp  []string
	}{
		{
			"sum",
			[]string{"sum(table.a)"},
		},
		{
			"sumif",
			[]string{"sumif(table.month,10,table.count)"},
		},
		{
			"if",
			[]string{"if(table.month>10,table.count1,table.count2)"},
		},
		{
			"and",
			[]string{"and(table.year=2011,table.month=6)"},
		},
		{
			"or",
			[]string{"or(table.year=2011,table.year=2012)"},
		},
		{
			"month",
			[]string{"month(\"1991-1-1\")"},
		},
		{
			"year",
			[]string{"year(\"1991-1-1\")"},
		},
		{
			"round",
			[]string{
				"round(1.56)",
				"round(table.a)",
			},
		},
		{
			"rounddown",
			[]string{
				"rounddown(1.56)",
				"rounddown(table.a)",
			},
		},
		{
			"roundup",
			[]string{
				"roundup(1.56)",
				"roundup(table.a)",
			},
		},
		{
			"count",
			[]string{
				"count(1.56)",
				"count(table.a)",
			},
		},
		{
			"&",
			[]string{
				"table.a&table.b",
			},
		},
	}
	for _, f := range funs {
		for _, exp := range f.Exp {
			r, err := Parse(exp)
			if err != nil {
				t.Error(err)
			}
			if r != 0 {

			}
		}
	}
}

func Parse(s string) (r float64, err error) {
	toks, err := engine.Parse(s)
	if err != nil {
		return 0, err
	}
	ast := engine.NewAST(toks, s)
	if ast.Err != nil {
		return 0, ast.Err
	}
	ar := ast.ParseExpression()
	if ast.Err != nil {
		return 0, ast.Err
	}
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	if ar != nil {
		fmt.Printf("ExprAST: %+v\n", ar)
	}
	return 0, err
}

// https://www.yoytang.com/math-expression-engine.html
