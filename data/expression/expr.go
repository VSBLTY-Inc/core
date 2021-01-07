package expression

import (
	"github.com/VSBLTY-Inc/core/data"
	"github.com/VSBLTY-Inc/core/data/resolve"
)

type Factory interface {
	NewExpr(exprStr string) (Expr, error)
}

type Expr interface {
	Eval(scope data.Scope) (interface{}, error)
}

type FactoryCreatorFunc func(resolve.CompositeResolver) Factory
