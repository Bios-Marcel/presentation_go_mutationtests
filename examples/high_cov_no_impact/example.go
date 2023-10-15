package highcovnoimpact

import (
	"fmt"
	"os"
)

func do(x int) {
	if x > 0 {
		sideEffectA(x)
	}

	sideEffectB(x)
}

func sideEffectA(x int) {
	toStr := fmt.Sprintf("%d_A", x)
	_ = os.WriteFile(toStr, []byte(toStr), 0o600)
}

func sideEffectB(x int) {
	toStr := fmt.Sprintf("%d_B", x)
	_ = os.WriteFile(toStr, []byte(toStr), 0o600)
}
