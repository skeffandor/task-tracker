package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	mngr "github.com/skeffandor/task-tracker/internal/manager"
)

func InitCLI(tm *mngr.TaskManager) {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Println("Task CLI. Type 'help' for commands.")
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		line = strings.TrimSpace(line)
		args := strings.SplitN(line, " ", 2)
		cmd := args[0]

		switch cmd {
		case "add":
			if len(args) < 2 {
				fmt.Println("Usage: add <description>")
				continue
			}

			AddedId := tm.Add(args[1])
			fmt.Printf("Task %d added.\n", AddedId)

		case "update":
			parts := strings.SplitN(args[1], " ", 2)
			if len(parts) < 2 {
				fmt.Println("Usage: update <id> <new description>")
				continue
			}
			id, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Invalid ID.")
				continue
			}

			if ok := tm.Update(mngr.Id(id), parts[1]); !ok {
				fmt.Println("Task not found.")
			} else {
				fmt.Println("Task updated.")
			}

		case "delete":
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid ID.")
				continue
			}

			if ok := tm.Delete(mngr.Id(id)); !ok {
				fmt.Println("Task not found.")
			} else {
				fmt.Println("Task deleted.")
			}

		case "mark-in-progress":
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid ID.")
				continue
			}

			if ok := tm.ChangeStatus(mngr.Id(id), mngr.InProgress); !ok {
				fmt.Println("Task not found.")
			} else {
				fmt.Printf("Task %d marked as %s.\n", id, mngr.InProgress)
			}

		case "mark-done":
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid ID.")
				continue
			}

			if ok := tm.ChangeStatus(mngr.Id(id), mngr.Done); !ok {
				fmt.Println("Task not found.")
			} else {
				fmt.Printf("Task %d marked as %s.\n", id, mngr.Done)
			}

		case "list":
			var statusArg mngr.Status
			if len(args) < 2 {
				statusArg = mngr.Any
			} else {
				statusArg = mngr.Status(args[1])
			}

			if statusArg.IsValid() {
				if c := tm.List(statusArg); c == 0 {
					fmt.Println("No tasks.")
				}
			} else {
				fmt.Println("Unknown argumrnt. Use <todo>, <in-progress> or <done>.")
			}

		case "help":
			fmt.Println("Commands:")
			fmt.Println("  add <description>")
			fmt.Println("  update <id> <new description>")
			fmt.Println("  delete <id>")
			fmt.Println("  mark-in-progress <id>")
			fmt.Println("  mark-done <id>")
			fmt.Println("  list")
			fmt.Println("  exit")

		case "exit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Type 'help' for list.")
		}
	}
}
