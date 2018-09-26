# base contains all the needed dependencies from dep
FROM rodrigodmd/ml-mutant-srv:base

ENV REPO_DIR $GOPATH/src/github.com/rodrigodmd/ml-mutant-srv
ENV APP_DIR $REPO_DIR/gateway

# Clean service dir
RUN rm -rf $APP_DIR && mkdir $APP_DIR

# Copy service source code
ADD ./ $APP_DIR

# Compile the binary 
RUN cd $APP_DIR && go build -o mutant-srv main.go

WORKDIR $APP_DIR
# Set the entrypoint
ENTRYPOINT [ "./mutant-srv" ] 
