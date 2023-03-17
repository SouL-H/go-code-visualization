# Go Code Visualization

## First Step

```go
go get -u github.com/ofabry/go-callvis
```

## Second Step

```go
export PATH=$PATH:$(go env GOPATH)/bin
```

## Third Step
```go
go-callvis <package_name>

exp:go-callvis go-visualization
```

> The application is listening on port 7878. The output is given as Svg format and you can click on the function you want and visualise it.

 - [For more details](github.com/ofabry/go-callvis)
 - [CLI Option](https://github.com/ofabry/go-callvis#options)

## Example

```go
go-callvis go-visualization
```

![Logo](/demo/demo.png)

