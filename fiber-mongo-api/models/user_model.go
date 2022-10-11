package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}

// type VMs struct {
// 	Id             primitive.ObjectID `json:"id,omitempty"`
// 	AccountType    string             `json:"AccountType,omitempty" validate:"required"`    //AWS,Azure,GCP
// 	Account        string             `json:"Account,omitempty" validate:"required"`        //AWS Account,Azure Subscription
// 	Group          string             `json:"Group,omitempty" validate:"required"`          //AWS: Region, Azure:Resource Group
// 	VMName         string             `json:"VMName,omitempty" validate:"required"`         //Virtual Machine Name
// 	VMId           string             `json:"VMId,omitempty" validate:"required"`           //Virtual Machine ID
// 	VMIpAddress    string             `json:"VMIpAddress,omitempty" validate:"required"`    //VirtualMachine IP Address
// 	VMState        string             `json:"VMState,omitempty" validate:"required"`        //Virtual Machine state power on/off
// 	VMStateShedule string             `json:"VMStateShedule,omitempty" validate:"required"` //Virtual Machine state shedule time

// }
