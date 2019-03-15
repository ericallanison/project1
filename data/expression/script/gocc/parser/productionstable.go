// Code generated by gocc; DO NOT EDIT.

package parser

import "github.com/project-flogo/core/data/expression/script/gocc/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Fscript	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Fscript : Expr	<<  >>`,
		Id:         "Fscript",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Fscript : TernaryExpr	<<  >>`,
		Id:         "Fscript",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "||" Expr1	<< ast.NewBoolExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr",
		NTType:     2,
		Index:      3,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr : Expr1	<<  >>`,
		Id:         "Expr",
		NTType:     2,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr1 : Expr1 "&&" Expr2	<< ast.NewBoolExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr1",
		NTType:     3,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr1 : Expr2	<<  >>`,
		Id:         "Expr1",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr2 : Expr2 "==" Expr3	<< ast.NewCmpExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr2",
		NTType:     4,
		Index:      7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmpExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr2 : Expr2 "!=" Expr3	<< ast.NewCmpExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr2",
		NTType:     4,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmpExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr2 : Expr2 "<" Expr3	<< ast.NewCmpExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr2",
		NTType:     4,
		Index:      9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmpExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr2 : Expr2 "<=" Expr3	<< ast.NewCmpExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr2",
		NTType:     4,
		Index:      10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmpExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr2 : Expr2 ">" Expr3	<< ast.NewCmpExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr2",
		NTType:     4,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmpExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr2 : Expr2 ">=" Expr3	<< ast.NewCmpExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr2",
		NTType:     4,
		Index:      12,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCmpExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr2 : Expr3	<<  >>`,
		Id:         "Expr2",
		NTType:     4,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr3 : Expr3 "+" Expr4	<< ast.NewArithExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr3",
		NTType:     5,
		Index:      14,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewArithExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr3 : Expr3 "-" Expr4	<< ast.NewArithExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr3",
		NTType:     5,
		Index:      15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewArithExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr3 : Expr4	<<  >>`,
		Id:         "Expr3",
		NTType:     5,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr4 : Expr4 "*" Expr5	<< ast.NewArithExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr4",
		NTType:     6,
		Index:      17,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewArithExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr4 : Expr4 "/" Expr5	<< ast.NewArithExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr4",
		NTType:     6,
		Index:      18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewArithExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr4 : Expr4 "%!"(MISSING) Expr5	<< ast.NewArithExpr(X[0],X[1],X[2]) >>`,
		Id:         "Expr4",
		NTType:     6,
		Index:      19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewArithExpr(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `Expr4 : Expr5	<<  >>`,
		Id:         "Expr4",
		NTType:     6,
		Index:      20,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr5 : Expr6	<<  >>`,
		Id:         "Expr5",
		NTType:     7,
		Index:      21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr5 : "-" Expr5	<< ast.NewUnaryExpr(X[0], X[1]) >>`,
		Id:         "Expr5",
		NTType:     7,
		Index:      22,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewUnaryExpr(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Expr5 : "!" Expr5	<< ast.NewUnaryExpr(X[0], X[1]) >>`,
		Id:         "Expr5",
		NTType:     7,
		Index:      23,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewUnaryExpr(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Expr6 : PrimaryExpr	<<  >>`,
		Id:         "Expr6",
		NTType:     8,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr6 : ident "(" Args ")"	<< ast.NewFuncExpr(X[0], X[2]) >>`,
		Id:         "Expr6",
		NTType:     8,
		Index:      25,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFuncExpr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Expr6 : functionName "(" Args ")"	<< ast.NewFuncExpr(X[0], X[2]) >>`,
		Id:         "Expr6",
		NTType:     8,
		Index:      26,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFuncExpr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `PrimaryExpr : Literal	<<  >>`,
		Id:         "PrimaryExpr",
		NTType:     9,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `PrimaryExpr : "(" Expr ")"	<< X[1], nil >>`,
		Id:         "PrimaryExpr",
		NTType:     9,
		Index:      28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `TernaryExpr : TernaryArgument "?" TernaryArgument ":" TernaryArgument	<< ast.NewTernaryExpr(X[0], X[2], X[4]) >>`,
		Id:         "TernaryExpr",
		NTType:     10,
		Index:      29,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTernaryExpr(X[0], X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `TernaryArgument : Expr	<<  >>`,
		Id:         "TernaryArgument",
		NTType:     11,
		Index:      30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TernaryArgument : TernaryExpr	<<  >>`,
		Id:         "TernaryArgument",
		NTType:     11,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `TernaryArgument : "(" TernaryExpr ")"	<< ast.NewTernaryArgument(X[1]) >>`,
		Id:         "TernaryArgument",
		NTType:     11,
		Index:      32,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTernaryArgument(X[1])
		},
	},
	ProdTabEntry{
		String: `BoolLit : "true"	<<  >>`,
		Id:         "BoolLit",
		NTType:     12,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `BoolLit : "false"	<<  >>`,
		Id:         "BoolLit",
		NTType:     12,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NilLit : "nil"	<<  >>`,
		Id:         "NilLit",
		NTType:     13,
		Index:      35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `NilLit : "null"	<<  >>`,
		Id:         "NilLit",
		NTType:     13,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Literal : intLit	<< ast.NewLiteral("int",X[0]) >>`,
		Id:         "Literal",
		NTType:     14,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLiteral("int",X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : floatLit	<< ast.NewLiteral("float",X[0]) >>`,
		Id:         "Literal",
		NTType:     14,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLiteral("float",X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : stringLit	<< ast.NewLiteral("string",X[0]) >>`,
		Id:         "Literal",
		NTType:     14,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLiteral("string",X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : BoolLit	<< ast.NewLiteral("bool",X[0]) >>`,
		Id:         "Literal",
		NTType:     14,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLiteral("bool",X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : NilLit	<< ast.NewLiteral("nil",X[0]) >>`,
		Id:         "Literal",
		NTType:     14,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLiteral("nil",X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : ref Ref	<< ast.NewRefExpr(X[0], X[1]) >>`,
		Id:         "Literal",
		NTType:     14,
		Index:      42,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRefExpr(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Ref : selector	<< ast.Concat(X[0]) >>`,
		Id:         "Ref",
		NTType:     15,
		Index:      43,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Concat(X[0])
		},
	},
	ProdTabEntry{
		String: `Ref : Indexer	<< ast.Concat(X[0]) >>`,
		Id:         "Ref",
		NTType:     15,
		Index:      44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Concat(X[0])
		},
	},
	ProdTabEntry{
		String: `Ref : Ref selector	<< ast.Concat(X[0], X[1]) >>`,
		Id:         "Ref",
		NTType:     15,
		Index:      45,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Concat(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Ref : Ref Indexer	<< ast.Concat(X[0], X[1]) >>`,
		Id:         "Ref",
		NTType:     15,
		Index:      46,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Concat(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Indexer : index	<< ast.Concat(X[0]) >>`,
		Id:         "Indexer",
		NTType:     16,
		Index:      47,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Concat(X[0])
		},
	},
	ProdTabEntry{
		String: `Indexer : "[" Expr "]"	<< ast.ConcatIndexer(X[0], X[1], X[2]) >>`,
		Id:         "Indexer",
		NTType:     16,
		Index:      48,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ConcatIndexer(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `Args : empty	<<  >>`,
		Id:         "Args",
		NTType:     17,
		Index:      49,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `Args : ExprList	<<  >>`,
		Id:         "Args",
		NTType:     17,
		Index:      50,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ExprList : Expr	<< ast.NewExprList(X[0]) >>`,
		Id:         "ExprList",
		NTType:     18,
		Index:      51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewExprList(X[0])
		},
	},
	ProdTabEntry{
		String: `ExprList : ExprList "," Expr	<< ast.AppendToExprList(X[0], X[2]) >>`,
		Id:         "ExprList",
		NTType:     18,
		Index:      52,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendToExprList(X[0], X[2])
		},
	},
}
