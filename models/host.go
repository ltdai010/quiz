package models

import (
	"errors"
	"math/rand"
	"quiz/temp"
)

var (
	hosts map[int]*Host
)

type Host struct {
	Name 	   			string
	Code       			int
	NumberOfParticipant int
}

func generateCode() int {
	var r int
	for {
		r = rand.Intn(900000)
		r += 100000 - 1
		if hosts[r] != nil {
			continue
		}
		break
	}
	return r
}

func init() {
	hosts = make(map[int]*Host)
}

func AddHost(host temp.HostUpdate) int {
	var h Host
	h.Code = generateCode()
	h.Name = host.Name
	h.NumberOfParticipant = host.NumberOfParticipant
	hosts[h.Code] = &h
	return h.Code
}

func GetOne(code int) (object *Host, err error) {
	if v, ok := hosts[code]; ok {
		return v, nil
	}
	return nil, errors.New("code not exist")
}

func GetAllHost() map[int]*Host {
	return hosts
}

func Update(Code int, host *temp.HostUpdate) (err error) {
	if v, ok := hosts[Code]; ok {
		v.Name = host.Name
		v.NumberOfParticipant = host.NumberOfParticipant
		return nil
	}
	return errors.New("code not exist")
}

func Delete(Code int) {
	delete(hosts, Code)
}

