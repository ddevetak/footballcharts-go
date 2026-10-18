# footballcharts-go

[![Go Reference](https://pkg.go.dev/badge/github.com/ddevetak/footballcharts-go.svg)](https://pkg.go.dev/github.com/ddevetak/footballcharts-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-emerald.svg)](LICENSE)

Zero-dependency Go client for the keyless **[Football Charts](https://www.football-charts.com)** REST API (`https://www.football-charts.com/api/public/v1`).

## Features

- **93 Football Leagues Worldwide:** Covers top flights, lower divisions (EFL Championship, League One, League Two, National League North & South, German 3. Liga & Regionalliga, Spanish Primera & Segunda Federación, Italian Serie C, Brasileirão Série B), and women's leagues.
- **Dixon–Coles Match Probabilities:** Calibrated 1X2, Over/Under 2.5, BTTS, and exact scoreline probabilities ([Track Record & Brier Calibration](https://www.football-charts.com/track-record)).
- **10,000-Run Monte Carlo Season Projections:** Daily simulated final standings, title odds, promotion playoff probabilities, and relegation survival points cutoffs.
- **No API Key Required:** Current and previous season endpoints are open and keyless under CC-BY-4.0.

## Installation

```bash
go get github.com/ddevetak/footballcharts-go@v0.1.0
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ddevetak/footballcharts-go"
)

func main() {
	client := footballcharts.NewClient()
	ctx := context.Background()

	leagues, err := client.GetLeagues(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Fetched metadata for Football Charts leagues: %v\n", leagues["count"])
}
```

## Links & Resources

- **Football Charts Homepage:** [https://www.football-charts.com](https://www.football-charts.com)
- **API & MCP Documentation:** [https://www.football-charts.com/developers](https://www.football-charts.com/developers)
- **Open Datasets (CC-BY-4.0 CSV & JSON):** [https://www.football-charts.com/data](https://www.football-charts.com/data)
- **Interactive Dixon–Coles & Fair-Odds Calculator:** [https://www.football-charts.com/tools/dixon-coles-calculator](https://www.football-charts.com/tools/dixon-coles-calculator)
- **Model Calibration Record:** [https://www.football-charts.com/track-record](https://www.football-charts.com/track-record)

## License

MIT License. Data provided by [Football Charts](https://www.football-charts.com) under CC-BY-4.0.
