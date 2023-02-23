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


// https://scim.eu-west-1.amazonaws.com/vdl9578616c-559f-42aa-9862-dae465c94998/scim/v2/
// 2fc6456d-9119-4484-ba08-4dc309d943d1:e9a72e2c-d4c9-4a38-809b-df03ad7b6b57:QOv7fgfd/j7fR4UiJBe7sSxUsMVmrNrWyckAqTlQ8iWKkWrCaUnSvamkLc/Bgb0mHNNjraAXZN0ZWTVpIM/jlGats6OhCAgdUkFA+xZx4IMJpTeqx77shNQmmAPBHkOgTsIOwRyrgihABZ2Mi6FM0o7ocCvWGFPVlvM=:Bck2L18BdFdmX3nBfGc6LWxl2CXWWMJB5YQEif/wMd68KPKMz4dLbAN9Uq26/VyI/l6sNuR4x5iLf6w9VbrnKnBkBS9FJeHRxicl6bzhst7ICWw6Uo24rk0F6AVGvTs/h/VdQrh5Tr+ObjAs9VPfFBe7JZgrJHfFaSucVsas6QRj+dHGHbelno/bre7LUOZA30tbeoSHKTf2JY51uG5j7m3npAiwGoCQku6BocC8yTE8WmW8xIxCO3ArFqC7FdoxPm29Nf1olUG1j45AxsweF8IK96gMhHrkpl0C19bZ8W4Fj8wvto+jzfpsN4CMnbv63h7oAt2JtbCfxwaWFgUEBQ==
