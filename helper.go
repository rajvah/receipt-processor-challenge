package main

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	layoutDate = "2006-01-02"
	layoutTime = "15:04"
	start      = "14:00"
	end        = "16:00"
)

// driver function to calculate points
// takes reciept object, returns Point, err
func calculatePoints(reciept *Receipt) (Point, error) {
	point := Point{}

	calculatedPoints := countAlphanumeric(reciept.Retailer)

	totalCount, err := countTotalAmount(reciept.Total)

	if err != nil {
		return point, err
	}
	calculatedPoints += totalCount

	calculatedPoints += int64(len(reciept.Items) / 2 * 5)

	for _, item := range reciept.Items {
		calculatedPoints += countItemDescription(item)
	}

	dateCount, err := countDate(reciept.PurchaseDate)
	if err != nil {
		return point, err
	}

	calculatedPoints += dateCount

	timeCount, err := countTime(reciept.PurchaseTime)
	if err != nil {
		return point, err
	}

	calculatedPoints += timeCount

	point.Points = calculatedPoints
	return point, nil

}

// takes retailer name as string, return calculated points
func countAlphanumeric(name string) int64 {
	count := int64(0)
	for _, char := range name {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// takes total amount  as string, return calculated points
func countTotalAmount(total string) (int64, error) {
	value, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return 0, err
	}

	// Round the value to the nearest dollar

	rounded := math.Round(value)

	res := 0
	// Check if the rounded value is equal to the original value
	if rounded == value {
		res += 50
	}
	if math.Mod(value, 0.25) == 0 {
		res += 25
	}
	return int64(res), nil
}

// for each item, calculates points for that
// takes Item as input, return calculated points
func countItemDescription(item Item) int64 {

	newDescription := strings.Trim(item.ShortDescription, " ")
	frac := len(newDescription) % 3
	if frac == 0 {
		value, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0
		}
		return int64(math.Ceil(value * 0.2))
	}

	return 0
}

// takes purchaseDate as string, return calculated points
func countDate(purchaseDate string) (int64, error) {

	val, err := time.Parse(layoutDate, purchaseDate)

	if err != nil {
		return 0, err
	}

	// ignoring year and month as we only need to check day
	_, _, day := val.Date()

	if day%2 == 1 {
		return 6, nil
	}
	return 0, nil
}

// takes purchaseTime as string, return calculated points
func countTime(purchaseTime string) (int64, error) {

	startTime, _ := time.Parse(layoutTime, start)
	endTime, _ := time.Parse(layoutTime, end)

	currTime, err := time.Parse(layoutTime, purchaseTime)
	if err != nil {
		return 0, err
	}
	if currTime.After(startTime) && currTime.Before(endTime) {
		return 10, nil
	}
	return 0, nil
}
