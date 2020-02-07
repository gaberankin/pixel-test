# This is me, messing around with a game library in go

I'd tried to mess around with [ebiten](https://github.com/hajimehoshi/ebiten), but found that drawing primitives isn't very straightforward out of the box for it (although it seems pretty rad, so I'll definitely revisit it later).  I switched to messing around with [pixel](https://github.com/faiface/pixel) because it supports drawing primitives fairly easily with its imdraw api.

I'm shooting to make this an easy build & run, although it relies on go modules, so your environment will need to support that if you want to play with this as well.

The included `Makefile` is just a convenience, so if your system doesn't support make you can just run `go build`.  I prefer to put my binaries, while developing, in a place out of the way so the `Makefile` is set up to drop the binary in a `bin/` directory.  If you're good with not worrying about that, go for it.

Right now the idea is to make a geometry-wars-like game with very primitive shapes and graphics.  Ultimately this will end up being a mind-dump as I tool around and mess with things.