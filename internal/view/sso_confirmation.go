package view

import (
	"context"
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/clawscli/claws/internal/action"
	"github.com/clawscli/claws/internal/registry"
	"github.com/clawscli/claws/internal/ui"
)

// SSOLoginConfirmationView displays a confirmation dialog for SSO login
type SSOLoginConfirmationView struct {
	ctx      context.Context
	registry *registry.Registry
	errorMsg string
	styles   ssoConfirmationStyles
}

type ssoConfirmationStyles struct {
	title lipgloss.Style
	error lipgloss.Style
	box   lipgloss.Style
	yes   lipgloss.Style
	no    lipgloss.Style
	bold  lipgloss.Style
}

func newSSOConfirmationStyles() ssoConfirmationStyles {
	t := ui.Current()
	return ssoConfirmationStyles{
		title: lipgloss.NewStyle().Bold(true).Foreground(t.Primary),
		error: lipgloss.NewStyle().Foreground(t.Danger),
		box:   lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(t.Border).Padding(0, 1).MarginTop(1),
		yes:   lipgloss.NewStyle().Bold(true).Foreground(t.Success),
		no:    lipgloss.NewStyle().Bold(true).Foreground(t.Danger),
		bold:  lipgloss.NewStyle().Bold(true),
	}
}

// NewSSOLoginConfirmationView creates a new SSO login confirmation view
func NewSSOLoginConfirmationView(ctx context.Context, registry *registry.Registry, errorMsg string) *SSOLoginConfirmationView {
	return &SSOLoginConfirmationView{
		ctx:      ctx,
		registry: registry,
		errorMsg: errorMsg,
		styles:   newSSOConfirmationStyles(),
	}
}

// Init implements tea.Model
func (v *SSOLoginConfirmationView) Init() tea.Cmd {
	return nil
}

// Update handles input for the confirmation view
func (v *SSOLoginConfirmationView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "y", "Y", "enter":
			// User confirmed - trigger SSO login
			exec := &action.SimpleExec{
				Command:    "aws sso login",
				ActionName: action.ActionNameLogin,
				SkipAWSEnv: true,
			}
			return v, tea.Exec(exec, func(err error) tea.Msg {
				if err != nil {
					return ErrorMsg{Err: err}
				}
				return RefreshMsg{}
			})
		case "n", "N", "esc":
			// User cancelled - return to normal error display
			return v, func() tea.Msg { return HideModalMsg{} }
		}
	}
	return v, nil
}

// ViewString renders the confirmation dialog
func (v *SSOLoginConfirmationView) ViewString() string {
	s := v.styles

	var out string
	out += s.title.Render("SSO Token Error") + "\n\n"

	// Truncate error message if too long
	errorMsg := v.errorMsg
	if len(errorMsg) > 500 {
		errorMsg = errorMsg[:497] + "..."
	}
	out += s.error.Render(fmt.Sprintf("Error: %s", errorMsg)) + "\n\n"

	confirmContent := s.bold.Render("Confirm SSO Login") + "\n"
	confirmContent += "Perform SSO login to refresh credentials?\n\n"
	confirmContent += "Press " + s.yes.Render("[Y]") + " to login or " + s.no.Render("[N]") + " to cancel"

	out += s.box.Render(confirmContent)

	return out
}

// View implements tea.Model
func (v *SSOLoginConfirmationView) View() tea.View {
	return tea.NewView(v.ViewString())
}

// SetSize implements View
func (v *SSOLoginConfirmationView) SetSize(width, height int) tea.Cmd {
	return nil
}

// HasActiveInput implements View
func (v *SSOLoginConfirmationView) HasActiveInput() bool {
	return false
}

// StatusLine implements View
func (v *SSOLoginConfirmationView) StatusLine() string {
	return "SSO Login Confirmation • Y to login • N to cancel"
}
