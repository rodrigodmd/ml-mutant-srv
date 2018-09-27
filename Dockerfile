# This image is intended to be the base image containing all the dependencies 
# needed to build the binary inside the docker image
FROM library/golang

# Copy code into image
ENV REPO_DIR $GOPATH/src/github.com/rodrigodmd/ml-mutant-srv
RUN mkdir -p $REPO_DIR
ADD ./ $REPO_DIR

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN cd $REPO_DIR && dep ensure