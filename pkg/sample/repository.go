package sample

type Repository struct {}

func (r *Repository) GetSampleById (id int) (string) {
	// println("GetSampleById!")
	if (id != 1) {
		panic("id not available")
	}

	return "sample"
}