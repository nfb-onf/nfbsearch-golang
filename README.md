## Go/Revel/Elasticsearch proof of concept

### Usage

    follow installation steps
    revel run nfbsearch
    visit /search?q=something+to+search

### Installation

#### Install Go

##### Mac OS X

    brew install go

##### Tarballs

    visit http://golang.org/doc/install#tarball

#### Install Dependencies

    go get github.com/robfig/revel
    go get github.com/robfig/revel/revel
    go get github.com/mattbaird/elastigo

#### Build

    revel build $GOPATH/nfbsearch ~/build/nfbsearch

### Test

Run revel development server:

    revel run nfbsearch
    visit /@tests