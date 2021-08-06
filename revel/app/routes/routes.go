// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tAddressController struct {}
var AddressController tAddressController


func (_ tAddressController) GetPersonAddressByID(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AddressController.GetPersonAddressByID", args).URL
}


type tErrorController struct {}
var ErrorController tErrorController


func (_ tErrorController) GetErrorByID(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("ErrorController.GetErrorByID", args).URL
}


type tPersonController struct {}
var PersonController tPersonController


func (_ tPersonController) GetPersonByID(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("PersonController.GetPersonByID", args).URL
}


