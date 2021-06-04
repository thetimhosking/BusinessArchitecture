package main

type Requirement struct {
	ID         int
	Title      string
	Definition string
	Alias      string
	Rank       int
}

func AddRequirement(Title string, Definition string, Alias string, Rank int) (int, error) {

	var r Requirement
	var nextID int

	nextID = 1

	r.Alias = Alias
	r.Definition = Definition
	r.Title = Title
	r.Rank = Rank
	r.ID = nextID

	return nextID, nil

}
