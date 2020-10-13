I am trying to write a custom terraform provider. In that, I have a resource called dash_panel which looks something like this:

```
resource "dash_panel" "http_reliability" {
  name = "http_reliability"
  chart {
    title = "Reliability Test1"
    options = {
      "charting.chart" = "radialGauge"
    }
  }
}
```

Just to illustrate the issue, I have made this a local-only resource.

This resource is expected to return an xml that looks like this:

```
<panel>
    <chart>
        <title>Reliability Test1</title>
        <option name="charting.chart">radialGauge</option>
    </chart>
</panel>
```

This is working as expected.

The problem starts when I try to update the chart. For example, If change the title to New Title

What I expect to get is:

```
<panel>
  <chart>
    <title>New Title</title>
    <option name="charting.chart">radialGauge</option>
  </chart>
</panel>
```

What I am getting is:

```
<panel>
  <chart>
    <title>New Title</title>
    <option name="charting.chart">radialGauge</option>
  </chart>
  <chart></chart>
</panel>
```

Note: The placement of the empty chart is arbitrary, it is either before or after the chart with values.

Moreover, I noticed that there are two charts being passed in the schema.ResourceData of the update function of this resource.

I also noticed that this problem vanishes when I remove the options map from the chart schema.

Do let me know what I am doing wrong?

Other details:

    Full Code: https://github.com/nabarunchatterjee/terraform-provider-test

    Terraform version: 0.13.2

    Golang version: 1.14.4


