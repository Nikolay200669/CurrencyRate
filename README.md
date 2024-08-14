currency-app/
│
├── cmd/
│   └── currency-app/
│       └── main.go
│
├── internal/
│   ├── api/
│   │   └── privatbank.go
│   ├── config/
│   │   └── config.go
│   ├── currency/
│   │   └── currency.go
│   ├── storage/
│   │   ├── csv.go
│   │   ├── json.go
│   │   └── xml.go
│   ├── tray/
│   │   └── tray.go
│   └── utils/
│       └── logger.go
│
├── pkg/
│   └── autostart/
│       ├── autostart_windows.go
│       ├── autostart_linux.go
│       └── autostart_darwin.go
│
├── configs/
│   ├── config.json
│   └── currency.json
│
├── data/
│   └── currency.csv
│
├── assets/
│   └── icon.png
│
├── go.mod
├── go.sum
├── README.md
└── Makefile
