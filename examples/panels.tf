
resource "dash_panel" "http_reliability" {
  name = "http_reliability"
  chart {
    title = "Reliability Test1"
    options = {
      "charting.chart" = "radialGauge"
    }
  }
}

