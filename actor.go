package main

type ActorId int

const (
	NoSuchActor            ActorId = 0
	ProducerActor          ActorId = 1
	LevelOneDispatherActor ActorId = 2
	LevelTwoDispatherActor ActorId = 3
	DetailerActor          ActorId = 4
)

func (id ActorId) String() string {
	switch id {
	case NoSuchActor:
		return "NoSuchActor"
	case ProducerActor:
		return "Producer"
	case LevelOneDispatherActor:
		return "LevelOneDispather"
	case LevelTwoDispatherActor:
		return "LevelTwoDispather"
	case DetailerActor:
		return "Detailer"
	}
	return "error Actor"
}

var (
	NumbersForProducer = []int{
		1, 2, 3, 4, 5,
	}
	NumberForLevelOneDispathers = []int{
		6, 7, 8, 9, 10,
	}
	NumbersForLevelTwoDispather = []int{
		11, 12, 13, 14, 15,
	}
	NumbersForDetailer = []int{
		16, 17, 18, 19, 20,
	}
	NumbersForTruck = []int{
		21, 22, 23, 24, 25,
	}
)

func RoleIdToActor(id int) ActorId {
	for _, v := range NumbersForProducer {
		if v == id {
			return ProducerActor
		}
	}
	for _, v := range NumberForLevelOneDispathers {
		if v == id {
			return LevelOneDispatherActor
		}
	}
	for _, v := range NumbersForLevelTwoDispather {
		if v == id {
			return LevelTwoDispatherActor
		}
	}
	for _, v := range NumbersForDetailer {
		if v == id {
			return DetailerActor
		}
	}
	return NoSuchActor
}
