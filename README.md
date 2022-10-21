## My Text Adventure

Just doing this for fun. Seeing where it goes...Does anyone play these anymore, anyway?

### Install

Couple ways to install. If your go env is set up properly, all you need to do is

```shell
$ go install github.com/unclassedpenguin/textgamev2@latest
```

And then run 

```shell
$ textgamev2
```

Or If you want to do it manually:

- Clone this repository.
- cd into folder
- run: 

```shell
$ go build main.go
```
- Then just execute!

```shell
$ ./main
```

### To-do:

- Intro
- ~~Add the rope item~~
- Make inventory a little prettier instead of just printing out an array. 
  - Have it say something if inventory is empty. Ie "Your inventory appears to be empty" or "Doesn't seem to be anything in your inventory"
- Continue work on area descriptions. Could be more descriptive/better written.
- Need to come up with a mini game for fighting the monster...skill, or luck?
  - Its luck at the moment. But not bad...
- ~~Add inventory to an area. So if you pick up the axe in one area and drop it in another, it is in that area.~~ 
  - ~~Turns out this is slightly more complicated than I thought....Not sure at the moment how to implement it.~~
  - (I was able to do this. Its maybe a bit more complicated than it needed to be. With a whole bunch of global variables. But it works...)
- ~~Change "please enter a valid option"~~
  - ~~maybe "Please try another option, or 'help' if you're not sure"~~
- Go through and format(clear unnecessary whitespace etc.) and comment code.
