package govalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)


func TestValidation(t *testing.T) {
	validate  :=  validator.New()

	if validate == nil {
		t.Error("validate is nil")
	}
}


func TestValidationField(t *testing.T) {
	validate   :=  validator.New()
	var user string = "arif"

	err :=  validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}


func TestValidasiDuaVar(t *testing.T) {
	validate    :=  validator.New()

	password    		:= "rahasia"
	confirmPassword		:= "salah"
	// cek kesamaan dua var dengan "eqfield"
	err :=  validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}


func TestMultipleTag(t *testing.T) {
	validate    :=  validator.New()
	var user string = "122"

	// number, numeric dll
	err := validate.Var(user, "required,number")
	if err != nil {
		fmt.Println(err.Error())
	}
}


func TestTagParameter(t *testing.T) {
	validate    :=  validator.New()
	user := "439"

	err := validate.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}



func TestValidationStruct(t *testing.T) {
	validate    :=  validator.New()
	type LoginRequest struct {
		Username  string `validate:"required,email"`
		Password  string `validate:"required,min=5"`
	}

	loginRequest := LoginRequest{
		Username: "eki",
		Password: "eki",
	}
	err  := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}



func TestValidationErrors(t *testing.T) {
	validate    :=  validator.New()
	type LoginRequest struct {
		Username  string `validate:"required,email"`
		Password  string `validate:"required,min=5"`
	}

	loginRequest := LoginRequest{
		Username: "eki",
		Password: "eki",
	}
	err  := validate.Struct(loginRequest)// Test implementation here
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors{
			fmt.Println("error", fieldErr.Field(), "on tag", fieldErr.Tag(), "with error", fieldErr.Error())
		}
	}
}



func TestName(t *testing.T) {
	validate    :=  validator.New()
	type RegisterUser struct {
		Username  string `validate:"required,email"`
		Password  string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	registerUser := RegisterUser{
		Username: "eki",
		Password: "ekiss",
		ConfirmPassword: "ekiss",
	}
	err  := validate.Struct(registerUser)
	if err != nil {
		fmt.Println(err.Error())
	}
}



func TestValidationNestedStruct(t *testing.T) {
	validate    :=  validator.New()

	type Address struct {
		City     string   	`validate:"required"`
		Country  string     `validate:"required"`
	}
	
	type User struct {
		Id     string   	`validate:"required"`
		Name    string   	`validate:"required"`
		Address  Address    `validate:"required"`
	}

	request  :=  User{
		Id: "",
		Name: "",
		Address: Address{
			City: "",
			Country: "",
		},
	}
	err  := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}


func TestValidationCollection(t *testing.T) {
	// Test implementation here
}