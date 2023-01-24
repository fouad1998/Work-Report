package reporter

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

var headers = [][]string{
	{"Total jours du mois", "=C8+C16+C24+C32+C40"},
	{"Total heures du mois", "=C9+C17+C25+C33+C41"},
	{"Taux jours €", "100"},
	{"Taux horaire €", "=C4/8"},
	{"Total mois €", "=C5*C3"},
}

const layout = "2006-01-02"
const SHEET_NAME = "sheet1"

func (r *Reporter) Excel(date time.Time) error {
	if err := r.read(); err != nil {
		return err
	}

	file := excelize.NewFile()
	style, err := styles(file)
	if err != nil {
		return err
	}

	for i, head := range headers {
		file.SetCellValue(SHEET_NAME, "B"+strconv.Itoa(i+2), head[0])
		file.SetCellFormula(SHEET_NAME, "C"+strconv.Itoa(i+2), head[1])

		file.SetCellStyle(SHEET_NAME, "B"+strconv.Itoa(i+2), "B"+strconv.Itoa(i+2), style.header)
		file.SetCellStyle(SHEET_NAME, "C"+strconv.Itoa(i+2), "C"+strconv.Itoa(i+2), style.header)

	}

	var _, month, _ = date.Date()
	var heightIndex = len(headers) + 1
	var insertRows int
	for i := 0; i < 31; i++ {
		start := date.Add(time.Duration(i * 24 * int(time.Hour)))
		if start.Weekday() == time.Saturday || start.Weekday() == time.Sunday {
			continue
		}

		if start.Month() != month {
			break
		}

		if insertRows%5 == 0 {
			heightIndex += 2
			file.SetCellValue(SHEET_NAME, "B"+strconv.Itoa(heightIndex), "Jours")
			file.SetCellStyle(SHEET_NAME, "B"+strconv.Itoa(heightIndex), "B"+strconv.Itoa(heightIndex), style.subheader)

			cells := "C" + strconv.Itoa(heightIndex+2) + ":C" + strconv.Itoa(heightIndex+6)
			file.SetCellFormula(SHEET_NAME, "C"+strconv.Itoa(heightIndex), "=(SUM("+cells+")/8*4)/4")
			file.SetCellStyle(SHEET_NAME, "C"+strconv.Itoa(heightIndex), "C"+strconv.Itoa(heightIndex), style.subheader)

			heightIndex += 1
			file.SetCellValue(SHEET_NAME, "B"+strconv.Itoa(heightIndex), "Heures")
			file.SetCellFormula(SHEET_NAME, "C"+strconv.Itoa(heightIndex), "=SUM("+cells+")")

			file.SetCellStyle(SHEET_NAME, "B"+strconv.Itoa(heightIndex), "B"+strconv.Itoa(heightIndex), style.subheader)
			file.SetCellStyle(SHEET_NAME, "C"+strconv.Itoa(heightIndex), "C"+strconv.Itoa(heightIndex), style.subheader)
		}

		var item Item = Item{}
		for _, i := range r.Items {
			if i.Date == start.Format("2006-01-02") {
				item = i
				break
			}
		}

		heightIndex += 1
		insertRows += 1
		file.SetCellValue(SHEET_NAME, "B"+strconv.Itoa(heightIndex), start.Format(layout))
		file.SetCellStyle(SHEET_NAME, "B"+strconv.Itoa(heightIndex), "B"+strconv.Itoa(heightIndex), style.content)

		file.SetCellValue(SHEET_NAME, "C"+strconv.Itoa(heightIndex), item.Hours)
		file.SetCellStyle(SHEET_NAME, "C"+strconv.Itoa(heightIndex), "C"+strconv.Itoa(heightIndex), style.content)

		var report []excelize.RichTextRun

		if len(item.Meets) > 0 {
			report = append(report, excelize.RichTextRun{
				Text: "Participated meets:\n",
				Font: &excelize.Font{
					Bold:      true,
					Underline: "double",
					Size:      12,
				},
			})
			for _, e := range item.Meets {
				report = append(report, excelize.RichTextRun{
					Text: e + "\n",
					Font: &excelize.Font{
						Size: 11,
					},
				})
			}

			report = append(report, excelize.RichTextRun{
				Text: strings.Repeat("\n", 1),
			})
		}

		if len(item.Tasks.Done) > 0 {
			report = append(report, excelize.RichTextRun{
				Text: "Completed tasks:\n",
				Font: &excelize.Font{
					Bold:      true,
					Underline: "double",
					Size:      12,
				},
			})

			for _, t := range item.Tasks.Done {
				report = append(report, excelize.RichTextRun{
					Text: t + "\n",
					Font: &excelize.Font{
						Size: 11,
					},
				})
			}

			report = append(report, excelize.RichTextRun{
				Text: strings.Repeat("\n", 1),
			})
		}

		if len(item.Tasks.Progress) > 0 {
			report = append(report, excelize.RichTextRun{
				Text: "Pending tasks:\n",
				Font: &excelize.Font{
					Bold:      true,
					Underline: "double",
					Size:      12,
				},
			})
			for _, t := range item.Tasks.Progress {
				report = append(report, excelize.RichTextRun{
					Text: t + "\n",
					Font: &excelize.Font{
						Size: 11,
					},
				})
			}

			report = append(report, excelize.RichTextRun{
				Text: strings.Repeat("\n", 1),
			})
		}

		if len(item.Contributions) > 0 {
			report = append(report, excelize.RichTextRun{
				Text: "Gitlab:\n",
				Font: &excelize.Font{
					Bold:      true,
					Underline: "double",
					Size:      12,
				},
			})
			for _, t := range item.Contributions {
				report = append(report, excelize.RichTextRun{
					Text: t.Action + ": ",
					Font: &excelize.Font{
						Size: 11,
						Bold: true,
					},
				})

				report = append(report, excelize.RichTextRun{
					Text: t.Name + "\n",
					Font: &excelize.Font{
						Size: 11,
					},
				})
			}

			report = append(report, excelize.RichTextRun{
				Text: strings.Repeat("\n", 1),
			})
		}

		if item.Note != "" {
			report = append(report, excelize.RichTextRun{
				Text: "Other tasks:\n",
				Font: &excelize.Font{
					Bold:      true,
					Underline: "double",
					Size:      12,
				},
			})
			report = append(report, excelize.RichTextRun{
				Text: item.Note,
				Font: &excelize.Font{
					Size: 11,
				},
			})
		}

		file.SetCellRichText(SHEET_NAME, "D"+strconv.Itoa(heightIndex), report)
		file.SetCellStyle(SHEET_NAME, "D"+strconv.Itoa(heightIndex), "B"+strconv.Itoa(heightIndex), style.content)
		file.SetRowHeight(SHEET_NAME, heightIndex, 50)
	}

	file.SetColWidth(SHEET_NAME, "D", "D", 50)
	file.SetColWidth(SHEET_NAME, "B", "C", 20)
	if err := file.SaveAs("report.xlsx"); err != nil {
		return err
	}

	filePath, _ := os.Getwd()
	filePath = "file://" + filePath
	filePath += "/report.xlsx"
	exec.Command("xdg-open", filePath).Start()

	return nil
}
