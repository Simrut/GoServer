.PHONY=all clean install

all: webserver

webserver: verifyToken.go generateToken.go
	go build

install:
	#mkdir -p mkosi.extra/opt
	cp webserver /home/nomis/Documents/PSI_Proj/PsiTasks/VM/tokenserver/mkosi.extra/opt
clean:
	$(RM) webserver