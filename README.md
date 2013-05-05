travl
=====

**build**:

	% go get "github.com/gorilla/mux"
	% go install ./...
	

**run**:

	% bin/travl
	
**get help**:

	% bit/travl -help
	

**use**:

	create
	% curl -X POST localhost:8080/res
	{
		id:1
	}

	delete
	% curl -X DEL localhost:8080/res/1
	{
		id:1
		delete: "sucess"
	}
	