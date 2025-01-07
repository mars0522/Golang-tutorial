package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func WriteBalanceToFile(bal float64, filename string) error {
	balance := fmt.Sprint(bal) // Other way to convert any data type to string
	err := os.WriteFile(filename, []byte(balance), 0644)
	if err != nil {
		fmt.Println("error writing balance to file: ", err)
		return err
	}
	return nil
}
func ReadFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, err
	}
	return balance, nil

}
