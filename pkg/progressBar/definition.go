package progressBar

import (
	"errors"
	"fmt"
	"github.com/Renjie-Woo/Doraemon/utils"
	"log"
)

type ProgressBar struct {
	title                              string  // 进度条标题
	rateGraph                          string  // 进度条图标
	totalCount                         float64 // 数据总量
	currentCount                       float64 // 当前完成量
	percent                            float64 // 完成比例
	currentBar                         string  // 当前条形图
	unit                               string  // 单位
	initErr                            error
	_hasTotalOrCurrentCountInitialized bool
}

func (pb *ProgressBar) SetTitle(title string) *ProgressBar {
	if utils.IsStringEmpty(title) {
		title = defaultTitle
	}
	pb.title = title

	return pb
}

func (pb *ProgressBar) GetTitle() string {
	return pb.title
}

func (pb *ProgressBar) SetGraph(graph string) *ProgressBar {
	if utils.IsStringEmpty(graph) {
		graph = defaultGraph
	}
	pb.rateGraph = graph

	return pb
}

func (pb *ProgressBar) GetGraph() string {
	return pb.rateGraph
}

// SetTotalCount
// support all number format such as int,int32,int64,uint,uint32,uint64,float,float32,float64, and their string ones
func (pb *ProgressBar) SetTotalCount(total float64) *ProgressBar {
	if total <= 0 {
		pb.initErr = fmt.Errorf(invalidTotalCountError)
		return pb
	}
	pb.totalCount = total
	if pb._hasTotalOrCurrentCountInitialized && pb.totalCount < pb.currentCount {
		pb.initErr = errors.New(fmt.Sprintf(currentCountGreaterThanTotalOneError, pb.currentCount, pb.totalCount))
	} else {
		pb._hasTotalOrCurrentCountInitialized = true
	}
	//return nil
	return pb
}

// SetCurrentCount
// support all number format such as int,int32,int64,uint,uint32,uint64,float,float32,float64, and their string ones
func (pb *ProgressBar) SetCurrentCount(current float64) *ProgressBar {
	if current < 0 {
		pb.initErr = fmt.Errorf(invalidCurrentCountError)
		return pb
	}
	pb.currentCount = current
	if pb._hasTotalOrCurrentCountInitialized && pb.totalCount < pb.currentCount {
		pb.initErr = errors.New(fmt.Sprintf(currentCountGreaterThanTotalOneError, pb.currentCount, pb.totalCount))
	} else {
		pb._hasTotalOrCurrentCountInitialized = true
	}
	return pb
}

func (pb *ProgressBar) setPercent(percent float64) *ProgressBar {
	pb.percent = percent

	return pb
}

func (pb *ProgressBar) getPercent() float64 {
	return pb.percent
}

func (pb *ProgressBar) setCurrentBar() string {
	var currentPercent = pb.getPercent()
	var newestPercent = pb.currentCount / pb.totalCount
	increase := int(newestPercent*100) - int(currentPercent*100)
	pb.setPercent(newestPercent)
	for i := 0; i < increase; i++ {
		pb.currentBar += pb.GetGraph()
	}
	return pb.currentBar
}

func (pb *ProgressBar) SetUnit(uint string) *ProgressBar {
	pb.unit = uint

	return pb
}

func NewProgressBar(title string, current, total float64) ProgressBar {
	var bar = ProgressBar{}
	bar.SetTitle(title)
	bar.SetCurrentCount(current)
	bar.SetTotalCount(total)
	bar.SetGraph(defaultGraph)
	if bar.initErr != nil {
		log.Fatalf("please check your bar settings: %v\n", bar.initErr)
	}
	return bar
}

func (pb *ProgressBar) Run(current float64) {
	pb.SetCurrentCount(current)
	currentBar := pb.setCurrentBar()
	rate := pb.getPercent() * 100
	strRate := utils.ParseFloatToStringWithAccuracy(rate, 2)
	strCurrent := utils.ParseFloatToStringWithAccuracy(pb.currentCount, 1)
	strTotal := utils.ParseFloatToStringWithAccuracy(pb.totalCount, 1)
	fmt.Printf("\r%s: [%-100s]%8s%%  %8s %s/%s %s", pb.title, currentBar, strRate, strCurrent, pb.unit, strTotal, pb.unit)
	if pb.currentCount == pb.totalCount {
		fmt.Println()
	}
}
