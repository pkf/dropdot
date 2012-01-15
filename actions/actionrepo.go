package main

type CreateActionFunc func(jsonString string) Actor

type ActionRepo struct {
	actionDict map[string] CreateActionFunc
}

func NewActionRepo() *ActionRepo {
	ar := new(ActionRepo)
	ar.actionDict = make(map[string] CreateActionFunc)

	return ar
}

func (this *ActionRepo) Add(actionCaption string, caFunc CreateActionFunc) {
	if caFunc == nil {
		panic("caFunc cannot be nil")
	}

	this.actionDict[actionCaption] = caFunc
}

func (this *ActionRepo) Get(actionCaption string, jsonString string) Actor {
	if afunc, ok := this.actionDict[actionCaption]; ok {
		return afunc(jsonString)
	}	

	return nil
}
