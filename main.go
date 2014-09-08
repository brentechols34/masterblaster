package main

import ()

func main() {

}

func NextId () func() uint32 {
	var id uint32
	id = 0

	return func() uint32 {
		return id++
	}
}