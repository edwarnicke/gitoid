gitoid provides a simple library to compute gitoids (git object ids)

By default it produces gitoids for git object type blob using sha1:

```go
var reader os.Reader
gitoidHash, err := gitoid.New(reader)
fmt.Println(gitoidHash)
// Output: 261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64
fmt.Println(gitoidHash.URI())
// Output: gitoid:blob:sha1:261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64
```

but it can easily be used to compute gitoids using sha256:

```go
var reader os.Reader
gitoidHash, err := gitoid.New(reader, gitoid.WithSha256())
fmt.Println(gitoidHash)
// Output: ed43975fbdc3084195eb94723b5f6df44eeeed1cdda7db0c7121edf5d84569ab
fmt.Println(gitoidHash.URI())
// Output: gitoid:blob:sha256:ed43975fbdc3084195eb94723b5f6df44eeeed1cdda7db0c7121edf5d84569ab
```

or compute gitoids for another git object type:

```go
var reader os.Reader
gitoidHash, err := gitoid.New(reader, gitoid.WithGitObjectType(gitoid.COMMIT))
```

or assert a contentLength to be read from the reader (note: if contentLength bytes are unavailable gitoid.New will return an error):
```go
var reader os.Reader
var contentLength int64
gitoidHash, _ := gitoid.New(reader, gitoid.WithContentLength(contentLength))
fmt.Println(gitoidHash)
// Output: 261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64
```

if you have a string or an array of bytes it's simple to compute the gitoid by using bytes.NewBuffer:
```go
input := []byte("example")
	gitoidHash, _ := gitoid.New(bytes.NewBuffer(input))
	fmt.Println(gitoidHash)
	// Output: 96236f8158b12701d5e75c14fb876c4a0f31b963
	fmt.Println(gitoidHash.URI())
// Output: gitoid:blob:sha1:96236f8158b12701d5e75c14fb876c4a0f31b963
```


