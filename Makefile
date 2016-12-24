FARIY_API_LDFLAGS = -ldflags
all: fairy-api 
init:
	mkdir -p output
	rm -rf output/*
	mkdir -p output/fairy-api

fairy-api: init
	echo "fairy-api is building..."
	#cd fairy-api/ && go build $(FARIY_API_LDFLAGS) -o ../output/fairy-api/fairy-api
	cd fairy-api/ && go build -o ../output/fairy-api/fairy-api

clean:
	rm -rf output
