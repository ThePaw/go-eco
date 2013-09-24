package ser

// Test of obj fns aganist Kostopoulos values
import (
	"fmt"
	//	"math"
	"os"
	"testing"
)

func TestObjFns(t *testing.T) {
	file, err := os.Open("/home/pac/live/ext/src/code.google.com/p/ser/dat/artif/AR/ShuffledAR50.csv") // For read access.
	if err != nil {
		panic("file does not exist")
	}
	a := ReadCsvMatrix64(file)

	n := a.Rows()
	p := NewIntVector(n)
	p.Order()

	//AntiRobinson
	a.SimToDist()

	fmt.Println("G1Gain: ", G1Gain(a, p))
	fmt.Println("G2Gain: ", G2Gain(a, p))
	fmt.Println("G3Gain: ", G3Gain(a, p))
	fmt.Println("G4Gain: ", G4Gain(a, p))
	fmt.Println("HGain: ", HGain(a, p))
	fmt.Println("StrengLoss1: ", StrengLoss1(a, p))
	fmt.Println("StrengLoss2: ", StrengLoss2(a, p))
	fmt.Println("StrengLoss3: ", StrengLoss3(a, p))
	fmt.Println("InertiaGain: ", InertiaGain(a, p))
	fmt.Println("LeastSquaresLoss: ", LeastSquaresLoss(a, p))
	fmt.Println("MooreStressDisLoss: ", MooreStressDisLoss(a, p))
	fmt.Println("VonNeumannStressDisLoss: ", VonNeumannStressDisLoss(a, p))
	fmt.Println("GARLoss10: ", GARLoss5(a, p))
	fmt.Println("RGARLoss10: ", RGARLoss5(a, p))
	fmt.Println("HamiltonLoss: ", HamiltonLoss(a, p))
	fmt.Println("ParabolaLoss: ", ParabolaLoss(a, p))
	fmt.Println("QAPGain: ", QAPGain(a, p))
	fmt.Println("CompatibilityGain: ", CompatibilityGain(a, p))
	fmt.Println("WeightedCompatibilityGain: ", WeightedCompatibilityGain(a, p))
	fmt.Println("AREventsViolationLoss: ", AREventsViolationLoss(a, p))
	fmt.Println("WeightedAREventsViolationLoss: ", WeightedAREventsViolationLoss(a, p))
	fmt.Println("DoublyWeightedAREventsViolationLoss: ", DoublyWeightedAREventsViolationLoss(a, p))
	fmt.Println("GeneralizedARLoss10: ", GeneralizedARLoss10(a, p))
	fmt.Println("RelativeGARLoss10: ", RelativeGARLoss10(a, p))
	fmt.Println("EffectivenessGain: ", EffectivenessGain(a, p))
	fmt.Println("BertinLossSim: ", BertinLossSim(a, p))
	fmt.Println("MEffGainDis: ", MEffGainDis(a, p))
}
