# go-react-boilerplate
Demo application that shows how to use a react frontend with a Go backend.

### dependencies
- yarn
- npm
- go
- dep

#### On macOS (with brew)
```bash
brew install yarn npm go dep
```

### Quickstart
```bash
git clone https://github.com/sharath/go-react-boilerplate.git
dep ensure
cd client && yarn && cd ..
go run App.go
```

The server is running at [localhost:3000](http://localhost:3000/). You can see the API at [localhost:3000/api/v1/users](http://localhost:3000/api/v1/users)

### Todo
- [ ] Support for HTTPS
- [ ] MongoDB integration