# :crystal_ball: fortune-api :crystal_ball:

[![forthebadge](https://forthebadge.com/images/badges/60-percent-of-the-time-works-every-time.svg)](https://forthebadge.com)

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e4942f64299a4a84b90eb7631ac26d5a)](https://www.codacy.com/app/sarah256/fortune-api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=sarah256/fortune-api&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/sarah256/fortune-api)](https://goreportcard.com/report/github.com/sarah256/fortune-api)

## Description
A RESTful API for the classic fortune command-line utility (in the fortune-mod package). Every query will return a random fortune!
These fortunes are originally from the [fortune-mod](https://github.com/shlomif/fortune-mod) repository, from the original command line utility.

It is hosted at [https://fortuneapi.herokuapp.com/](https://fortuneapi.herokuapp.com/).

## Adding Fortunes
To add a new fortune, go to the `datfiles` directory, and choose the correspoding file to contribute to.  If none of the files fit, feel free to make your own!  As with the original, fortunes deemed to be not-so-wholesome should be confined to the `datfiles/off` directory.

Simply type your fortune, formatted as you wish, and then follow it with a newline containing a `%`.  That's all!

## Running Locally
To run the API locally, ensure that you have Go 1.7+ installed.

The only dependency needed is mux, which can be installed by running:
```bash
$ go get github.com/gorilla/mux
```

In the repository folder, use the following to run the API server:
```bash
$ go run main.go
```

The API is hosted at [https://fortuneapi.herokuapp.com/](https://fortuneapi.herokuapp.com/), but you can change this in the `main()` function in the `main.go` file to run it locally.

I recommend changing it as follows:

Replace these lines in `main()`:
```go
	port := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, router))
```

With this line:
```go
    log.Fatal(http.ListenAndServe(":8080", router))
```
Now you can access the API at `localhost:8080`!

## Endpoints
No additional path is needed if you would like a completely random fortune.

You can get a fortune from a specific genre at `https://fortuneapi.herokuapp.com/<genre>`, replacing 'genre' with the desired datfile name.  The genre must be the exact name of one of the files in the `datfiles` directory.
For example, `https://fortuneapi.herokuapp.com/computers` *could* give you an output such as:
```
"We are experiencing system trouble -- do not adjust your terminal.\n"
```
*Note: the output is always random, so even with the same URL, the output will likely vary with every query*
