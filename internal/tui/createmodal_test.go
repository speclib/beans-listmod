package tui

import (
	"context"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hmans/beans/pkg/bean"
	"github.com/hmans/beans/pkg/config"
)

type fakeBeanGetter struct {
	beans map[string]*bean.Bean
}

func (f fakeBeanGetter) Bean(_ context.Context, id string) (*bean.Bean, error) {
	return f.beans[id], nil
}

func newTestCreateModal(parentID string) createModalModel {
	lookup := fakeBeanGetter{beans: map[string]*bean.Bean{
		"m1": {ID: "m1", Type: "milestone", Title: "Release 1.0"},
		"e1": {ID: "e1", Type: "epic", Title: "Terminal work"},
	}}
	return newCreateModalModel(lookup, config.Default(), parentID, 100, 40)
}

func TestCreateModalFocusCycle(t *testing.T) {
	m := newTestCreateModal("m1")
	if m.focus != focusTitle {
		t.Fatalf("focus = %v, want focusTitle", m.focus)
	}

	m, _ = m.Update(tea.KeyMsg{Type: tea.KeyTab})
	if m.focus != focusParent {
		t.Errorf("after tab: focus = %v, want focusParent", m.focus)
	}

	m, _ = m.Update(tea.KeyMsg{Type: tea.KeyShiftTab})
	if m.focus != focusTitle {
		t.Errorf("after shift+tab: focus = %v, want focusTitle", m.focus)
	}

	m, _ = m.Update(tea.KeyMsg{Type: tea.KeyShiftTab})
	if m.focus != focusParent {
		t.Errorf("shift+tab from title: focus = %v, want focusParent", m.focus)
	}
}

func TestCreateModalOpensParentPicker(t *testing.T) {
	m := newTestCreateModal("m1")
	m.textInput.SetValue("New thing")
	m, _ = m.Update(tea.KeyMsg{Type: tea.KeyTab})

	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("enter on parent row returned no command")
	}

	msg, ok := cmd().(openParentPickerMsg)
	if !ok {
		t.Fatalf("message = %T, want openParentPickerMsg", cmd())
	}
	if msg.beanTitle != "New thing" {
		t.Errorf("beanTitle = %q, want %q", msg.beanTitle, "New thing")
	}
	if len(msg.beanTypes) != 1 || msg.beanTypes[0] != "task" {
		t.Errorf("beanTypes = %v, want [task]", msg.beanTypes)
	}
	if msg.currentParent != "m1" {
		t.Errorf("currentParent = %q, want m1", msg.currentParent)
	}
}

func TestCreateModalCreatesWithParent(t *testing.T) {
	m := newTestCreateModal("e1")
	m.textInput.SetValue("Child task")

	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("enter on title returned no command")
	}

	msg, ok := cmd().(beanCreatedMsg)
	if !ok {
		t.Fatalf("message = %T, want beanCreatedMsg", cmd())
	}
	if msg.title != "Child task" {
		t.Errorf("title = %q, want %q", msg.title, "Child task")
	}
	if msg.parentID != "e1" {
		t.Errorf("parentID = %q, want e1", msg.parentID)
	}
}

func TestCreateModalCreatesWithoutParent(t *testing.T) {
	m := newTestCreateModal("")
	m.textInput.SetValue("Standalone task")

	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	msg, ok := cmd().(beanCreatedMsg)
	if !ok {
		t.Fatalf("message = %T, want beanCreatedMsg", cmd())
	}
	if msg.parentID != "" {
		t.Errorf("parentID = %q, want empty", msg.parentID)
	}
}

func TestCreateModalEmptyTitleCloses(t *testing.T) {
	m := newTestCreateModal("m1")

	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("enter with empty title returned no command")
	}
	if _, ok := cmd().(closeCreateModalMsg); !ok {
		t.Fatalf("message = %T, want closeCreateModalMsg", cmd())
	}
}

func TestCreateModalEscapeCloses(t *testing.T) {
	m := newTestCreateModal("m1")

	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if cmd == nil {
		t.Fatal("esc returned no command")
	}
	if _, ok := cmd().(closeCreateModalMsg); !ok {
		t.Fatalf("message = %T, want closeCreateModalMsg", cmd())
	}
}

func TestCreateModalViewShowsParent(t *testing.T) {
	m := newTestCreateModal("m1")
	view := m.View()
	if !strings.Contains(view, "milestone Release 1.0") {
		t.Errorf("view does not show parent type and title:\n%s", view)
	}
	if !strings.Contains(view, "Parent") {
		t.Errorf("view does not show parent row:\n%s", view)
	}
}

func TestCreateModalViewShowsNone(t *testing.T) {
	m := newTestCreateModal("")
	m.focus = focusParent
	view := m.View()
	if !strings.Contains(view, "(none)") {
		t.Errorf("view does not show (none):\n%s", view)
	}
	if !strings.Contains(view, "change") || !strings.Contains(view, "title") {
		t.Errorf("view does not show parent focus hints:\n%s", view)
	}
}

func TestParentSelectedUpdatesCreateModal(t *testing.T) {
	app := &App{
		state:         viewParentPicker,
		previousState: viewCreateModal,
		createModal:   newTestCreateModal("m1"),
	}
	app.createModal.focus = focusParent

	app.Update(parentSelectedMsg{parentID: "e1"})

	if app.createModal.parentID != "e1" {
		t.Errorf("parentID = %q, want e1", app.createModal.parentID)
	}
	if app.createModal.focus != focusTitle {
		t.Errorf("focus = %v, want focusTitle", app.createModal.focus)
	}
	if app.state != viewCreateModal {
		t.Errorf("state = %v, want viewCreateModal", app.state)
	}
}

func TestCreateModalEscapeReturnsToBackground(t *testing.T) {
	app := &App{
		state:            viewCreateModal,
		previousState:    viewList,
		createBackground: viewList,
		createModal:      newTestCreateModal("m1"),
	}

	_, cmd := app.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if cmd == nil {
		t.Fatal("esc returned no command")
	}
	app.Update(cmd())

	if app.state != viewList {
		t.Errorf("state = %v, want viewList", app.state)
	}
}

func TestCreateModalEscapeAfterPickerSelection(t *testing.T) {
	app := &App{
		state:            viewParentPicker,
		previousState:    viewCreateModal,
		createBackground: viewList,
		createModal:      newTestCreateModal("m1"),
	}

	app.Update(parentSelectedMsg{parentID: "e1"})
	if app.state != viewCreateModal {
		t.Fatalf("after selection: state = %v, want viewCreateModal", app.state)
	}

	_, cmd := app.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if cmd == nil {
		t.Fatal("esc returned no command")
	}
	app.Update(cmd())
	if app.state != viewList {
		t.Errorf("after esc: state = %v, want viewList", app.state)
	}
}

func TestCreateModalEscapeAfterPickerCancel(t *testing.T) {
	app := &App{
		state:            viewParentPicker,
		previousState:    viewCreateModal,
		createBackground: viewList,
		createModal:      newTestCreateModal("m1"),
	}

	app.Update(closeParentPickerMsg{})
	if app.state != viewCreateModal {
		t.Fatalf("after cancel: state = %v, want viewCreateModal", app.state)
	}

	_, cmd := app.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if cmd == nil {
		t.Fatal("esc returned no command")
	}
	app.Update(cmd())
	if app.state != viewList {
		t.Errorf("after esc: state = %v, want viewList", app.state)
	}
}

func TestCloseParentPickerReturnsToCreateModal(t *testing.T) {
	app := &App{
		state:         viewParentPicker,
		previousState: viewCreateModal,
		createModal:   newTestCreateModal("m1"),
	}

	app.Update(closeParentPickerMsg{})

	if app.state != viewCreateModal {
		t.Errorf("state = %v, want viewCreateModal", app.state)
	}
	if app.createModal.parentID != "m1" {
		t.Errorf("parentID = %q, want m1", app.createModal.parentID)
	}
}
