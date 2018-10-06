package html

import (
	"fmt"
	"strings"
)

func VideosDiv(videoUrls []string, detailUrls []string, col int) string {
	div := ""
	videos := make([]string, len(videoUrls))
	videoTemplate := `<td width='%f%%'><video  src="%s" controls></video></td>`
	videoDetailTemplate := `<td width='%f%%'><a href="%s" target="_blank"><video  src="%s" controls></video></a></td>`
	videoCount := len(videoUrls)
	var width float64
	width = 100.0 / float64(col)
	for i, url := range videoUrls {
		var video string
		if detailUrls == nil {
			video = fmt.Sprintf(videoTemplate, width, url)
		} else {
			video = fmt.Sprintf(videoDetailTemplate, width, detailUrls[i], url)
		}
		videos[i] = video
	}
	videoRowTemp := make([]string, col)
	videoRow := make([]string, 0)
	for i := 0; i < videoCount; i += col {
		for j := 0; j < col; j++ {
			index := i + j
			if index < videoCount {
				videoRowTemp[j] = videos[index]
			} else {
				videoRowTemp[j] = ""
			}
		}
		row := fmt.Sprintf("<tr>%s</tr>", strings.Join(videoRowTemp, "\n"))
		videoRow = append(videoRow, row)
	}
	div = fmt.Sprintf("<div><table border='2'>%s</table></div>", strings.Join(videoRow, "\n"))
	return div
}

func Html(title string, divs string) string {
	htmlTemplate := `
<html>
<head>
<title> %s </title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<style>
body
  {
  background-color:gray;
  }
video
{
	max-width:100%%
}
</style>
</head>
<body>
%s
</body>
</html>`
	html := fmt.Sprintf(htmlTemplate, title, divs)
	return html
}
