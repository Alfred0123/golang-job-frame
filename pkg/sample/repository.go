package sample

type repository struct {}

func (r *repository) GetSampleById (id int) (string) {

	if (id != 1) {
		panic("id not available")
	}

	return "sample"
}