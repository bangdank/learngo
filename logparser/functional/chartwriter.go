package main

// func chartWriter(w io.Writer) outputFn {
// 	return func(res []result) error {
// 		return chartWrite(w, res)
// 	}
// }

// func chartWrite(w io.Writer, res []result) error {
// 	sort.Slice(res, func(i, j int) bool {
// 		return res[i].domain > res[j].domain
// 	})

// 	donut := chart.DonutChart{
// 		Title: "Total Visits Per Domain",
// 		TitleStyle: chart.Style{
// 			FontSize:  35,
// 			Show:      true,
// 			FontColor: chart.ColorAlternateGreen,
// 		},
// 		Width:  1920,
// 		Height: 800,
// 	}

// 	for _, r := range res {
// 		v := chart.Value{
// 			Label: r.domain + r.page + ": " + strconv.Itoa(r.visits),
// 			Value: float64(r.visits),
// 			Style: chart.Style{
// 				FontSize: 14,
// 			},
// 		}

// 		donut.Values = append(donut.Values, v)
// 	}

// 	return donut.Render(chart.SVG, w)
// }
