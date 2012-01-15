exe: clean
exe: clean
	windres resource.rc -o temp-rc.o 
	8g app.go util.go vars.go actions\actionrepo.go actions\engine.go actions\actor.go forms\mainform.go controls\statebutton.go controls\progressbutton.go controls\imgbutton.go
	gopack grc _go_.8 app.8 temp-rc.o
	8l -o dropdot.exe -s _go_.8
	rm *.8 *.o
clean:
	rm -f *.8 *.o *.exe