package facade

func NewAPI() API {
	return &apiImpl{
		s1: newStep1(),
		s2: newStep2(),
		s3: newStep3(),
	}
}

//API is facade interface of facade package
type API interface {
	Run()
}

//facade implement
type apiImpl struct {
	s1 *step1
	s2 *step2
	s3 *step3
}

func (a *apiImpl) Run() {
	a.s1.run()
	a.s2.run()
	a.s3.run()
}
