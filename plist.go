package main

const plist string = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>Label</key>
  <string>dev.jpedro.device</string>

  <key>ProgramArguments</key>
  <array>
    <string>device</string>
    <string>sync</string>
  </array>

  <key>Nice</key>
  <integer>1</integer>

  <key>StartInterval</key>
  <integer>10</integer>

  <key>RunAtLoad</key>
  <true/>

  <key>StandardOutPath</key>
  <string>/tmp/device.log</string>
  <key>StandardErrorPath</key>
  <string>/tmp/device.err</string>
</dict>
</plist>
`
