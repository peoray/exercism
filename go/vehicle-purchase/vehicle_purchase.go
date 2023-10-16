package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    msg := "is clearly the better choice."
    if (option1 < option2) {
        return fmt.Sprintf("%s %s", option1, msg)
	}
	return fmt.Sprintf("%s %s", option2, msg)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    const LESS_THAN_THREE_YEARS = 0.8
    const ATLEAST_TEN_YEARS = 0.5
    const ATLEAST_THREE_YEARS_LESS_THAN_TEN_YEARS = 0.7

    if age < 3 {
		return originalPrice * LESS_THAN_THREE_YEARS
    } 
	if age < 10 {
		return originalPrice * ATLEAST_THREE_YEARS_LESS_THAN_TEN_YEARS
    } 
    return originalPrice * ATLEAST_TEN_YEARS
}
