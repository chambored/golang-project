package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Cell represents a mobile phone with various properties.
type Cell struct {
	// company phone comes from
	oem string
	// model of phone
	model string
	// year of launch announced
	launchAnnounced uint
	// year of launch
	launchStatus string
	// dimensions of the phone's body
	bodyDimensions string
	// weight of phone's body
	bodyWeight float32
	// type of sim card
	bodySim string
	// type of display
	displayType string
	// size of display in inches
	displaySize float32
	// resolution of display
	displayResolution string
	// any features that are sensors
	featuresSensors string
	// platform of the operating system of the phone
	platformOS string
}

// NewCell creates a new Cell with the given properties.
func NewCell(oem string, model string, launchAnnounced uint,
	launchStatus string, bodyDimensions string, bodyWeight float32,
	bodySim string, displayType string, displaySize float32,
	displayResolution string, featuresSensors string, platformOS string) *Cell {
	return &Cell{
		oem:               oem,
		model:             model,
		launchAnnounced:   launchAnnounced,
		launchStatus:      launchStatus,
		bodyDimensions:    bodyDimensions,
		bodyWeight:        bodyWeight,
		bodySim:           bodySim,
		displayType:       displayType,
		displaySize:       displaySize,
		displayResolution: displayResolution,
		featuresSensors:   featuresSensors,
		platformOS:        platformOS,
	}
}

// String implements the Stringer interface for the Cell struct.
func (c Cell) String() string {
	return fmt.Sprintf("\nOEM: %s\nModel: %s\nLaunch Announced: %d\nLaunch Status: %s\nBody Dimensions: %s\nBody Weight: %.2f g\nSIM: %s\nDisplay Type: %s\nDisplay Size: %.2f in\nDisplay Resolution: %s\nSensors: %s\nPlatform OS: %s\n", c.oem, c.model, c.launchAnnounced, c.launchStatus, c.bodyDimensions, c.bodyWeight, c.bodySim, c.displayType, c.displaySize, c.displayResolution, c.featuresSensors, c.platformOS)
}

// averageWeight calculates the average weight of the phones
// in the given map. Only phones with a non-zero weight are considered
// in the average.
func averageWeight(cells map[string]*Cell) float32 {
	var totalWeight float32
	var count int

	for _, cell := range cells {
		currentWeight := cell.bodyWeight
		if currentWeight > 0 {
			totalWeight += currentWeight
			count++
		}
	}

	if count > 0 {
		return totalWeight / float32(count)
	} else {
		return 0
	}
}

// averageDisplaySize calculates the average display size of the phones
// in the given map. Only phones with a non-zero display size are considered
// in the average.
func averageDisplaySize(cells map[string]*Cell) float32 {
	var totalSize float32
	var count int

	for _, cell := range cells {
		currentSize := cell.displaySize
		if currentSize > 0 {
			totalSize += currentSize
			count++
		}
	}

	if count > 0 {
		return totalSize / float32(count)
	} else {
		return 0
	}
}

// YearCounts represents the number of phones launched in each year.
type YearCounts struct {
	Counts map[uint]int
	Years  []uint
}

// countPhonesByYear counts the number of phones launched in each year
// and returns a YearCounts object which includes a sorted list of years
// for which data exists.
func countPhonesByYear(cells map[string]*Cell) YearCounts {
	counts := make(map[uint]int)

	for _, cell := range cells {
		if cell.launchAnnounced != 0 {
			counts[cell.launchAnnounced]++
		}
	}

	// Collect the years in a slice.
	years := make([]uint, 0, len(counts))
	for year := range counts {
		years = append(years, year)
	}

	// Sort the years.
	sort.Slice(years, func(i, j int) bool { return years[i] < years[j] })

	return YearCounts{
		Counts: counts,
		Years:  years,
	}
}

// countUniqueOS counts the number of unique operating systems
// used in the given map of cells.
func countUniqueOS(cells map[string]*Cell) int {
	osSet := make(map[string]struct{})

	for _, cell := range cells {
		osSet[cell.platformOS] = struct{}{}
	}

	return len(osSet)
}

// countPhonesByOEM counts the number of phones produced by each OEM
// in the given map of cells.
func countPhonesByOEM(cells map[string]*Cell) map[string]int {
	oemCount := make(map[string]int)
	for _, cell := range cells {
		oemCount[cell.oem]++
	}
	return oemCount
}

// findLatestPhoneByOEM finds the most recently launched phone by each OEM
// in the given map of cells.
func findLatestPhoneByOEM(cells map[string]*Cell) map[string]*Cell {
	oemLatest := make(map[string]*Cell)
	for _, cell := range cells {
		if oemLatest[cell.oem] == nil || cell.launchAnnounced > oemLatest[cell.oem].launchAnnounced {
			oemLatest[cell.oem] = cell
		}
	}
	return oemLatest
}

// findHeaviestAndLightestPhones finds the heaviest and lightest phones
// in the given map of cells. Phones with a weight of zero are excluded.
func findHeaviestAndLightestPhones(cells map[string]*Cell) (*Cell, *Cell) {
	var heaviest, lightest *Cell
	for _, cell := range cells {
		// Exclude cells with bodyWeight of zero
		if cell.bodyWeight == 0 {
			continue
		}
		if heaviest == nil || cell.bodyWeight > heaviest.bodyWeight {
			heaviest = cell
		}
		if lightest == nil || cell.bodyWeight < lightest.bodyWeight {
			lightest = cell
		}
	}
	return heaviest, lightest
}

// averageWeightByOEM calculates the average weight of the phones
// for each OEM in the given map. Only phones with a non-zero weight are considered
// in the average. It returns a map of OEMs and their average phone weights.
func averageWeightByOEM(cells map[string]*Cell) map[string]float32 {
	oemWeights := make(map[string]float32)
	oemCounts := make(map[string]int)

	for _, cell := range cells {
		currentWeight := cell.bodyWeight
		if currentWeight > 0 {
			oemWeights[cell.oem] += currentWeight
			oemCounts[cell.oem]++
		}
	}

	averageWeights := make(map[string]float32)
	for oem, totalWeight := range oemWeights {
		count := oemCounts[oem]
		if count > 0 {
			averageWeights[oem] = totalWeight / float32(count)
		} else {
			averageWeights[oem] = 0
		}
	}

	return averageWeights
}

func findOEMWithHighestAverageWeight(cells map[string]*Cell) string {
	averages := averageWeightByOEM(cells)
	var maxOEM string
	var maxAvg float32
	for oem, avg := range averages {
		if avg > maxAvg {
			maxOEM = oem
			maxAvg = avg
		}
	}
	return maxOEM
}

// PhoneDetails struct to hold the oem and model of a phone
type PhoneDetails struct {
	oem   string
	model string
}

// findPhonesAnnouncedAndReleasedDifferentYears checks if there are any phones that were announced
// in one year and released in another. If such phones exist, it returns their OEM and model.
func findPhonesAnnouncedAndReleasedDifferentYears(cells map[string]*Cell) []PhoneDetails {
	var phoneDetails []PhoneDetails
	for _, cell := range cells {
		announcedYear := cell.launchAnnounced
		releasedYear := cell.launchReleased
		if announcedYear != releasedYear {
			phoneDetails = append(phoneDetails, PhoneDetails{cell.oem, cell.model})
		}
	}
	return phoneDetails
}

// countPhonesWithOneSensor counts the number of phones with only one feature sensor
func countPhonesWithOneSensor(cells map[string]*Cell) int {
	count := 0
	for _, cell := range cells {
		sensors := strings.Split(cell.featuresSensors, ",")
		if len(sensors) == 1 {
			count++
		}
	}
	return count
}

// findMostLaunchesIn2000s returns the year in the 2000s that had the most phone launches
func findMostLaunchesIn2000s(yearCounts YearCounts) uint {
	var maxYear uint
	var maxCount int

	for year, count := range yearCounts.Counts {
		if year >= 2000 && year < 2010 && count > maxCount {
			maxYear = year
			maxCount = count
		}
	}

	return maxYear
}

// parseYear extracts a 4-digit year from a string. If no 4-digit year is found or if
// an error occurs during conversion, it returns nil. Otherwise, it returns a pointer to
// the extracted year.
func parseYear(yearStr string) *uint {
	// Find any 4-digit number in the string
	re := regexp.MustCompile("\\b\\d{4}\\b")
	match := re.FindString(yearStr)

	// If no match was found, return nil
	if match == "" {
		return nil
	}

	// Convert the match to an integer
	year, err := strconv.Atoi(match)
	// If an error occurred during conversion, return nil
	if err != nil {
		return nil
	}

	// Convert the year to an uint and return a pointer to it
	yearUint := uint(year)
	return &yearUint
}

// parseWeight extracts a weight in grams from a string. If no valid weight is found or if
// an error occurs during conversion, it returns nil. Otherwise, it returns a pointer to
// the extracted weight.
func parseWeight(weightStr string) *float32 {
	// Find any number followed by " g" in the string
	re := regexp.MustCompile("(\\d+(\\.\\d+)?)\\s* g")
	match := re.FindStringSubmatch(weightStr)

	// If no match was found, return nil
	if len(match) == 0 {
		return nil
	}

	// Convert the match to a float
	weight, err := strconv.ParseFloat(match[1], 32)
	// If an error occurred during conversion, return nil
	if err != nil {
		return nil
	}

	// Convert the weight to a float32 and return a pointer to it
	weightFloat := float32(weight)
	return &weightFloat
}

// parseSim checks a string for the values "no" or "yes" (case-insensitive). If either
// of these values is found, it returns nil. Otherwise, it returns a pointer to the original string.
func parseSim(simStr string) *string {
	if strings.ToLower(simStr) == "no" || strings.ToLower(simStr) == "yes" {
		return nil
	}
	return &simStr
}

// parseSize extracts a size in inches from a string. If no valid size is found or if
// an error occurs during conversion, it returns nil. Otherwise, it returns a pointer to
// the extracted size.
func parseSize(sizeStr string) *float32 {
	// Find any number followed by " inches" in the string
	re := regexp.MustCompile("(\\d+(\\.\\d+)?)\\s* inches")
	match := re.FindStringSubmatch(sizeStr)

	// If no match was found, return nil
	if len(match) == 0 {
		return nil
	}

	// Convert the match to a float
	size, err := strconv.ParseFloat(match[1], 32)
	// If an error occurred during conversion, return nil
	if err != nil {
		return nil
	}

	// Convert the size to a float32 and return a pointer to it
	sizeFloat := float32(size)
	return &sizeFloat
}

// parseSensors checks if a string consists only of digits (with a decimal point). If so,
// it returns nil. Otherwise, it returns a pointer to the original string.
func parseSensors(sensorStr string) *string {
	re := regexp.MustCompile("^(\\d+(\\.\\d+)?)$")
	match := re.FindStringSubmatch(sensorStr)

	// If a match was found, return nil
	if len(match) > 0 {
		return nil
	}

	return &sensorStr
}

// parsePlatformOS checks if a string consists only of digits (with a decimal point). If so,
// it returns nil. Otherwise, it returns a pointer to the first comma-separated element in
// the original string.
func parsePlatformOS(osStr string) *string {
	re := regexp.MustCompile("^(\\d+(\\.\\d+)?)$")
	match := re.FindStringSubmatch(osStr)

	// If a match was found, return nil
	if len(match) > 0 {
		return nil
	}

	// Split the string on the comma and return a pointer to the first element
	parts := strings.SplitN(osStr, ",", 2)
	return &parts[0]
}

func main() {
	// initialize a new map where the keys are strings
	// and the values are pointers to Cell structs. The map is named cells
	cells := make(map[string]*Cell)

	// open the file named cells.csv for reading. it returns
	// an *os.File and an error. if the file opens successfully,
	// err will be nil, else it will contain information about the problem
	file, err := os.Open("resources/cells.csv")

	// if there is an error opening the file, the program terminates execution
	// and prints the error information
	if err != nil {
		panic(err)
	}

	// create a csv reader from the file
	reader := csv.NewReader(file)

	// read the first line of cells.csv and ignore it,
	// panic if there is an error
	if _, err := reader.Read(); err != nil {
		panic(err)
	}

	// for loop continuing indefinitely
	for {
		// reads the next line from the csv, if there is an error in this
		// operation err will not be nil, stores a record of the fields in line
		line, err := reader.Read()
		// if there is an error in reading the line, break out of the loop
		if err != nil {
			break
		}

		// Parse the year from the third column of the line
		launchPtr := parseYear(line[2])
		// Create a variable to hold the year. This will default to 0 if parsing fails
		var launchYear uint = 0
		// If parsing was successful (i.e., the pointer is not nil), set launchYear to the parsed year
		if launchPtr != nil {
			launchYear = *launchPtr
		}

		// Parse the weight from the sixth column of the line
		weightPtr := parseWeight(line[5])
		// Create a variable to hold the weight. This will default to 0.0 if parsing fails
		var weight float32 = 0.0
		// If parsing was successful, set weight to the parsed weight
		if weightPtr != nil {
			weight = *weightPtr
		}

		// Parse the SIM from the seventh column of the line
		simPtr := parseSim(line[6])
		// Create a variable to hold the SIM. This will default to an empty string if parsing fails
		var sim = ""
		// If parsing was successful, set sim to the parsed SIM
		if simPtr != nil {
			sim = *simPtr
		}

		// Parse the size from the ninth column of the line
		sizePtr := parseSize(line[8])
		// Create a variable to hold the size. This will default to 0.0 if parsing fails
		var size float32 = 0.0
		// If parsing was successful, set size to the parsed size
		if sizePtr != nil {
			size = *sizePtr
		}

		// Parse the sensors from the eleventh column of the line
		sensorPtr := parseSensors(line[10])
		// Create a variable to hold the sensors. This will default to an empty string if parsing fails
		var sensors = ""
		// If parsing was successful, set sensors to the parsed sensors
		if sensorPtr != nil {
			sensors = *sensorPtr
		}

		// Parse the OS from the twelfth column of the line
		osPtr := parsePlatformOS(line[11])
		// Create a variable to hold the OS. This will default to an empty string if parsing fails
		var osPlat = ""
		// If parsing was successful, set osPlat to the parsed OS
		if osPtr != nil {
			osPlat = *osPtr
		}

		// create a new Cell and assign it cell using the NewCell func
		// pulling data from the line record containing all cell fields
		cell := NewCell(line[0], line[1], launchYear, line[3],
			line[4], weight, sim, line[7],
			size, line[9], sensors, osPlat)

		// create a string combination of the manufacturer and model
		// of the phone, seperated by a dash
		key := fmt.Sprintf("%s-%s", cell.oem, cell.model)

		// insert cell into the cells map, where the key is the formatted
		// string from above
		cells[key] = cell
	}

	// Printing the details of a specific cell phone
	fmt.Print("Sample Cell Phone Output: ")
	fmt.Print(cells["Google-Pixel 4 XL"])
	fmt.Println()

	// Calculating and displaying the statistics for the cell phone collection
	fmt.Println("Collection Statistics:")
	fmt.Printf("Average cell weight: %.2f g \n", averageWeight(cells))
	fmt.Printf("Average cell size: %.2f in \n", averageDisplaySize(cells))
	fmt.Printf("Number of Unique Operating Systems: %d\n", countUniqueOS(cells))
	fmt.Printf("There are %d phones with only one feature sensor.\n", countPhonesWithOneSensor(cells))

	// Finding and printing the heaviest and lightest phones
	heaviest, lightest := findHeaviestAndLightestPhones(cells)
	fmt.Printf("Heaviest Phone: %s\n", heaviest.oem+"'s "+heaviest.model+", "+fmt.Sprintf("%.2f", heaviest.bodyWeight)+" g")
	fmt.Printf("Lightest Phone: %s\n", lightest.oem+"'s "+lightest.model+", "+fmt.Sprintf("%.2f", lightest.bodyWeight)+" g")
	fmt.Println()

	// Find the OEM with the highest average weight, and print the result
    fmt.Println("The OEM with the highest average phone body weight is:", findOEMWithHighestAverageWeight(cells))

	// Counting phones released each year and printing the result
	fmt.Println("Number of cell announcements by year:")
	counts := countPhonesByYear(cells)
	for _, year := range counts.Years {
		fmt.Printf("%d: %d\n", year, counts.Counts[year])
	}
	fmt.Printf("The year with the most phone launches in the 2000s was %d.\n", findMostLaunchesIn2000s(counts))
	fmt.Println()

	// Counting phones by OEM and finding the latest phone model for each OEM
	fmt.Println("Count of phones and the latest model by each OEM:")
	fmt.Printf("%-15s %-10s %-25s\n", "OEM", "Count", "Latest Model")
	oemCounts := countPhonesByOEM(cells)
	latestPhones := findLatestPhoneByOEM(cells)
	for oem, count := range oemCounts {
		fmt.Printf("%-15s %-10d %-25s\n", oem, count, latestPhones[oem].model)
	}

	// Find the phones that were announced and released in different years
    phones := findPhonesAnnouncedAndReleasedDifferentYears(cells)
    // Print the result
    if len(phones) > 0 {
        fmt.Println("The following phones were announced and released in different years:")
        for _, phone := range phones {
            fmt.Printf("OEM: %s, Model: %s\n", phone.oem, phone.model)
        }
    } else {
        fmt.Println("No phones were announced and released in different years.")
    }
}
