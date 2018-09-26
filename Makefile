dockerbase:
	dep ensure
	docker build -t rodrigodmd/ml-mutant-srv:base .
	docker push rodrigodmd/ml-mutant-srv:base