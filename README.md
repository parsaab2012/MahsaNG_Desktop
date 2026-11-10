# MahsaNG Desktop

**A cross-platform desktop VPN client written in Go.**

An independent desktop port of the MahsaNG Android application, originally developed by MATOK.

This repository is a personal fork maintained by [Parsa](https://github.com/parsaab2012), featuring improvements to subscription fetching, UI attribution, and Windows network recovery.

[Original Repository](https://github.com/bhdrnmatin/MahsaNG_Desktop) · [Releases](https://github.com/parsaab2012/MahsaNG_Desktop/releases) · [Report an Issue](https://github.com/parsaab2012/MahsaNG_Desktop/issues)

---

## Features

### GET CONFIG
- Fetches configurations from multiple public subscription providers.
- Deduplicates configurations.
- Maintains a weighted distribution across providers.
- Skips configurations that failed testing during the previous 24 hours.
- Uses GitHub Contents API for selected GitHub-hosted subscription sources.

### TEST ALL
- Concurrent HTTPS latency testing with 16 workers.
- Stops early after finding enough working configurations.
- Tests real HTTPS payloads instead of relying only on connection establishment.
- Classifies failures by cause:
  - Reset
  - Timeout
  - Throttled
  - Blocked
  - Bad configuration

### Speed Test
Measures download throughput through a selected configuration.

### Configuration Management
- **Delete Invalid:** Removes configurations that failed testing.
- **Delete All:** Removes the entire configuration list after confirmation.
- **Clipboard Import:** Supports individual share links, multiple links, and Base64 subscription content.

### Connect
Starts a local SOCKS5/HTTP proxy and, where supported, a TUN interface to route system traffic through the selected server.

### Serverless Configurations
Supports complete Xray configurations that do not require a traditional proxy server.

System-wide TUN support depends on the operating system and network interface capabilities. On Windows and macOS, serverless configurations use local SOCKS5/HTTP proxy mode.

### Automatic Wi-Fi Recovery (Windows)

This fork introduces an optional Windows-only network recovery feature.

When enabled:
1. The user disconnects from the VPN through the application.
2. MahsaNG finishes tunnel cleanup.
3. The application attempts to restart the detected physical Wi-Fi adapter.
4. The recovery result is displayed in the application's status message.

The feature is disabled by default.

**Requirements and limitations:**
- Windows only.
- Requires Administrator privileges.
- Temporarily interrupts the Wi-Fi connection.
- Only proceeds when exactly one matching physical wireless adapter is detected.

---

## Supported Configurations

Supported share-link protocols:

- VMess
- VLESS
- Trojan
- Shadowsocks

Supported transports and security options depend on the configuration and parser implementation, including TCP, WebSocket, gRPC, HTTP/H2, TLS, and REALITY.

---

## Fork Improvements

Compared with the original repository, this fork includes:

### Subscription Provider Improvements
- Updated selected GitHub-hosted subscription sources to use GitHub Contents API.
- Added appropriate GitHub API request headers.
- Updated the built-in provider list.
- Removed the previously configured Epodonios subscription source.

### UI Attribution
- Added a link to Parsa's GitHub profile.
- Preserved the original MATOK GitHub and Telegram credits.

### Windows Network Recovery
- Added an optional Wi-Fi recovery checkbox.
- Added persistent preference storage.
- Added Windows-specific adapter recovery logic.

---

## Download

Prebuilt binaries will be published through the GitHub Releases page:

**[MahsaNG Desktop Releases](https://github.com/parsaab2012/MahsaNG_Desktop/releases)**

No prebuilt release has been published in this fork yet. Until the first release is available, build the application from source.

---

## Usage

1. Launch MahsaNG Desktop.
2. Press **GET CONFIG** or import your own configuration using the `+` button.
3. Press **TEST ALL**.
4. Press **SORT** to arrange configurations by speed.
5. Select a configuration and press **Connect**.
6. Press **Disconnect** when finished.

On Windows, the optional Wi-Fi recovery setting can be enabled from the main interface.

---

## Build from Source

Requirements:
- Go (check `go.mod` for the required version).
- A C compiler, because Fyne uses CGO.
- Platform-specific development dependencies.

### Linux

Install the required development packages:

```bash
sudo apt update
sudo apt install gcc libgl1-mesa-dev xorg-dev
```

Build:

```bash
go build -o mahsang ./cmd/mahsang
go build -o mahsang-cli ./cmd/mahsang-cli
```

### Windows

Use MinGW-w64 GCC, for example through MSYS2 UCRT64.

With MSYS2 UCRT64 installed:

```powershell
$env:Path = "C:\msys64\ucrt64\bin;$env:Path"
$env:CGO_ENABLED = "1"
```

Build:

```powershell
go build -o mahsang.exe ./cmd/mahsang
```

For system-wide VPN connections and optional Wi-Fi recovery, run the application as Administrator.

Keep `wintun.dll` next to the executable.

### macOS

Install Xcode Command Line Tools:

```bash
xcode-select --install
```

Then build:

```bash
go build -o mahsang ./cmd/mahsang
```

---

## Testing

Run the complete Go test suite:

```bash
go test ./...
```

---

## CLI

The project also includes a command-line interface:

```bash
mahsang-cli get
mahsang-cli test <subURL|->
mahsang-cli ping <share-link>
mahsang-cli connect <link>
mahsang-cli speed <link>
```

The CLI shares the same underlying engine and configuration processing components.

---

## Architecture

```text
Provider
   |
   v
Parser
   |
   v
Tester
   |
   v
Core (xray-core)
   |
   v
TUN / tun2socks
```

The project is organized into independent internal packages:

| Package | Responsibility |
|---|---|
| `core` | Xray engine and tunnel management |
| `parser` | Share-link parsing |
| `provider` | Subscription retrieval |
| `tester` | Configuration testing |
| `store` | Persistent configuration storage |
| `tun` | System-wide routing |
| `ui` | Fyne desktop interface |
| `model` | Application data models |
| `logx` | Logging |

---

## Known Limitations

- Public configuration sources are volatile and may become unavailable.
- Configuration testing does not guarantee long-term availability.
- System-wide Connect requires elevated privileges.
- Windows and macOS routing may behave differently from Linux.
- Windows Wi-Fi recovery is optional and may temporarily interrupt wireless connectivity.
- DNS requests to some LAN resolvers may bypass the tunnel.
- Serverless configurations may use local proxy mode depending on the operating system.

---

## Credits

### Original Project

Developed by **MATOK**.

- GitHub: [matinbhdrn77](https://github.com/matinbhdrn77)
- Telegram: [@mat_bh](https://t.me/mat_bh)
- Original repository: [MahsaNG_Desktop](https://github.com/bhdrnmatin/MahsaNG_Desktop)

### Fork and Modifications

Maintained by **Parsa**.

- GitHub: [parsaab2012](https://github.com/parsaab2012)

---

## License

This project retains the original MIT License.

Original copyright notice:

```text
Copyright (c) 2026 matin
```

See [LICENSE](LICENSE) for the full license text.
