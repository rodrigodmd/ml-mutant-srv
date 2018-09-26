FROM library/golang

ENV REPO_DIR $GOPATH/src/github.com/rodrigodmd/ml-mutant-srv
RUN mkdir -p $REPO_DIR

ADD ./ $REPO_DIR


