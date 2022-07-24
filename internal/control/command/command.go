package command

type Command struct {
  Name string
  Description string
  Execute func([]string)
}

func CreateCommand(name string, description string, execute func([]string)) *Command {
  return struct {
    Name: name,
    Description: description,
    Execute: execute,
  }
}
