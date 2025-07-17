package generalpurposeconverter

//Feet To Meters
func FtoM(f Feet) Meters { return Meters(f * 0.348)}
//Meters To Feet
func MtoF(m Meters) Feet { return Feet(m * 3.28084) }

//Lbs to Kilograms
func LtoK(lbs Pounds) Kilograms { return Kilograms(lbs * 0.453592) }
//Kilograms
func KtoL(k Kilograms) Pounds { return Pounds(k * 2.2)}


