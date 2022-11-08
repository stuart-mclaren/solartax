package main

import (
	"flag"
	"fmt"
)

func main() {
	taxRate := flag.Float64("incometaxrate", 0.4,
		"Income tax rate. For 40% use 0.4 here.")
	annualTaxAllowance := flag.Float64("annualtaxallowance", 200.0,
		"Income tax allowance. Eg 200.0 for first 200 euro income "+
			"every year to be untaxed")
	unitSellPrice := flag.Float64("unitsellprice", 0.14,
		"Price for selling a unit to the grid. eg 0.14 for 14 cents per unit")
	unitBuyPrice := flag.Float64("unitbuyprice", 0.31,
		"Price for selling a unit to the grid. eg 0.31 for 31 cents per unit")
	unitSellPriceInflation := flag.Float64("unitsellpriceinflation", 0.04,
		"Rate at which price for selling to "+
			"the grid increases. Eg for 5% per "+
			"year use 0.05 here")
	unitBuyPriceInflation := flag.Float64("unitbuypriceinflation", 0.04,
		"Rate at which price for selling to "+
			"the grid increases. Eg for 5% per "+
			"year use 0.05 here")
	solarPanelLifetime := flag.Float64("solarpanellifetime", 25.0,
		"Lifespan of solar panels (duration of investment). "+
			"Eg 25.0 for 25 years")
	selltoGridFraction := flag.Float64("selltogridfraction", 0.66,
		"Fraction of generated energy sold to grid. Eg if you export "+
			"2/3 of generated electricity and use 1/3, set to 0.66. "+
			"If you use 2/3 and export 1/3, set to 0.33")
	totalAnnualGeneration := flag.Float64("totalannualgeneration", 3200.0,
		"Total amount of electricity generated in a year in kilowatts. "+
			"Eg for 3.2 megawatts set to 3200")

	flag.Parse()

	fmt.Println("Input:")
	fmt.Printf("incometaxrate: %f\n", *taxRate)
	fmt.Printf("unitsellprice: %f\n", *unitSellPrice)
	fmt.Printf("unitbuyprice: %f\n", *unitBuyPrice)
	fmt.Printf("unitsellpriceinflation: %f\n", *unitSellPriceInflation)
	fmt.Printf("unitbuypriceinflation: %f\n", *unitBuyPriceInflation)
	fmt.Printf("solarpanellifetime: %f\n", *solarPanelLifetime)
	fmt.Printf("selltogridfraction: %f\n", *selltoGridFraction)
	fmt.Printf("totalAnnualGeneration: %f\n\n", *totalAnnualGeneration)

	currentUnitSellPrice := *unitSellPrice
	currentUnitBuyPrice := *unitBuyPrice

	var totalIncome float64 = 0
	var totalIncomeTax float64 = 0
	var totalSavings float64 = 0
	var taxableIncome float64

	for i := 1; i <= int(*solarPanelLifetime); i++ {
		fmt.Printf("Year %d:\n", i)
		fmt.Printf("Sell unit rate (euro): %f\n", currentUnitSellPrice)
		fmt.Printf("Buy unit rate (euro): %f\n", currentUnitBuyPrice)
		generatedUnits := (*totalAnnualGeneration)
		fmt.Printf("Generated units (kilowatts): %0.0f\n", generatedUnits)
		income := generatedUnits * *selltoGridFraction * currentUnitSellPrice
		valueOfConsumedSolarElectricity := generatedUnits *
			(1.0 - *selltoGridFraction) * currentUnitBuyPrice

		fmt.Printf("Value of consumed solar electricity units for this "+
			"year (euro): %f\n", valueOfConsumedSolarElectricity)
		totalSavings += valueOfConsumedSolarElectricity

		fmt.Printf("Income from selling this year (euro): %0.0f\n", income)
		if income < *annualTaxAllowance {
			taxableIncome = 0
		} else {
			taxableIncome = (income - *annualTaxAllowance)
		}

		tax := taxableIncome * *taxRate
		fmt.Printf("Taxable income from selling this year (euro): %0.0f\n",
			taxableIncome)
		fmt.Printf("Income tax due from selling this year (euro): %0.0f\n",
			tax)
		totalIncome += income
		fmt.Printf("Total income from selling so far (euro): %0.0f\n",
			totalIncome)
		fmt.Printf("Total savings so far (euro): %0.0f\n",
			totalSavings)
		fmt.Printf("Total pre-tax income plus savings so far (euro): %0.0f\n",
			totalIncome+totalSavings)
		totalIncomeTax += tax
		fmt.Printf("Total post-tax income plus savings so far (euro): %0.0f\n",
			totalIncome-totalIncomeTax+totalSavings)

		fmt.Printf("Total income tax due from selling so far (euro): %0.0f\n",
			totalIncomeTax)

		// price inflation
		currentUnitSellPrice = currentUnitSellPrice +
			(currentUnitSellPrice * *unitSellPriceInflation)
		currentUnitBuyPrice = currentUnitBuyPrice +
			(currentUnitBuyPrice * *unitBuyPriceInflation)
		fmt.Println()

	}
}
