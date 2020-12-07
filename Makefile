CXX = go

SRC = ./tests

all: clean test

test:
	$(CXX) test $(SRC) -v

clean:
	$(CXX) clean -testcache