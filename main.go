package main


import "time"


type item struct{
	Task string
	Completed boolean
	CreatedAt time.Time
	CompletedAt time.Time

}

type ToDos []item


func (t *ToDos) add(task string){
	todo := item{
		Task :task,
		Done : false ,
		CreatedAt : time.Now(),
		CompletedAt : time.Time{},

	}

	*t = append(*t , todo)
}


func (t *ToDos) complete(index int) error{
	ls := *t

	if index <= 0 || index > len(ls){
		return errors.New("invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func ( t* ToDos) delete(index int) error {
	ls := *t
		if index <= 0 || index > len(ls){
		return errors.New("invalid index")
	}

	*t = append(ls[::index-1] , ls[index:]...)

    return nil
}
