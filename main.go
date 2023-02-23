package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	// "gopkg.in/jpedro/color.v1"
	"github.com/jpedro/color"
)

var (
    client  *db.Client
    ctx     context.Context
    serial  string
    model   string
    version string
    uuid string
)

func run(command string, args ...string) string {
    cmd := exec.Command(command, args...)
    bytes, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("Error calling command '%s': %v\n", command, err)
    }
    return string(bytes)
}

func grab(stdout string, key string) string {
    lines := strings.Split(stdout, "\n")
    found := ""
    for _, line := range lines {
        if strings.Contains(line, key) {
            parts := strings.Split(line, ":")
            found = strings.TrimSpace(parts[1])
        }
    }
    return found
}

func loadHardware() {
    stdout := run("system_profiler", "SPHardwareDataType")
    serial = grab(stdout, "Serial Number (system):")
    model = grab(stdout, "Model Name:")
    uuid = grab(stdout, "Hardware UUID:")
}

func loadSoftware() {
    stdout := run("sw_vers")
    version = grab(stdout, "ProductVersion:")
}

func connect() {
    ctx = context.Background()
    config := &firebase.Config{
        ProjectID:   "jpedro-store",
        DatabaseURL: "https://jpedro-store.firebaseio.com",
    }
    app, err := firebase.NewApp(ctx, config)
    if err != nil {
        log.Fatalf("Error initializing app: %v\n", err)
    }

    client, err = app.Database(ctx)
    if err != nil {
        log.Fatalf("Error connecting to database: %v\n", err)
    }
}

func save(path string, record any) {
    err := client.NewRef(path).Set(ctx, record)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    go func() {
        fmt.Println("Hello")
    }()

    loadHardware()
    loadSoftware()
    fmt.Printf("Serial: %s.\n", serial)
    fmt.Printf("Model name: %s.\n", model)
    fmt.Printf("Version: %s.\n", version)

    connect()
    fmt.Println("Client:", client)

    save("devices/"+serial, Device{
        Serial:  serial,
        Model: model,
        Version: version,
        UUID: uuid,
    })
    fmt.Printf("Device saved: %s.\n", color.Green(serial))

    save("messages/"+serial, Message{
        Content: "Hei " + time.Now().Format("2006-01-31 01:37:58"),
        Created: time.Now(),
    })

    var stored Device
    if err := client.NewRef("devices/"+serial).Get(ctx, &stored); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Device %s has UUID %s.\n", color.Green(stored.Serial), color.Green(stored.UUID))
}
