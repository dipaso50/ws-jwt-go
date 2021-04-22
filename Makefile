exeName=jwtExample
sourceName=main.go
dockerImageTag=jws-go-img

all: compileLinux compileWindows

compileLinux:
	go build -o bin/linux/$(exeName) $(sourceName)

compileWindows:
	GOOS=windows GOARCH=386 go build -o bin/windows/$(exeName).exe $(sourceName)

install:
	cp bin/linux/$(exeName) ${GOPATH}/bin/

release: all
ifdef tagv
	mkdir -p bin/release/$(tagv)/windows; cp bin/windows/$(exeName).exe bin/release/$(tagv)/windows/
	mkdir -p bin/release/$(tagv)/linux; cp bin/linux/$(exeName) bin/release/$(tagv)/linux/ 
	cp README.MD bin/release/$(tagv)
	zip -r bin/release/$(exeName)_$(tagv).zip bin/release/$(tagv)
else
	@echo 'tagv not defined'
endif

dockerRun:
	docker run --env-file env.list --rm -it -p 9000:9000  $(dockerImageTag) 

dockerBuild:
	docker build -t $(dockerImageTag) .

clean:
	rm -f bin/linux/$(exeName) ; rm -f bin/windows/$(exeName).exe ; rm -rf bin/release
	