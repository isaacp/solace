// package main

// // A simple program demonstrating the text input component from the Bubbles
// // component library.

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"time"

// 	"github.com/charmbracelet/bubbles/textinput"
// 	"github.com/charmbracelet/bubbles/viewport"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// 	"github.com/creack/pty"
// 	"github.com/solace/terminal"
// )

// var p *os.File
// var term terminal.Terminal

// type leader struct {
// 	file *os.File
// }

// func (l *leader) Initialize() error {
// 	c := exec.Command("/bin/zsh")
// 	file, err := pty.Start(c)
// 	if err != nil {
// 		return errors.New("could not start shell")
// 	}
// 	l.file = file
// 	return nil
// }

// func (l *leader) Write(message string) (int, error) {
// 	return l.file.WriteString(message)
// }

// func (l *leader) Read() (string, error) {
// 	arr := make([]byte, 16384)
// 	l.file.Read(arr)
// 	return string(arr), nil
// }

// func (l *leader) Close() error {
// 	return l.file.Close()
// }

// var l leader

// func main() {
// 	// c := exec.Command("/bin/zsh")
// 	// pt, err := pty.Start(c)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	os.Exit(1)
// 	// }
// 	// p = pt

// 	l = leader{}
// 	l.Initialize()
// 	defer l.Close()
// 	//term.Run()
// 	pr := tea.NewProgram(
// 		initialModel(),
// 		//tea.WithAltScreen(),
// 		//tea.WithMouseCellMotion(),
// 	)
// 	if _, err := pr.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// type (
// 	errMsg error
// )

// type model struct {
// 	textInput textinput.Model
// 	content   string
// 	ready     bool
// 	viewport  viewport.Model
// 	err       error
// }

// func initialModel() model {
// 	ti := textinput.New()
// 	ti.Placeholder = "<command>"
// 	ti.Focus()
// 	ti.CharLimit = 300
// 	ti.Width = 300

// 	vp := viewport.New(100, 300)
// 	vp.Style = lipgloss.NewStyle()
// 	vp.HighPerformanceRendering = true
// 	vp.MouseWheelEnabled = true

// 	return model{
// 		textInput: ti,
// 		viewport:  vp,
// 		err:       nil,
// 	}
// }

// func (m model) Init() tea.Cmd {
// 	return textinput.Blink
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd
// 	var cmd1 tea.Cmd
// 	var cmds []tea.Cmd

// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.Type {
// 		case tea.KeyCtrlC, tea.KeyEsc:
// 			return m, tea.Quit
// 		case tea.KeyEnter:
// 			l.Write(m.textInput.Value() + "\r")
// 			time.Sleep(100 * time.Millisecond)
// 			// t := make([]byte, 16092)
// 			t, _ := l.Read()
// 			m.content = t
// 			//fmt.Println(t)
// 			//m.viewport.SetContent(t)
// 			m.textInput.Reset()
// 		}

// 	case tea.WindowSizeMsg:
// 		if !m.ready {
// 			// Since this program is using the full size of the viewport we
// 			// need to wait until we've received the window dimensions before
// 			// we can initialize the viewport. The initial dimensions come in
// 			// quickly, though asynchronously, which is why we wait for them
// 			// here.
// 			m.viewport = viewport.New(msg.Width, msg.Height)
// 			m.viewport.HighPerformanceRendering = true
// 			m.viewport.SetContent(m.content)
// 			m.ready = true

// 			// This is only necessary for high performance rendering, which in
// 			// most cases you won't need.
// 			//
// 			// Render the viewport one line below the header.
// 			//m.viewport.YPosition = headerHeight + 1
// 		} else {
// 			m.viewport.Width = msg.Width
// 			m.viewport.Height = msg.Height
// 		}

// 		if m.viewport.HighPerformanceRendering {
// 			// Render (or re-render) the whole viewport. Necessary both to
// 			// initialize the viewport and when the window is resized.
// 			//
// 			// This is needed for high-performance rendering only.
// 			cmds = append(cmds, viewport.Sync(m.viewport))
// 		}

// 	// We handle errors just like any other message
// 	case errMsg:
// 		m.err = msg
// 		return m, nil
// 	}

// 	m.textInput, cmd = m.textInput.Update(msg)
// 	cmds = append(cmds, cmd)
// 	m.viewport, cmd1 = m.viewport.Update(msg)
// 	cmds = append(cmds, cmd1)
// 	fmt.Println(m.ready)
// 	return m, tea.Batch(cmds...)
// }

// func (m model) View() string {
// 	if !m.ready {
// 		return "\n  Initializing..."
// 	}
// 	return fmt.Sprintf("%s\n%s\n\n%s",
// 		m.viewport.View(),
// 		m.textInput.View(),
// 		"(esc to quit)",
// 	) + "\n"
// }

// // package main

// // import (
// // 	"fmt"
// // 	"log"
// // 	"time"

// // 	"github.com/solace/terminal"
// // )

// // func main() {
// // 	term := terminal.NewTerminal(terminal.Zsh)
// // 	term.OpenShell()
// // 	if err := term.Run(); err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	line := ""
// // 	for line != "quit" {
// // 		fmt.Scanln(&line)
// // 		term.Write(line)
// // 		time.Sleep(1000)
// // 		message, _ := term.Read()
// // 		fmt.Println(message)
// // 	}
// // }

package main

// An example program demonstrating the pager component from the Bubbles
// component library.

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// You generally won't need this unless you're processing stuff with
// complicated ANSI escape sequences. Turn it on if you notice flickering.
//
// Also keep in mind that high performance rendering only works for programs
// that use the full size of the terminal. We're enabling that below with
// tea.EnterAltScreen().
const useHighPerformanceRenderer = false

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

type model struct {
	content  string
	ready    bool
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.content)
			m.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			// Render (or re-render) the whole viewport. Necessary both to
			// initialize the viewport and when the window is resized.
			//
			// This is needed for high-performance rendering only.
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m model) headerView() string {
	title := titleStyle.Render("Mr. Pager")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Load some text for our viewport
	content, err := os.ReadFile("main.go")
	if err != nil {
		fmt.Println("could not load file:", err)
		os.Exit(1)
	}

	p := tea.NewProgram(
		model{content: string(content)},
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
