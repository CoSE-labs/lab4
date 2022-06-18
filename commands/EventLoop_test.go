package commands

import (
  "testing"  
  "github.com/stretchr/testify/assert"
)

func TestELoop(t *testing.T) {
  revCmd1 := &RevCommand{
    arg: "Github!",
  }

  revCmd2 := &RevCommand{
    arg: "11223344556677889900",
  }

  printCmd := &PrintCommand{
    line: "Printing...",
  }

  eventLoop := new(EventLoop)
  eventLoop.Start()

  cmd1 := RevCommand(*revCmd1)
  cmd2 := RevCommand(*revCmd2)
  cmd3 := PrintCommand(*printCmd)

  assert.Equal(t, false, eventLoop.stop)
  assert.Equal(t, 0, len(eventLoop.queue.CmdArray))

  eventLoop.Post(&cmd1)
  eventLoop.Post(&cmd2)
  eventLoop.Post(cmd3)

  assert.Equal(t, 3, len(eventLoop.queue.CmdArray))
  eventLoop.AwaitFinish()
  assert.Equal(t, true, eventLoop.stop)
  assert.Equal(t, 0, len(eventLoop.queue.CmdArray))
}