package main

import "fmt"

// PODEMOS ESCALAR DIFERENTES TIPOS DE ROCA, DE HIELO, CONCRETO, CANTO,ETXC...
// TRATAMOS DE QUE EL PROGRAMA DEPENDDA DE UN COMPORTAMIENTO Y NO DE UNA IMPLEMENTACION

type RockClimer struct {
	rocksClimbed int
	kind         int
	sp           SaftyPlace
}

func newRockClimber(sp SaftyPlace) *RockClimer {
	return &RockClimer{sp: sp}
}

type SaftyPlace interface {
	placeSafeties()
}

type IceSafety struct {
}

func (is IceSafety) placeSafeties() {
	fmt.Println("Placing ice safeties")
}

func (rc *RockClimer) climb() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.placeSafeties()
	}
}

func (rc *RockClimer) placeSafeties() {
	fmt.Println("Placing safeties")
}

func main() {

	// podemos inyectar la dependencia que queramos (SaftyPlace)
	rc := newRockClimber(IceSafety{})

	for i := 0; i < 10; i++ {
		rc.climb()
	}

}
