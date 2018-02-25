package main

import (
	pb "go-time/vessel-service/proto/vessel"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

const (
	dbName                = "shippy"
	consignmentCollection = "vessels"
)

/*
Repository bla bla
*/
type Repository interface {
	Create(*pb.Vessel) error
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Close()
}

/*
VesselRepository bla bla
*/
type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(consignmentCollection)
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)

	if err != nil {
		return nil, err
	}

	return vessel, nil
}

/*
Close bla bla
*/
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

/*
Create bla bla
*/
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}
