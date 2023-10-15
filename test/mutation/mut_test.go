//go:build mutation

package mutation_test

import (
	"go/ast"
	"go/types"
	"testing"

	"github.com/gtramontina/ooze"
	"github.com/gtramontina/ooze/viruses"
	"github.com/gtramontina/ooze/viruses/arithmetic"
	"github.com/gtramontina/ooze/viruses/arithmeticassignment"
	"github.com/gtramontina/ooze/viruses/arithmeticassignmentinvert"
	"github.com/gtramontina/ooze/viruses/bitwise"
	"github.com/gtramontina/ooze/viruses/comparison"
	"github.com/gtramontina/ooze/viruses/comparisoninvert"
	"github.com/gtramontina/ooze/viruses/comparisonreplace"
	"github.com/gtramontina/ooze/viruses/floatdecrement"
	"github.com/gtramontina/ooze/viruses/floatincrement"
	"github.com/gtramontina/ooze/viruses/integerdecrement"
	"github.com/gtramontina/ooze/viruses/integerincrement"
	"github.com/gtramontina/ooze/viruses/loopbreak"
	"github.com/gtramontina/ooze/viruses/loopcondition"
	"github.com/gtramontina/ooze/viruses/rangebreak"
)

func TestMutation(t *testing.T) {
	ooze.Release(t,
		// Tests are always run with the test directory as the working
		// directory, therefore we can reliably hardcode the root directory.
		ooze.WithRepositoryRoot("../../"),
		// -v: verbose output
		// -count=1: disable test caching
		// -failfast: stop on first failure
		// -timeout: Slow tests mean deadlock, meaning failure caught by timeout
		ooze.WithTestCommand("go test -v -count=1 -timeout=100ms -failfast ./examples/..."),
		// Parallelises the different mutated suites, but not the tests
		// functions, these are still controlled by t.Parallel().
		ooze.Parallel(),
		// Explicitly list viruses, since that's required when listing ANY
		// custom virus.
		ooze.WithViruses(
			arithmetic.New(),
			[]viruses.Virus{
				// Custom virus
				// customVirus{},
				arithmeticassignment.New(),
				arithmeticassignmentinvert.New(),
				bitwise.New(),
				comparison.New(),
				comparisoninvert.New(),
				comparisonreplace.New(),
				floatdecrement.New(),
				floatincrement.New(),
				integerdecrement.New(),
				integerincrement.New(),
				loopbreak.New(),
				loopcondition.New(),
				rangebreak.New(),
			}...),
		ooze.ForceColors(),
		ooze.IgnoreSourceFiles(`\.git`),
	)
}

type customVirus struct{}

func (c customVirus) Incubate(node ast.Node, typeInfo *types.Info) []*viruses.Infection {
	call, ok := node.(*ast.CallExpr)
	if !ok {
		return nil
	}

	ident, ok := call.Fun.(*ast.Ident)
	if !ok {
		return nil
	}

	if ident.Name == "sideEffectB" {
		return []*viruses.Infection{
			viruses.NewInfection(
				"Replace 'sideEffectB' with 'println'",
				func() {
					ident.Name = "println"
				},
				func() {
					ident.Name = "sideEffectB"
				},
			),
		}
	}

	return nil
}
