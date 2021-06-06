## Angelo

TUI app to practice music.

## Alternatives

If I don't find an easy way to do it with Go, I could use a browser app instead. For now, trying to figure out how to change the speed without affecting pitch.

## Possible libraries for UI
- https://github.com/gdamore/tcell
- https://github.com/rivo/tview
= https://github.com/jroimartin/gocui

## Repos to watch
- https://github.com/skanehira/github-tui

## Using ffplay

### What I need from ffplay?
- [ ] Looping
  Use `-ss` to set an initial time.
  I can use `-t` to set how much time do I want to loop.
  Then use the `-loop` option with 0 to repeat forever.
- [ ] Speed change
- [ ] Start from a specific time

### Full example
``` bash
❯ ffplay -v warning -nodisp -autoexit -af "atempo=1" -ss 4 -t 5 -loop 0 samples/ton-doux-sourire.mp3
```

