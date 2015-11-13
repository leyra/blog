### Blog

I'll come up with a better name for this eventually. This is pretty much an
example web application using [Leyra](https://github.com/leyra/leyra). It's
super easy to get up and running provide you have:

  - Go (tested on 1.5)
  - Make
  - [Gresh](https://github.com/leyra/gresh) (optional)

[Gresh](https://github.com/leyra/gresh) is a cli tool I've been putting together
to assist in managing web applications I build with Leyra. It is under active
development but as of writing this should work fine for the purpose of playing
around with this application.

Note: Leyra (currently) needs a local Go environment. This has been done for a
few reasons. Firstly to keep only the dependencies we need in place and up to
date. Also to have clearer import paths within the application.

### Install

##### With Gresh

```bash
$ gresh fetch leyra/blog
$ export GOPATH=$(pwd)/blog
$ cd blog/src/leyra
$ make run
```

##### Without Gresh

```bash
$ mkdir blog blog/src
$ git clone git@github.com:leyra/blog.git blog/src/leyra
$ export GOPATH=$(pwd)/blog
$ cd blog/src/leyra
$ make run
```

The first time this is ran it may take a couple of minutes to initially get the
dependencies, but after that it will not bother to fetch them again. If all goes
well, you should be able to head over to [localhost:3000](http://localhost:3000)
to see the blog all up and running.

### License

MIT
