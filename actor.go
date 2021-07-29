package main

type ActorId int

const (
	NoSuchActor       ActorId = 0
	Producer          ActorId = 1
	LevelOneDispather ActorId = 2
	LevelTwoDispather ActorId = 3
	Detailer          ActorId = 4
)

func (id ActorId) String() string {
	switch id {
	case NoSuchActor:
		return "NoSuchActor"
	case Producer:
		return "Producer"
	case LevelOneDispather:
		return "LevelOneDispather"
	case LevelTwoDispather:
		return "LevelTwoDispather"
	case Detailer:
		return "Detailer"
	}
	return "error Actor"
}

var (
	Producers = []int{
		1, 2, 3, 4, 5,
	}
	LevelOneDispathers = []int{
		6, 7, 8, 9, 10,
	}
	LevelTwoDispathers = []int{
		11, 12, 13, 14, 15,
	}
	Detailers = []int{
		16, 17, 18, 19, 20,
	}
)

func RoleIdToActor(id int) ActorId {
	for _, v := range Producers {
		if v == id {
			return Producer
		}
	}
	for _, v := range LevelOneDispathers {
		if v == id {
			return LevelOneDispather
		}
	}
	for _, v := range LevelTwoDispathers {
		if v == id {
			return LevelTwoDispather
		}
	}
	for _, v := range Detailers {
		if v == id {
			return Detailer
		}
	}
	return NoSuchActor
}
