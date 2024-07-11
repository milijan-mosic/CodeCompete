package lib

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

// func WriteJSONFile(filePath string, data taxi_models.WantedPoisRequest) {
// 	jsonData, err := json.MarshalIndent(data, "", "    ")
// 	if err != nil {
// 		log.Fatal("Error marshalling JSON -> ", err)
// 		return
// 	}

// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		log.Fatal("Error creating file -> ", err)
// 		return
// 	}
// 	defer file.Close()

// 	_, err = file.Write(jsonData)
// 	if err != nil {
// 		log.Fatal("Error writing JSON to file -> ", err)
// 		return
// 	}
// }

// func ReadBootstrappingData(filePath string) []taxi_models.PointsOfInterest {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		CheckForError("Error opening file -> ", err)
// 	}
// 	defer file.Close()

// 	data, err := io.ReadAll(file)
// 	if err != nil {
// 		CheckForError("Error reading file -> ", err)
// 	}

// 	var pois []taxi_models.PointsOfInterest

// 	if err := json.Unmarshal(data, &pois); err != nil {
// 		CheckForError("Error unmarshalling JSON -> ", err)
// 	}

// 	return pois
// }

func CheckForError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckSafely(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(filePath string) []string {
	content, err := os.ReadFile(filePath)
	Check(err)

	return strings.Split(string(content), "\n")
}

func WriteFile(filePath string, data string) {
	file, err := os.Create(filePath)
	Check(err)
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(data)
	Check(err)

	err = writer.Flush()
	Check(err)
}

// NOTE: Descending order
func SortListOfInts(list []int) {
	sort.Ints(list)
}

func StringToInt(char string) int {
	intVar, err := strconv.Atoi(char)
	Check(err)

	return intVar
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func Split(text, char string) []string {
	return strings.Split(text, char)
}

func Print(text string, data ...any) {
	fmt.Println(text, data)
}

func ContainsSubstring(str, substring string) bool {
	for i := 0; i < len(str)-len(substring)+1; i++ {
		if str[i:i+len(substring)] == substring {
			return true
		}
	}

	return false
}

func GenerateRandomString(length int16) string {
	bytes := make([]byte, (length+1)/2)

	_, err := rand.Read(bytes)
	Check((err))

	return base64.URLEncoding.EncodeToString(bytes)[:length]
}

func StringifySlice(inputSlice []string) string {
	jsonData, err := json.Marshal(inputSlice)
	Check(err)

	return string(jsonData)
}

func ParseSlice(jsonString string) []string {
	var list []string

	err := json.Unmarshal([]byte(jsonString), &list)
	Check(err)

	return list
}

func GenerateRandomPosition() (uint8, uint8, error) {
	positionOne, err := rand.Int(rand.Reader, big.NewInt(256))
	Check(err)

	positionTwo, err := rand.Int(rand.Reader, big.NewInt(256))
	Check(err)

	return uint8(positionOne.Int64()), uint8(positionTwo.Int64()), nil

}
