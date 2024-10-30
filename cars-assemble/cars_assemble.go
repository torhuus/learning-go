package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.00
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	perHour := CalculateWorkingCarsPerHour(productionRate, successRate)

	return int(perHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars. 10=9.5, 1=10
func CalculateCost(carsCount int) uint {
	remFromBulk := carsCount % 10
	bulk := carsCount / 10

	remFromBulkPrice := remFromBulk * 10000
	bulkPrice := bulk * 95000

	return uint(bulkPrice + remFromBulkPrice)

}
