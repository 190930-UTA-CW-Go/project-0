makefile_dir		:= $(abspath $(shell pwd))
export

go_package  	:= $(shell cat go.mod | grep '^module' | sed 's/module //')
go_test 		:= go test -count=1 -v
service := project-0

run:
	go build
	./$(service)
	rm $(service)

shell:
	docker exec -it pgcontainer bash
	# psql -U postgres
	# \l , \c accounts, \dt 