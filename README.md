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

*create*

	% curl -X POST localhost:1982/obj -d '{
		"id"		 : "8",
		"resolution" : "1min"
	}'
	201 CREATED { "id":"8" }

*delete*

	% curl -X DELETE localhost:1982/obj/8
	200 OK
		
*define*

	% curl -X PUT localhost:1982/obj/8/_av -d '{
		 "from" 	: "2013-05-05T00:00:00",
		 "to" 		: "2013-06-05T00:00:00",
		 "value"	: 1
	}'
	
	% curl -X PUT localhost:1982/obj/8/_av -d '{
		 "from" 	: "2013-05-05T00:00:00",
		 "duration" : "7d",
		 "value"	: 1
	}'
	
*add*

	% curl -X PUT localhost:1982/obj/8/_ev -d '{
		 "id" 		: "17",
		 "from" 	: "2013-05-07T18:00:00",
		 "duration" : "2h"
	}'
	
*retrieve*

	% curl 'localhost:1982/obj/13/_av?from=2013-05-05&to=2013-05-06&resolution=hour'
	{
		"from" 		: "2013-05-05T00:00:00",
		"to" 		: "2013-05-06T00:00:00",
		resolution	: "hour",
		internal_resolution: "1min",
		defined:	: [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1],
		available:	: [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
	}