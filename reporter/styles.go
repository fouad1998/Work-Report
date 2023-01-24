package reporter

import "github.com/xuri/excelize/v2"

type style struct {
	header,
	subheader,
	content int
}

func styles(file *excelize.File) (*style, error) {
	header, err := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 13,
		},
	})
	if err != nil {
		return nil, err
	}

	subheader, err := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		return nil, err
	}

	content, err := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Fill: excelize.Fill{
			Pattern: 1,
			Type:    "pattern",
			Color:   []string{"#fff2cc"},
		},
	})
	if err != nil {
		return nil, err
	}

	return &style{
		header:    header,
		subheader: subheader,
		content:   content,
	}, nil
}
