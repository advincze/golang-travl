travl
=====

**build**:

	// get sources
	% git clone git@github.com:advincze/golang-travl.git

	// setup gopath	
	% cd golang-travl
	% export GOPATH=`pwd`
	
	// get dependencies
	% go get "github.com/gorilla/mux"
	% go get "launchpad.net/gocheck"
	
	// build
	% go install ./...
	

**run**:

	% bin/travl
	
**get help**:

	% bin/travl -help
	

**use**:


*delete*

	% curl -X DELETE localhost:1982/bob
	200 OK
		
*define*

	% curl -X PUT localhost:1982/bob/_av -d '{
		 "from" 	: "2013-05-05T00:00:00Z",
		 "to" 		: "2013-06-05T00:00:00Z",
		 "value"	: 1
	}'
	
	% curl -X PUT localhost:1982/bob/_av -d '{
		 "from" 	: "2013-05-07T00:00:00Z",
		 "to" 		: "2013-05-09T00:00:00Z",
		 "value"	: 0
	}'
	


*add* //TODO

	% curl -X PUT localhost:1982/bob/_ev -d '{
		 "id" 		: "17",
		 "from" 	: "2013-05-07T18:00:00",
		 "duration" : "2h"
	}'
	

	
*retrieve*
	
	% curl 'localhost:1982/bob/_av?at=2013-05-05'
	{
		"at" 		: "2013-05-05T00:00:00",
		internal_resolution: "1min",
		defined:	: true
		available:	: true
	}
	
	% curl 'localhost:1982/bob/_av?from=2013-05-05&to=2013-05-06&resolution=hour'
	{
		"from" 		: "2013-05-05T00:00:00",
		"to" 		: "2013-05-06T00:00:00",
		resolution	: "hour",
		internal_resolution: "1min",
		available:	: [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
	}