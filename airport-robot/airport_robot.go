package airportrobot

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, say Greeter) string {
	return say.Greet(name)
}

type Italian struct {
	
}

func (it Italian) LanguageName() string {
	return "Italian"
}

func (it Italian) Greet(name string) string {
	return "I can speak " + it.LanguageName() + ": Ciao " + name + "!"
}

type Portuguese struct {
	
}

func (pg Portuguese) LanguageName() string {
	return "Portuguese"
}

func (pg Portuguese) Greet(name string) string {
	return "I can speak " + pg.LanguageName() + ": Olá " + name + "!"
}
