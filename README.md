travl
=====

build:

	% go get "github.com/gorilla/mux"
	% go install ./...
	

run:

	% bin/travl
	
get help:

	% bit/travl -help
	

use:

	% curl -X PUT localhost:8080/res
	{
		id:1
	}
	% curl -X PUT localhost:8080/res/1 -d '{name:"bob"}'
	{	
		id:1, 
		name="bob"
	}
	% curl localhost:8080/res/1
	{	
		id:1, 
		name="bob"
	}
	% curl -X DEL localhost:8080/res/1
	{
		id:1
		delete: "sucess"
	}
	