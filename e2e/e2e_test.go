package e2e

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

func TestUUID(t *testing.T) {
	pwd, _ := filepath.Abs(".")
	os.Chdir("../")
	defer func() {
		os.Chdir(pwd)
	}()

	cmd := exec.Command("go", "run", "./...", "uuid")
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf

	err := cmd.Run()
	if err != nil {
		t.Errorf("failed to exec command: %v", err)
		return
	}

	id := strings.TrimRight(buf.String(), "\n")
	if _, err := uuid.Parse(id); err != nil {
		t.Errorf("invalid uuid: %v", err)
	}
}

func TestUUIDV7(t *testing.T) {
	pwd, _ := filepath.Abs(".")
	os.Chdir("../")
	defer func() {
		os.Chdir(pwd)
	}()

	cmd := exec.Command("go", "run", "./...", "uuid")
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf

	err := cmd.Run()
	if err != nil {
		t.Errorf("failed to exec command: %v", err)
		return
	}

	id := strings.TrimRight(buf.String(), "\n")
	if _, err := uuid.Parse(id); err != nil {
		t.Errorf("invalid uuidv7: %v", err)
	}
}

func TestULID(t *testing.T) {
	pwd, _ := filepath.Abs(".")
	os.Chdir("../")
	defer func() {
		os.Chdir(pwd)
	}()

	cmd := exec.Command("go", "run", "./...", "ulid")
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf

	err := cmd.Run()
	if err != nil {
		t.Errorf("failed to exec command: %v", err)
		return
	}

	id := strings.TrimRight(buf.String(), "\n")
	if _, err := ulid.Parse(id); err != nil {
		t.Errorf("invalid ulid: %v", err)
	}
}
