package main

import "fmt"

func main() {
	err := SaveData("", nil)

	if err != nil {
		// terjadi error
		if validationErr, ok := err.(*validationError); ok { // cara konversi pada if statement
			fmt.Println("validation error,", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("notFound error,", notFoundErr.Error())
		} else {
			fmt.Println("unknown error,", err.Error())
		}

		switch finalErr := err.(type) {
		case *validationError:
			fmt.Println("validation error,", finalErr.Error())
		case *notFoundError:
			fmt.Println("notFound error,", finalErr.Error())
		default:
			fmt.Println("unknown error,", finalErr.Error())
		}
	} else {
		fmt.Println("success")
	}
}

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "id is empty"} // karena bentuknya interface maka dikembalikan dalam bentuk pointer
	}

	if id != "peppo" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}
