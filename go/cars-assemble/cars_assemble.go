package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return (float64(productionRate) * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(workingCarsPerHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const INDIVIDUAL_CAR_PRODUCTION_COST = 10_000
    const GROUP_CARS_PRODUCTION_COST = 95_000

    carGroup := carsCount / 10
    carRemaining := carsCount % 10

    totalCostOfOfProductionForCarGroup := uint(GROUP_CARS_PRODUCTION_COST * carGroup)
    totalCostOfProductionForIndividualCar := uint(INDIVIDUAL_CAR_PRODUCTION_COST * carRemaining)
    

    return totalCostOfOfProductionForCarGroup + totalCostOfProductionForIndividualCar
}
