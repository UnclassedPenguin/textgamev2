//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------
//
// Tyler(UnclassedPenguin) Text Adventure 2022
//
// Author: Tyler(UnclassedPenguin)
// GitHub: https://github.com/UnclassedPenguin/textgame/
//
//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------


package main

import (
  "fmt"
  "os"
  "flag"
  "strconv"
  tge "github.com/unclassedpenguin/textgameengine"
  "golang.org/x/term"
)


//-----------------------------------------------------------------------------
// Functions start
//-----------------------------------------------------------------------------

// Prints the help function when you type help or ?
func help() {
  fmt.Println("--------------------------------------------------------------------------------")
  fmt.Println("Help: ")
  fmt.Println("To move a direction, simply type the direction you want to go.")
  fmt.Println("       i.e.(north, south, east, or west)")
  fmt.Println("If there is an item, just try typing its name to pick it up.")
  fmt.Println("Type 'inv' to check whats in your inventory.")
  fmt.Println("Type 'look' to check your surroundings.")
  fmt.Println("Type 'exit' to exit the game.")
  fmt.Println("--------------------------------------------------------------------------------")
}

//I'd like to make a few phrases here and randomly pick one to say, just
//for some variety.
func cantGo(termWidth int) {
  str := []string{
    "I'm sorry, that way is blocked.",
    "I'm sorry, you can't go that way.",
    "Doesn't seem to be a path in that direction.",
    "It's not possible to go that way.",
  }
  s()
  tge.PrintSlow(str[tge.RandNumber(len(str))], termWidth, slowMode)
}

//s for give me some (s)pace
func s() {
  fmt.Print("\n")
}

func exit(i int) {
  s()
  fmt.Println("Thanks For Playing!")
  os.Exit(i)
}

//-----------------------------------------------------------------------------
// Functions end
//-----------------------------------------------------------------------------


//-----------------------------------------------------------------------------
// Areas start
//-----------------------------------------------------------------------------

func intro(termWidth int) string{
  var name string
  s()
  fmt.Println("Welcome to UnclassedPenguin TextAdventure!")
  fmt.Println("Author: Tyler(UnclassedPenguin)")
  fmt.Println("Github: https://github.com/UnclassedPenguin/textgamev2/")
  s()
  s()
  tge.PrintSlow("What's your name?", termWidth, slowMode)
  fmt.Print(" > ")
  fmt.Scan(&name)
  if name == "exit" {
    exit(0)
  }
  s()
  tge.PrintSlow("Hope you enjoy the game, " + name + ". Good luck!\nIf you get stuck, try 'help'.", termWidth, slowMode)
  s()
  fmt.Println("--------------------------------------------------------------------------------")
  s()
  tge.PrintSlow("You wake up and become aware of your surroundings. You're not quite sure what happened, and your head feels a bit cloudy. Your eyes slowly start to open........", termWidth, slowMode)
  return name
}


func startArea(player tge.Player, game tge.Game) {
  // validDirections = south, west

  var userchoice string
  var description string
  var baseDescription string
  var axeDescription string
  var swordDescription string
  var ropeDescription string
  var directions string

  baseDescription = "You find yourself in the middle of a forest. The trees surrounding you are tall and the canopy is thick, blocking nearly all the sunlight from coming through."

  if startAxe {
    axeDescription = " You see an axe leaning up against a tree."
  } else {
    axeDescription = ""
  }

  if startSword {
    swordDescription = " You see a sword lying on the ground."
  } else {
    swordDescription = ""
  }

  if startRope {
    ropeDescription = " You see a rope lying on the ground."
  } else {
    ropeDescription = ""
  }

  directions = "You can go south or west."

  description = baseDescription + axeDescription + swordDescription + ropeDescription + "\n" + directions

  tge.PrintSlow(description, game.TermWidth, slowMode)

  for true {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      cantGo(game.TermWidth)
    } else if userchoice == "east" {
      cantGo(game.TermWidth)
    } else if userchoice == "south" {
      s()
      tge.PrintSlow("You go south.", game.TermWidth, slowMode)
      sArea(player, game)
    } else if userchoice == "west" {
      s()
      tge.PrintSlow("You go west.", game.TermWidth, slowMode)
      wArea(player, game)
    } else if userchoice == "axe" {
      s()
      if tge.InventoryContains(&player.Inventory, "axe") {
        startAxe = true
        tge.PrintSlow("You drop the axe.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "axe")
      } else if startAxe {
        startAxe = false
        tge.PrintSlow("You pick up the axe. It's a nice heavy American felling axe.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "axe")
      } else {
        tge.PrintSlow("What axe?", game.TermWidth, slowMode)
      }
    } else if userchoice == "sword" {
      s()
      if tge.InventoryContains(&player.Inventory, "sword") {
        startSword = true
        tge.PrintSlow("You drop the sword.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "sword")
      } else if startSword {
        startSword = false
        tge.PrintSlow("You pick up the sword. It's a long sword.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "sword")
      } else {
        tge.PrintSlow("What sword?", game.TermWidth, slowMode)
      }
    } else if userchoice == "rope" {
      s()
      if tge.InventoryContains(&player.Inventory, "rope") {
        startRope = true
        tge.PrintSlow("You drop the rope.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "rope")
      } else if startRope {
        startRope = false
        tge.PrintSlow("You pick up the rope. It looks heavy enough to hold your weight.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "rope")
      } else {
        tge.PrintSlow("What rope?", game.TermWidth, slowMode)
      }
    //} else if userchoice == "?" {
      //checkLocalItems(startAxe, startSword, startRope)
    } else if userchoice == "look" {
      if startAxe {
        axeDescription = "You see an axe leaning up against a tree."
      } else {
        axeDescription = ""
      }

      if startSword {
        swordDescription = "You see a sword lying on the ground."
      } else {
        swordDescription = ""
      }

      if startRope {
        ropeDescription = "You see a rope lying on the ground."
      } else {
        ropeDescription = ""
      }

      directions := "You can go south or west."

      description := baseDescription + axeDescription + swordDescription + ropeDescription + "\n" + directions

      s()
      tge.PrintSlow(description, game.TermWidth, slowMode)
    } else if userchoice == "inv" {
      s()
      tge.PrintInventory(&player.Inventory)
      s()
    } else if userchoice == "help" || userchoice == "?" {
      s()
      help()
    } else if userchoice == "exit" {
      exit(0)
    } else {
      s()
      tge.PrintSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter another option, or try 'help'.\n", game.TermWidth, slowMode)
    }
  }
}


func wArea(player tge.Player, game tge.Game) {
  // validDirections = north, east, south

  var userchoice string
  var description string
  var baseDescription string
  var eventLog string
  var axeDescription string
  var swordDescription string
  var ropeDescription string
  var directions string

  baseDescription = "There is a little clearing in the trees here with a small pond, fed by a natural spring, which has a stream leading out of it to the south."

  if game.Events["log"] {
    eventLog = " To the north it looks like there is a path, but with a large log blocking the way."
    directions = "\nYou can go east or south."
  } else {
    eventLog = " To the north there is a path you cleared, with a large log split in half on either side."
    directions = "\nYou can go north, east, or south."
  }

  if wAxe {
    axeDescription = " You see an axe lying on the ground."
  } else {
    axeDescription = ""
  }

  if wSword {
    swordDescription = " You see a sword lying on the ground."
  } else {
    swordDescription = ""
  }

  if wRope {
    ropeDescription = " You see a rope lying on the ground."
  } else {
    ropeDescription = ""
  }

  description = baseDescription + eventLog + axeDescription + swordDescription + ropeDescription + directions

  s()
  tge.PrintSlow(description, game.TermWidth, slowMode)

  for true {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
        // if user has axe and log is still there
      if tge.InventoryContains(&player.Inventory, "axe") && game.Events["log"] == true {
        tge.PrintSlow("You use your axe to clear the log and travel north.", game.TermWidth, slowMode)
        game.Events["log"] = false
        nwArea(player, game)
        // if user has axe and log is not there
      } else if tge.InventoryContains(&player.Inventory, "axe") && game.Events["log"] == false {
        tge.PrintSlow("You travel north.", game.TermWidth, slowMode)
        nwArea(player, game)
        // if user has already cleared the log, dropped the axe back in startArea 
        // and comes back. So log not there, and doesn't have axe.
      } else if game.Events["log"] == false {
        tge.PrintSlow("You travel north.", game.TermWidth, slowMode)
        nwArea(player, game)
        // if user doesn't have axe and the log is still there
      } else {
        tge.PrintSlow("There is a log blocking the way! If only you had a way to clear it...", game.TermWidth, slowMode)
      }
    } else if userchoice == "east" {
      s()
      tge.PrintSlow("You go east.", game.TermWidth, slowMode)
      startArea(player, game)
    } else if userchoice == "south" {
      s()
      tge.PrintSlow("You go south.", game.TermWidth, slowMode)
      swArea(player, game)
    } else if userchoice == "west" {
      cantGo(game.TermWidth)
    } else if userchoice == "pond" {
      s()
      tge.PrintSlow("You look in the pond. There are some small fish swimming around.", game.TermWidth, slowMode)
    } else if userchoice == "fish" {
      s()
      tge.PrintSlow("You say hi to the fish in the pond, but they don't seem interested in being friends.", game.TermWidth, slowMode)
    } else if userchoice == "axe" {
      s()
      if tge.InventoryContains(&player.Inventory, "axe") {
        wAxe = true
        tge.PrintSlow("You drop the axe.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "axe")
      } else if wAxe {
        wAxe = false
        tge.PrintSlow("You pick up the axe. It's a nice heavy American felling axe.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "axe")
      } else {
        tge.PrintSlow("What axe?", game.TermWidth, slowMode)
      }
    } else if userchoice == "sword" {
      s()
      if tge.InventoryContains(&player.Inventory, "sword") {
        wSword = true
        tge.PrintSlow("You drop the sword.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "sword")
      } else if wSword {
        wSword = false
        tge.PrintSlow("You pick up the sword. It's a long sword.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "sword")
      } else {
        tge.PrintSlow("What sword?", game.TermWidth, slowMode)
      }
    } else if userchoice == "rope" {
      s()
      if tge.InventoryContains(&player.Inventory, "rope") {
        wRope = true
        tge.PrintSlow("You drop the rope.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "rope")
      } else if wRope {
        wRope = false
        tge.PrintSlow("You pick up the rope. It looks heavy enough to hold your weight.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "rope")
      } else {
        tge.PrintSlow("What rope?", game.TermWidth, slowMode)
      }
    //} else if userchoice == "?" {
      //checkLocalItems(wAxe, wSword, wRope)
    } else if userchoice == "look" {
      if game.Events["log"] {
        eventLog = " To the north it looks like there is a path, but with a large log blocking the way."
      } else {
        eventLog = " To the north there is a path you cleared, with a large log split in half on either side."
      }

      if wAxe {
        axeDescription = " You see an axe lying on the ground."
      } else {
        axeDescription = ""
      }

      if wSword {
        swordDescription = " You see a sword lying on the ground."
      } else {
        swordDescription = ""
      }

      if wRope {
        ropeDescription = " You see a rope lying on the ground."
      } else {
        ropeDescription = ""
      }

      directions = "\nYou can go north, east, or south."

      description = baseDescription + eventLog + axeDescription + swordDescription + ropeDescription + directions

      s()
      tge.PrintSlow(description, game.TermWidth, slowMode)
    } else if userchoice == "inv" {
      s()
      tge.PrintInventory(&player.Inventory)
      s()
    } else if userchoice == "help" || userchoice == "?" {
      s()
      help()
    } else if userchoice == "exit" {
      exit(0)
    } else {
      s()
      tge.PrintSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter another option, or try 'help'.\n", game.TermWidth, slowMode)
    }
  }
}


func nwArea(player tge.Player, game tge.Game) {
  // validDirections = south, west

  var userchoice string
  var description string
  var baseDescription string
  var axeDescription string
  var swordDescription string
  var ropeDescription string
  var directions string

  baseDescription = "There are tall trees all around you. The sun gleams through a few of the trees."

  if nwAxe {
    axeDescription = " You see an axe lying on the ground."
  } else {
    axeDescription = ""
  }

  if nwSword {
    swordDescription = " Is that something shiny behind that tree? It almost looks like it could be a sword..."
  } else {
    swordDescription = ""
  }

  if nwRope {
    ropeDescription = " You see a rope lying on the ground."
  } else {
    ropeDescription = ""
  }

  directions = "\nYou can only go south."

  description = baseDescription + swordDescription + axeDescription + ropeDescription + directions

  s()
  tge.PrintSlow(description, game.TermWidth, slowMode)

  for true {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      cantGo(game.TermWidth)
    } else if userchoice == "east" {
      cantGo(game.TermWidth)
    } else if userchoice == "south" {
      s()
      tge.PrintSlow("You go south.", game.TermWidth, slowMode)
      wArea(player, game)
    } else if userchoice == "west" {
      cantGo(game.TermWidth)
    } else if userchoice == "axe" {
      s()
      if tge.InventoryContains(&player.Inventory, "axe") {
        nwAxe = true
        tge.PrintSlow("You drop the axe.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "axe")
      } else if nwAxe {
        nwAxe = false
        tge.PrintSlow("You pick up the axe. It's a nice heavy American felling axe.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "axe")
      } else {
        tge.PrintSlow("What axe?", game.TermWidth, slowMode)
      }
    } else if userchoice == "sword" {
      s()
      if tge.InventoryContains(&player.Inventory, "sword") {
        nwSword = true
        tge.PrintSlow("You drop the sword.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "sword")
      } else if nwSword {
        nwSword = false
        tge.PrintSlow("You pick up the sword. It's a long sword.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "sword")
      } else {
        tge.PrintSlow("What sword?", game.TermWidth, slowMode)
      }
    } else if userchoice == "rope" {
      s()
      if tge.InventoryContains(&player.Inventory, "rope") {
        nwRope = true
        tge.PrintSlow("You drop the rope.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "rope")
      } else if nwRope {
        nwRope = false
        tge.PrintSlow("You pick up the rope. It looks heavy enough to hold your weight.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "rope")
      } else {
        tge.PrintSlow("What rope?", game.TermWidth, slowMode)
      }
    //} else if userchoice == "?" {
      //checkLocalItems(nwAxe, nwSword, nwRope)
    } else if userchoice == "look" {
      if nwAxe {
        axeDescription = " You see an axe lying on the ground."
      } else {
        axeDescription = ""
      }

      if nwSword {
        swordDescription = " Is that something shiny behind that tree? It almost looks like it could be a sword..."
      } else {
        swordDescription = ""
      }

      if nwRope {
        ropeDescription = " You see a rope lying on the ground."
      } else {
        ropeDescription = ""
      }

      directions = "\nYou can only go south."

      description = baseDescription + swordDescription + axeDescription + ropeDescription + directions

      s()
      tge.PrintSlow(description, game.TermWidth, slowMode)
    } else if userchoice == "inv" {
      s()
      tge.PrintInventory(&player.Inventory)
      s()
    } else if userchoice == "help" || userchoice == "?" {
      s()
      help()
    } else if userchoice == "exit" {
      exit(0)
    } else {
      s()
      tge.PrintSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter another option, or try 'help'.\n", game.TermWidth, slowMode)
    }
  }
}


func swArea(player tge.Player, game tge.Game) {
  // validDirections = north, east

  var userchoice string
  var description string
  var baseDescription string
  var axeDescription string
  var swordDescription string
  var ropeDescription string
  var directions string

  baseDescription = "The forest has cleared here, and there is a stream running from the north. The banks of the stream are covered in rocks. Don't slip!"

  if swAxe {
    axeDescription = " You see an axe lying on the ground."
  } else {
    axeDescription = ""
  }

  if swSword {
    swordDescription = " You see a sword lying on the ground."
  } else {
    swordDescription = ""
  }

  if swRope {
    ropeDescription = " You see a rope lying on the ground."
  } else {
    ropeDescription = ""
  }

  directions = "\nYou can go north or east."

  description = baseDescription + axeDescription + swordDescription + ropeDescription + directions

  s()
  tge.PrintSlow(description, game.TermWidth, slowMode)

  for true {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      tge.PrintSlow("You go north.", game.TermWidth, slowMode)
      wArea(player, game)
    } else if userchoice == "east" {
      s()
      tge.PrintSlow("You go east.", game.TermWidth, slowMode)
      sArea(player, game)
    } else if userchoice == "south" {
      cantGo(game.TermWidth)
    } else if userchoice == "west" {
      cantGo(game.TermWidth)
    } else if userchoice == "stream" {
      s()
      tge.PrintSlow("You bend down and use your hands to cup some water and drink it.", game.TermWidth, slowMode)
    } else if userchoice == "axe" {
      s()
      if tge.InventoryContains(&player.Inventory, "axe") {
        swAxe = true
        tge.PrintSlow("You drop the axe.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "axe")
      } else if swAxe {
        swAxe = false
        tge.PrintSlow("You pick up the axe. It's a nice heavy American felling axe.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "axe")
      } else {
        tge.PrintSlow("What axe?", game.TermWidth, slowMode)
      }
    } else if userchoice == "sword" {
      s()
      if tge.InventoryContains(&player.Inventory, "sword") {
        swSword = true
        tge.PrintSlow("You drop the sword.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "sword")
      } else if swSword {
        swSword = false
        tge.PrintSlow("You pick up the sword. It's a long sword.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "sword")
      } else {
        tge.PrintSlow("What sword?", game.TermWidth, slowMode)
      }
    } else if userchoice == "rope" {
      s()
      if tge.InventoryContains(&player.Inventory, "rope") {
        swRope = true
        tge.PrintSlow("You drop the rope.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "rope")
      } else if swRope {
        swRope = false
        tge.PrintSlow("You pick up the rope. It looks heavy enough to hold your weight.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "rope")
      } else {
        tge.PrintSlow("What rope?", game.TermWidth, slowMode)
      }
    //} else if userchoice == "?" {
      //checkLocalItems(swAxe, swSword, swRope)
    } else if userchoice == "look" {
      if swAxe {
        axeDescription = " You see an axe lying on the ground."
      } else {
        axeDescription = ""
      }

      if swSword {
        swordDescription = " You see a sword lying on the ground."
      } else {
        swordDescription = ""
      }

      if swRope {
        ropeDescription = " You see a rope lying on the ground."
      } else {
        ropeDescription = ""
      }

      directions = "\nYou can go north or east."

      description = baseDescription + axeDescription + swordDescription + ropeDescription + directions

      s()
      tge.PrintSlow(description, game.TermWidth, slowMode)
    } else if userchoice == "inv" {
      s()
      tge.PrintInventory(&player.Inventory)
      s()
    } else if userchoice == "help" || userchoice == "?" {
      s()
      help()
    } else if userchoice == "exit" {
      exit(0)
    } else {
      s()
      tge.PrintSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter another option, or try 'help'.\n", game.TermWidth, slowMode)
    }
  }
}


func sArea(player tge.Player, game tge.Game) {
  // validDirections = north, east, west

  var userchoice string
  var description string
  var baseDescription string
  var axeDescription string
  var swordDescription string
  var ropeDescription string
  var directions string

  baseDescription = "You find yourself in open grasslands. The land starts to rise gently to the east, with mountains in the distance."

  if sAxe {
    axeDescription = " You see an axe lying on the ground."
  } else {
    axeDescription = ""
  }

  if sSword {
    swordDescription = " You see a sword lying on the ground."
  } else {
    swordDescription = ""
  }

  if sRope {
    ropeDescription = " It looks like there is a rope laying in the grass."
  } else {
    ropeDescription = ""
  }

  directions = "\nYou can go north, east, or west."

  description = baseDescription + ropeDescription + axeDescription + swordDescription + directions

  s()
  tge.PrintSlow(description, game.TermWidth, slowMode)

  for true {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      tge.PrintSlow("You go north.", game.TermWidth, slowMode)
      startArea(player, game)
    } else if userchoice == "east" {
      s()
      tge.PrintSlow("You go east.", game.TermWidth, slowMode)
      seArea(player, game)
    } else if userchoice == "south" {
      cantGo(game.TermWidth)
    } else if userchoice == "west" {
      s()
      tge.PrintSlow("You go west.", game.TermWidth, slowMode)
      swArea(player, game)
    } else if userchoice == "axe" {
      s()
      if tge.InventoryContains(&player.Inventory, "axe") {
        sAxe = true
        tge.PrintSlow("You drop the axe.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "axe")
      } else if sAxe {
        sAxe = false
        tge.PrintSlow("You pick up the axe. It's a nice heavy American felling axe.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "axe")
      } else {
        tge.PrintSlow("What axe?", game.TermWidth, slowMode)
      }
    } else if userchoice == "sword" {
      s()
      if tge.InventoryContains(&player.Inventory, "sword") {
        sSword = true
        tge.PrintSlow("You drop the sword.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "sword")
      } else if sSword {
        sSword = false
        tge.PrintSlow("You pick up the sword. It's a long sword.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "sword")
      } else {
        tge.PrintSlow("What sword?", game.TermWidth, slowMode)
      }
    } else if userchoice == "rope" {
      s()
      if tge.InventoryContains(&player.Inventory, "rope") {
        sRope = true
        tge.PrintSlow("You drop the rope.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "rope")
      } else if sRope {
        sRope = false
        tge.PrintSlow("You pick up the rope. It looks heavy enough to hold your weight.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "rope")
      } else {
        tge.PrintSlow("What rope?", game.TermWidth, slowMode)
      }
    //} else if userchoice == "?" {
      //checkLocalItems(sAxe, sSword, sRope)
    } else if userchoice == "look" {
      if sAxe {
        axeDescription = " You see an axe lying on the ground."
      } else {
        axeDescription = ""
      }

      if sSword {
        swordDescription = " You see a sword lying on the ground."
      } else {
        swordDescription = ""
      }

      if sRope {
        ropeDescription = " It looks like there is a rope laying in the grass."
      } else {
        ropeDescription = ""
      }

      directions = "\nYou can go north, east, or west."

      description = baseDescription + ropeDescription + axeDescription + swordDescription + directions

      s()
      tge.PrintSlow(description, game.TermWidth, slowMode)
    } else if userchoice == "inv" {
      s()
      tge.PrintInventory(&player.Inventory)
      s()
    } else if userchoice == "help" || userchoice == "?" {
      s()
      help()
    } else if userchoice == "exit" {
      exit(0)
    } else {
      s()
      tge.PrintSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter another option, or try 'help'.\n", game.TermWidth, slowMode)
    }
  }
}


func seArea(player tge.Player, game tge.Game) {
  // validDirections = north, west

  // count for attempts at climbing cliff without rope.
  count := 0

  var userchoice string
  var description string
  var baseDescription string
  var axeDescription string
  var swordDescription string
  var ropeDescription string
  var directions string

  baseDescription = "The terrain has turned mountainous. There is a cliff to the north. You *might* be able to climb it..."

  if seAxe {
    axeDescription = " You see an axe lying on the ground."
  } else {
    axeDescription = ""
  }

  if seSword {
    swordDescription = " You see a sword lying on the ground."
  } else {
    swordDescription = ""
  }

  if seRope {
    ropeDescription = " You see a rope lying on the ground."
  } else {
    ropeDescription = ""
  }

  directions = "\nYou can go west."

  description = baseDescription + axeDescription + swordDescription + ropeDescription + directions

  s()
  tge.PrintSlow(description, game.TermWidth, slowMode)

  for true {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" || userchoice == "climb" {
      s()
      // If you have the rope, you are guaranteed to climb the cliff. 
      if tge.InventoryContains(&player.Inventory, "rope") {
        tge.PrintSlow("You use the rope to climb the cliff.", game.TermWidth, slowMode)
        eArea(player, game)
      // This is if you don't have a rope. Gives you a small change of
      // climbing the cliff. Currently 1/20 chance. Too small? Too big? 
      } else {
        if count < 5 {
          tge.PrintSlow("You decide to try free climbing the cliff...", game.TermWidth, slowMode)
          s()
          rn := tge.RandNumber(20)
          switch rn {
            case 7:
              tge.PrintSlow("You used your skill to successfully climb the cliff!", game.TermWidth, slowMode)
              eArea(player, game)
            default:
              count += 1
              tge.PrintSlow("You failed to climb the cliff, and fell to the bottom! Ouch!", game.TermWidth, slowMode)
          }
        } else {
          tge.PrintSlow("You're a little beat up from failed attempts. Maybe take a break, and try again later.", game.TermWidth, slowMode)
        }
      }
    } else if userchoice == "east" {
      cantGo(game.TermWidth)
    } else if userchoice == "south" {
      cantGo(game.TermWidth)
    } else if userchoice == "west" {
      s()
      tge.PrintSlow("You go west.", game.TermWidth, slowMode)
      sArea(player, game)
    } else if userchoice == "mountain" {
      s()
      tge.PrintSlow("To the east and south are mountains for miles.", game.TermWidth, slowMode)
    } else if userchoice == "axe" {
      s()
      if tge.InventoryContains(&player.Inventory, "axe") {
        seAxe = true
        tge.PrintSlow("You drop the axe.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "axe")
      } else if seAxe {
        seAxe = false
        tge.PrintSlow("You pick up the axe. It's a nice heavy American felling axe.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "axe")
      } else {
        tge.PrintSlow("What axe?", game.TermWidth, slowMode)
      }
    } else if userchoice == "sword" {
      s()
      if tge.InventoryContains(&player.Inventory, "sword") {
        seSword = true
        tge.PrintSlow("You drop the sword.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "sword")
      } else if seSword {
        seSword = false
        tge.PrintSlow("You pick up the sword. It's a long sword.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "sword")
      } else {
        tge.PrintSlow("What sword?", game.TermWidth, slowMode)
      }
    } else if userchoice == "rope" {
      s()
      if tge.InventoryContains(&player.Inventory, "rope") {
        seRope = true
        tge.PrintSlow("You drop the rope.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "rope")
      } else if seRope {
        seRope = false
        tge.PrintSlow("You pick up the rope. It looks heavy enough to hold your weight.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "rope")
      } else {
        tge.PrintSlow("What rope?", game.TermWidth, slowMode)
      }
    //} else if userchoice == "?" {
      //checkLocalItems(seAxe, seSword, seRope)
    } else if userchoice == "look" {
      if seAxe {
        axeDescription = " You see an axe lying on the ground."
      } else {
        axeDescription = ""
      }

      if seSword {
        swordDescription = " You see a sword lying on the ground."
      } else {
        swordDescription = ""
      }

      if seRope {
        ropeDescription = " You see a rope lying on the ground."
      } else {
        ropeDescription = ""
      }

      directions = "\nYou can go west."

      description = baseDescription + axeDescription + swordDescription + ropeDescription + directions

      s()
      tge.PrintSlow(description, game.TermWidth, slowMode)
    } else if userchoice == "inv" {
      s()
      tge.PrintInventory(&player.Inventory)
      s()
    } else if userchoice == "help" || userchoice == "?" {
      s()
      help()
    } else if userchoice == "exit" {
      exit(0)
    } else {
      s()
      tge.PrintSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter another option, or try 'help'.\n", game.TermWidth, slowMode)
    }
  }
}


func eArea(player tge.Player, game tge.Game) {
  // validDirections = north, south

  var userchoice string
  var description string
  var baseDescription string
  var axeDescription string
  var swordDescription string
  var ropeDescription string
  var directions string

  baseDescription = "You are on a high plateau. To the south is a cliff, and further south than that you can see large mountain ranges in the distance."

  if eAxe {
    axeDescription = " You see an axe lying on the ground."
  } else {
    axeDescription = ""
  }

  if eSword {
    swordDescription = " You see a sword lying on the ground."
  } else {
    swordDescription = ""
  }

  if eRope {
    ropeDescription = " You see a rope lying on the ground."
  } else {
    ropeDescription = ""
  }

  directions = "\nYou can go north or south"

  description = baseDescription + axeDescription + swordDescription + ropeDescription + directions

  s()
  tge.PrintSlow(description, game.TermWidth, slowMode)

  for true {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      tge.PrintSlow("You go north.", game.TermWidth, slowMode)
      neArea(player, game)
    } else if userchoice == "east" {
      cantGo(game.TermWidth)
    } else if userchoice == "south" {
      s()
      tge.PrintSlow("You go south.", game.TermWidth, slowMode)
      seArea(player, game)
    } else if userchoice == "west" {
      cantGo(game.TermWidth)
    } else if userchoice == "axe" {
      s()
      if tge.InventoryContains(&player.Inventory, "axe") {
        eAxe = true
        tge.PrintSlow("You drop the axe.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "axe")
      } else if eAxe {
        eAxe = false
        tge.PrintSlow("You pick up the axe. It's a nice heavy American felling axe.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "axe")
      } else {
        tge.PrintSlow("What axe?", game.TermWidth, slowMode)
      }
    } else if userchoice == "sword" {
      s()
      if tge.InventoryContains(&player.Inventory, "sword") {
        eSword = true
        tge.PrintSlow("You drop the sword.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "sword")
      } else if eSword {
        eSword = false
        tge.PrintSlow("You pick up the sword. It's a long sword.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "sword")
      } else {
        tge.PrintSlow("What sword?", game.TermWidth, slowMode)
      }
    } else if userchoice == "rope" {
      s()
      if tge.InventoryContains(&player.Inventory, "rope") {
        eRope = true
        tge.PrintSlow("You drop the rope.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "rope")
      } else if eRope {
        eRope = false
        tge.PrintSlow("You pick up the rope. It looks heavy enough to hold your weight.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "rope")
      } else {
        tge.PrintSlow("What rope?", game.TermWidth, slowMode)
      }
    //} else if userchoice == "?" {
      //checkLocalItems(eAxe, eSword, eRope)
    } else if userchoice == "look" {

      if eAxe {
        axeDescription = " You see an axe lying on the ground."
      } else {
        axeDescription = ""
      }

      if eSword {
        swordDescription = " You see a sword lying on the ground."
      } else {
        swordDescription = ""
      }

      if eRope {
        ropeDescription = " You see a rope lying on the ground."
      } else {
        ropeDescription = ""
      }

      directions = "\nYou can go north or south"

      description = baseDescription + axeDescription + swordDescription + ropeDescription + directions

      s()
      tge.PrintSlow(description, game.TermWidth, slowMode)
    } else if userchoice == "inv" {
      s()
      tge.PrintInventory(&player.Inventory)
      s()
    } else if userchoice == "help" || userchoice == "?" {
      s()
      help()
    } else if userchoice == "exit" {
      exit(0)
    } else {
      s()
      tge.PrintSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter another option, or try 'help'.\n", game.TermWidth, slowMode)
    }
  }
}


func neArea(player tge.Player, game tge.Game) {
  // validDirections = south, west

  var userchoice string
  var description string
  var baseDescription string
  var eventMonster string
  var axeDescription string
  var swordDescription string
  var ropeDescription string
  var directions string

  baseDescription = "You enter another forest area. Pines are surrounding you on all sides. There is a path to the west."

  if game.Events["monster"] {
    eventMonster = " When you look closer you see there is a monster standing there, blocking your path."
  } else {
    eventMonster = " The monster you have slain is laying to the side of the path heading west."
  }

  if neAxe {
    axeDescription = " You see an axe lying on the ground."
  } else {
    axeDescription = ""
  }

  if neSword {
    swordDescription = " You see a sword lying on the ground."
  } else {
    swordDescription = ""
  }

  if neRope {
    ropeDescription = " You see a rope lying on the ground."
  } else {
    ropeDescription = ""
  }

  directions = "\nYou can go west or south."

  description = baseDescription + eventMonster + axeDescription + swordDescription + ropeDescription + directions

  s()
  tge.PrintSlow(description, game.TermWidth, slowMode)

  for true {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      cantGo(game.TermWidth)
    } else if userchoice == "east" {
      cantGo(game.TermWidth)
    } else if userchoice == "south" {
      s()
      tge.PrintSlow("You go south.", game.TermWidth, slowMode)
      eArea(player, game)
    } else if userchoice == "west" {
      if tge.InventoryContains(&player.Inventory, "sword") && game.Events["monster"] {
        monsterFight(player, game)
      } else if game.Events["monster"]{
        s()
        tge.PrintSlow("I don't think you can fight the monster without a sword...", game.TermWidth, slowMode)
      } else {
        s()
        tge.PrintSlow("You go west.", game.TermWidth, slowMode)
        nArea(player, game)
      }
    } else if userchoice == "axe" {
      s()
      if tge.InventoryContains(&player.Inventory, "axe") {
        neAxe = true
        tge.PrintSlow("You drop the axe.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "axe")
      } else if neAxe {
        neAxe = false
        tge.PrintSlow("You pick up the axe. It's a nice heavy American felling axe.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "axe")
      } else {
        tge.PrintSlow("What axe?", game.TermWidth, slowMode)
      }
    } else if userchoice == "sword" {
      s()
      if tge.InventoryContains(&player.Inventory, "sword") {
        neSword = true
        tge.PrintSlow("You drop the sword.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "sword")
      } else if neSword {
        neSword = false
        tge.PrintSlow("You pick up the sword. It's a long sword.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "sword")
      } else {
        tge.PrintSlow("What sword?", game.TermWidth, slowMode)
      }
    } else if userchoice == "rope" {
      s()
      if tge.InventoryContains(&player.Inventory, "rope") {
        neRope = true
        tge.PrintSlow("You drop the rope.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "rope")
      } else if neRope {
        neRope = false
        tge.PrintSlow("You pick up the rope. It looks heavy enough to hold your weight.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "rope")
      } else {
        tge.PrintSlow("What rope?", game.TermWidth, slowMode)
      }
    //} else if userchoice == "?" {
      //checkLocalItems(neAxe, neSword, neRope)
    } else if userchoice == "look" {
      if game.Events["monster"] {
        eventMonster = "When you look closer you see there is a monster standing there, blocking your path."
      } else {
        eventMonster = "The monster you have slain is laying to the side of the path heading west."
      }

      if neAxe {
        axeDescription = " You see an axe lying on the ground."
      } else {
        axeDescription = ""
      }

      if neSword {
        swordDescription = " You see a sword lying on the ground."
      } else {
        swordDescription = ""
      }

      if neRope {
        ropeDescription = " You see a rope lying on the ground."
      } else {
        ropeDescription = ""
      }

      directions = "\nYou can go west or south."

      description = baseDescription + eventMonster + axeDescription + swordDescription + ropeDescription + directions

      s()
      tge.PrintSlow(description, game.TermWidth, slowMode)
    } else if userchoice == "inv" {
      s()
      tge.PrintInventory(&player.Inventory)
      s()
    } else if userchoice == "help" || userchoice == "?" {
      s()
      help()
    } else if userchoice == "exit" {
      exit(0)
    } else {
      s()
      tge.PrintSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter another option, or try 'help'.\n", game.TermWidth, slowMode)
    }
  }
}


func monsterFight(player tge.Player, game tge.Game) {
  if game.Events["monster"] {
    var userchoice int
    var damage int
    //var stringDamage string
    monsterHealth := 20
    s()
    tge.PrintSlow("You decide to fight the monster.", game.TermWidth, slowMode)

    for monsterHealth > 0 {
      // I don't remember why this is + 1? Maybe so it is never 0?
      rn := tge.RandNumber(10) + 1
      s()
      tge.PrintSlow("Pick a number between 1 and 10", game.TermWidth, slowMode)
      fmt.Print(" > ")
      fmt.Scan(&userchoice)
      if userchoice <= 10 && userchoice >= 0 {
        if userchoice > rn {
          damage = userchoice - rn
        } else {
          damage = rn - userchoice
        }
        monsterHealth -= damage
        s()
        tge.PrintSlow("You deal " + strconv.Itoa(damage) + " damage to the monster.", game.TermWidth, slowMode)
        //fmt.Println("--------diag---------------")
        //fmt.Println("Random Number:", rn)
        //fmt.Println("User choice:", userchoice)
        //fmt.Println("Damage:", damage)
        //fmt.Println("Monster Health:", monsterHealth)
        //fmt.Println("--------diag---------------")
        if monsterHealth <= 0 {
          s()
          tge.PrintSlow("You Defeated the monster!", game.TermWidth, slowMode)
          s()
          tge.PrintSlow("You go west.", game.TermWidth, slowMode)
          game.Events["monster"] = false
          nArea(player, game)
        }
      } else {
        s()
        tge.PrintSlow("Invalid Number!", game.TermWidth, slowMode)
      }
    }
  } else {
    s()
    tge.PrintSlow("You go west.", game.TermWidth, slowMode)
    nArea(player, game)
  }
}


func nArea(player tge.Player, game tge.Game) {
  // validDirections = north, east

  var userchoice string
  var description string
  var baseDescription string
  var axeDescription string
  var swordDescription string
  var ropeDescription string
  var directions string

  baseDescription = "The forest clears and you find yourself in a field of wildflowers. Purple, blue, yellow and red as far as the eye can see."

  if nAxe {
    axeDescription = " You see an axe lying on the ground."
  } else {
    axeDescription = ""
  }

  if nSword {
    swordDescription = " You see a sword lying on the ground."
  } else {
    swordDescription = ""
  }

  if nRope {
    ropeDescription = " You see a rope lying on the ground."
  } else {
    ropeDescription = ""
  }

  directions = "\nYou can go north or east."

  description = baseDescription + axeDescription + swordDescription + ropeDescription + directions

  s()
  tge.PrintSlow(description, game.TermWidth, slowMode)

  for true {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      tge.PrintSlow("You go north.", game.TermWidth, slowMode)
      exitArea(player, game)
    } else if userchoice == "east" {
      s()
      tge.PrintSlow("You go east.", game.TermWidth, slowMode)
      neArea(player, game)
    } else if userchoice == "south" {
      cantGo(game.TermWidth)
    } else if userchoice == "west" {
      cantGo(game.TermWidth)
    } else if userchoice == "flower" {
      s()
      tge.PrintSlow("You pick a flower and smell it.", game.TermWidth, slowMode)
    } else if userchoice == "axe" {
      s()
      if tge.InventoryContains(&player.Inventory, "axe") {
        nAxe = true
        tge.PrintSlow("You drop the axe.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "axe")
      } else if nAxe {
        nAxe = false
        tge.PrintSlow("You pick up the axe. It's a nice heavy American felling axe.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "axe")
      } else {
        tge.PrintSlow("What axe?", game.TermWidth, slowMode)
      }
    } else if userchoice == "sword" {
      s()
      if tge.InventoryContains(&player.Inventory, "sword") {
        nSword = true
        tge.PrintSlow("You drop the sword.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "sword")
      } else if nSword {
        nSword = false
        tge.PrintSlow("You pick up the sword. It's a long sword.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "sword")
      } else {
        tge.PrintSlow("What sword?", game.TermWidth, slowMode)
      }
    } else if userchoice == "rope" {
      s()
      if tge.InventoryContains(&player.Inventory, "rope") {
        nRope = true
        tge.PrintSlow("You drop the rope.", game.TermWidth, slowMode)
        tge.InventoryRemove(&player.Inventory, "rope")
      } else if nRope {
        nRope = false
        tge.PrintSlow("You pick up the rope. It looks heavy enough to hold your weight.", game.TermWidth, slowMode)
        tge.InventoryAdd(&player.Inventory, "rope")
      } else {
        tge.PrintSlow("What rope?", game.TermWidth, slowMode)
      }
    //} else if userchoice == "?" {
      //checkLocalItems(nAxe, nSword, nRope)
    } else if userchoice == "look" {
      if nAxe {
        axeDescription = " You see an axe lying on the ground."
      } else {
        axeDescription = ""
      }

      if nSword {
        swordDescription = " You see a sword lying on the ground."
      } else {
        swordDescription = ""
      }

      if nRope {
        ropeDescription = " You see a rope lying on the ground."
      } else {
        ropeDescription = ""
      }

      directions = "\nYou can go north or east."

      description = baseDescription + axeDescription + swordDescription + ropeDescription + directions

      s()
      tge.PrintSlow(description, game.TermWidth, slowMode)
    } else if userchoice == "inv" {
      s()
      tge.PrintInventory(&player.Inventory)
      s()
    } else if userchoice == "help" || userchoice == "?" {
      s()
      help()
    } else if userchoice == "exit" {
      exit(0)
    } else {
      s()
      tge.PrintSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter another option, or try 'help'.\n", game.TermWidth, slowMode)
    }
  }
}


func exitArea(player tge.Player, game tge.Game) {
  s()
  tge.PrintSlow("Congratulations, " + player.Name + "!", game.TermWidth, slowMode)
  tge.PrintSlow("You win!", game.TermWidth, slowMode)
  tge.PrintSlow("Hope you had fun!", game.TermWidth, slowMode)
  exit(0)
}


//-----------------------------------------------------------------------------
// Areas end
//-----------------------------------------------------------------------------


//-----------------------------------------------------------------------------
// Global variables start
//-----------------------------------------------------------------------------

// Users name

// slowMode is a cmd line flag to either print normally, or if true it prints 
// character by character.
var slowMode bool

var termWidth int

// Areas inventory
var startAxe bool
var startSword bool
var startRope bool
var wAxe bool
var wSword bool
var wRope bool
var nwAxe bool
var nwSword bool
var nwRope bool
var swAxe bool
var swSword bool
var swRope bool
var sAxe bool
var sSword bool
var sRope bool
var seAxe bool
var seSword bool
var seRope bool
var eAxe bool
var eSword bool
var eRope bool
var neAxe bool
var neSword bool
var neRope bool
var nAxe bool
var nSword bool
var nRope bool



//-----------------------------------------------------------------------------
// Global variables end
//-----------------------------------------------------------------------------

func main() {

  flag.BoolVar(&slowMode, "s", false, "Print out the text in slow mode")
  flag.Parse()

  // Initial value of area inventories
  startAxe = true //startArea has axe to start
  startSword = false
  startRope = false
  wAxe = false
  wSword = false
  wRope = false
  nwAxe = false
  nwSword = true // nwArea has sword to start
  nwRope = false
  swAxe = false
  swSword = false
  swRope = false
  sAxe = false
  sSword = false
  sRope = true // sArea has rope to start
  seAxe = false
  seSword = false
  seRope = false
  eAxe = false
  eSword = false
  eRope = false
  neAxe = false
  neSword = false
  neRope = false
  nAxe = false
  nSword = false
  nRope = false

  var name string
  termWidth, _, err := term.GetSize(0)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  name = intro(termWidth)
  player := tge.Player {
    Name: name,
    Score: 0,
  }

  game := tge.Game {
    Events: make(map[string]bool),
    TermWidth: termWidth,
  }

  game.Events = map[string]bool {
    "log":true,
    "monster":true,
  }

  s()
  startArea(player, game)
}
