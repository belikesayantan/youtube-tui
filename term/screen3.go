package term

import (
	"fmt"
	"log"
	"runtime"

	"github.com/belikesayantan/ytmusic-cli/ytsearch"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Renders the screen in the terminal
func Initscreen3(filetype string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	YTdesign := widgets.NewParagraph()
	YTdesign.SetRect(40, 0, 400, 9)
	YTdesign.Border = false
	YTdesign.TextStyle.Fg = ui.ColorRed
	YTdesign.Text = `
	██╗   ██╗████████╗      ██████╗ ██╗      █████╗ ██╗   ██╗███████╗██████╗ 
	╚██╗ ██╔╝╚══██╔══╝      ██╔══██╗██║     ██╔══██╗╚██╗ ██╔╝██╔════╝██╔══██╗
	 ╚████╔╝    ██║         ██████╔╝██║     ███████║ ╚████╔╝ █████╗  ██████╔╝
	  ╚██╔╝     ██║         ██╔═══╝ ██║     ██╔══██║  ╚██╔╝  ██╔══╝  ██╔══██╗
	   ██║      ██║         ██║     ███████╗██║  ██║   ██║   ███████╗██║  ██║
	   ╚═╝      ╚═╝         ╚═╝     ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝╚═╝  ╚═╝`

	title := widgets.NewParagraph()
	title.Text = "-- A Online Terminal Music App"
	title.Border = false
	title.TextStyle.Fg = ui.ColorRed
	title.SetRect(105, 9, 149, 12)

	query := ""
	st := &keypress{}
	prompt := func(st *keypress) {

		ui.Clear()
		p := widgets.NewParagraph()
		p.Title = "Search Box"
		p.Text = fmt.Sprintf("%s", st.queryRender)

		p.SetRect(40, 10, 100, 13)

		ui.Render(YTdesign, title, p)
	}

	prompt(st)

	menuEvents := ui.PollEvents()

	for {
		e := <-menuEvents

		if e.Type == ui.KeyboardEvent && e.ID == "<Enter>" {
			results := ytsearch.GetMusicList(query)
			QueryResults(query, results, filetype)
			query = ""
			break
		}
		if e.Type == ui.KeyboardEvent && e.ID == "<Escape>" {
			Initscreen2()
		}
		if e.Type == ui.KeyboardEvent && e.ID == "<Space>" {
			query += " "
			st.queryRender = query
			prompt(st)
		}
		if runtime.GOOS == "windows" {
			if e.Type == ui.KeyboardEvent && e.ID == "<C-<Backspace>>" {
				if len(query) == 0 {
					break
				}
				query = query[:len(query)-1]
				st.queryRender = query
				prompt(st)
			}
		} else {
			if e.Type == ui.KeyboardEvent && e.ID == "<Backspace>" {
				if len(query) == 0 {
					break
				}
				query = query[:len(query)-1]
				st.queryRender = query
				prompt(st)
			}
		}
		if runtime.GOOS == "windows" {
			if e.Type == ui.KeyboardEvent && e.ID != "<Escape>" && e.ID != "<Space>" && e.ID != "<C-<Backspace>>" {
				query += e.ID
				st.queryRender = query
				prompt(st)
			}
		} else {
			if e.Type == ui.KeyboardEvent && e.ID != "<Escape>" && e.ID != "<Space>" && e.ID != "<Backspace>" {
				query += e.ID
				st.queryRender = query
				prompt(st)
			}
		}

	}

}