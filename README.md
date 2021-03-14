# blindtest-crop
[![build](https://github.com/aminekaabachi/blindtest-crop/actions/workflows/build.yaml/badge.svg)](https://github.com/aminekaabachi/blindtest-crop/)
[![test](https://github.com/aminekaabachi/blindtest-crop/actions/workflows/test.yaml/badge.svg)](https://github.com/aminekaabachi/blindtest-crop/)
[![codecov](https://codecov.io/gh/aminekaabachi/blindtest-crop/branch/main/graph/badge.svg?token=MS0JGLZOR8)](https://codecov.io/gh/aminekaabachi/blindtest-crop)

Utility to crop MP3 to first X seconds for blind test.

## Building the project

The `Makefile` contains all neccessary goals for build steps.

Simply type in `make help` to list all available goals.

## Running

Simply run the application using the mandatory parameter: `./blindtest-crop -input-path <MP3_FOLDER>`

The application then prints the following output:

```

```

## Arguments
```
  -input-path string
        MP3 songs input folder path.
  -output-path string
        MP3 croped songs output path. (default "./output")
  -seconds int
        Number of seconds to crop. (default 30)
```

## Building and running in Docker

Build the application using `make build-docker` and run the following command: `docker run org/blindtestcrop:latest --input-path <MP3_FOLDER>`