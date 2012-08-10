package sim

import (
	"code.google.com/p/go-eco/eco/aux"
	"fmt"
	"testing"
)

// Test of two identical rows
func TestIdent(t *testing.T) {
	fmt.Println("Values of sim/dis indices for identical samples")
	var (
		data, out *Matrix
		x         float64
	)
	data = GetBoolIdent()

	out = WhittakerBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Whittaker: ", x)

	out = MountfordBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Mountford: ", x)

	out = YuleBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Yule: ", x)

	out = WilsonShmidaBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("WilsonShmida: ", x)

	out = CoCoGastonBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("CoCoGaston: ", x)

	out = OchiaiBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Ochiai: ", x)

	out = JaccardBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Jaccard: ", x)

	out = SorensenBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Sorensen: ", x)

	out = Williams1Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Williams1: ", x)

	out = Williams2Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Williams2: ", x)

	out = WeiherBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Weiher: ", x)

	out = StilesBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Stiles: ", x)

	out = SorgenfreiBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Sorgenfrei: ", x)

	out = SokalSneath1Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("SokalSneath1: ", x)

	out = SokalSneath2Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("SokalSneath2: ", x)

	out = SokalSneath3Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("SokalSneath3: ", x)

	out = SokalSneath4Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("SokalSneath4: ", x)

	out = SokalSneath5Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("SokalSneath5: ", x)

	out = Simpson1Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Simpson1: ", x)

	out = Simpson2Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Simpson2: ", x)

	out = SimpleMatchingBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("SimpleMatching: ", x)

	out = RuggieroBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Ruggiero: ", x)

	out = Routledge1Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Routledge1: ", x)

	out = Routledge2Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Routledge2: ", x)

	out = Routledge3Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Routledge3: ", x)

	out = RogersTanimotoBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("RogersTanimoto: ", x)

	out = PeirceBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Peirce: ", x)

	out = MichaelBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Michael: ", x)

	out = McConnaghBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("McConnagh: ", x)

	out = MargaleffBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Margaleff: ", x)

	out = ManhattanBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Manhattan: ", x)

	out = MagurranBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Magurran: ", x)

	out = MaarelBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Maarel: ", x)

	out = Lennon1Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Lennon1: ", x)

	out = Lennon2Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Lennon2: ", x)

	out = Legendre1Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Legendre1: ", x)

	out = Legendre2Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Legendre2: ", x)

	out = LandeBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Lande: ", x)

	out = LamontBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Lamont: ", x)

	out = Kulczynski1Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Kulczynski1: ", x)

	out = Kulczynski2Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Kulczynski2: ", x)

	out = Johnson1Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Johnson1: ", x)

	out = Johnson2Bool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Johnson2: ", x)

	out = HarteBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Harte: ", x)

	out = HarrisonBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Harrison: ", x)

	out = HamannBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Hamann: ", x)

	out = GowerBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Gower: ", x)

	out = FossumBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Fossum: ", x)

	out = ForbesBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Forbes: ", x)

	out = EyraudBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Eyraud: ", x)

	out = EuclidBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Euclid: ", x)

	out = DivergenceBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Divergence: ", x)

	out = DiceBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Dice: ", x)

	out = DennisBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Dennis: ", x)

	out = CodyBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Cody: ", x)

	out = ChiSquaredBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("ChiSquared: ", x)

	/* not implementyed yet
	out = BrayCurtisBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("BrayCurtis: ", x)

	out = BinomialBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Binomial: ", x)
	*/
	out = BaroniUrbaniBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("BaroniUrbani: ", x)

	out = FagerBool_S(data)
	x = out.Get(0, 1)
	fmt.Println("Fager: ", x)
}
