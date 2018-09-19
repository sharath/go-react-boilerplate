# go-react-boilerplate
Demo application that shows how to use a react frontend with a Go backend.

### dependencies
- yarn
- npm
- go
- dep

### Quickstart
```bash
go get github.com/sharath/go-react-boilerplate
cd $GOPATH/src/github.com/sharath/go-react-boilerplate
dep ensure
cd client && yarn && yarn build
cd .. && go run App.go
```

# Project Structure
```bash
.
├── App.go
├── Gopkg.toml
├── README.md
├── backend
│   └── Backend.go
├── client
│   ├── package.json
│   ├── public
│   │   ├── favicon.ico
│   │   ├── index.html
│   │   └── manifest.json
│   ├── src
│   │   ├── App.css
│   │   ├── App.js
│   │   ├── App.test.js
│   │   ├── components
│   │   │   └── demo
│   │   │       └── demo.js
│   │   ├── index.css
│   │   ├── index.js
│   │   └── logo.svg
│   └── yarn.lock
└── frontend -> client/build
```

The server is running at [localhost:3000](http://localhost:3000/). You can see the API at [localhost:3000/api/v1/users](http://localhost:3000/api/v1/users)
