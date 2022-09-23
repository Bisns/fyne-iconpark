package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Bisns/fyne-iconpark/iconpark/Base"
	"github.com/Bisns/fyne-iconpark/internal/cmd/fyne-iconpark/theme"
)

func main() {
	a := app.NewWithID("fyne-iconpark")
	t := a.Preferences().StringWithFallback("Theme", "Light")
	a.Settings().SetTheme(&theme.MyTheme{Theme: t})
	a.SetIcon(theme.Ico)
	w := a.NewWindow("Fyne IconPark 图标库")
	w.Resize(fyne.NewSize(680, 600))

	// 按 esc 关闭 dialog
	w.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case fyne.KeyEscape:
			if len(w.Canvas().Overlays().List()) > 0 {
				w.Canvas().Overlays().Top().Hide()
			}
		}
	})

	Categorys := []string{"基础", "安全 & 防护", "办公文档", "编辑", "表情", "测量 & 试验", "抽象图形", "电商财产", "动物", "多媒体音乐", "服饰", "符号标识", "工业", "化妆美妆", "几何图形", "建筑", "箭头方向", "交流沟通", "交通旅游", "界面组件", "链接", "美颜调整", "母婴儿童", "能源 & 生命", "品牌", "生活", "时间日期", "食品", "手势动作", "数据", "数据图表", "体育运动", "天气", "星座", "医疗健康", "硬件", "用户人名", "游戏", "其它"}
	tabsIcon := container.NewDocTabs(tabItem(w, Categorys)...)
	tabsIcon.SetTabLocation(container.TabLocationLeading)

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("图标", Base.HomeIcon(), container.NewPadded(tabsIcon)),
		container.NewTabItemWithIcon("设置", Base.SettingIcon(), container.NewPadded(settingsWindow())),
	)

	w.SetContent(tabs)

	w.CenterOnScreen()
	w.ShowAndRun()
}

func tabItem(w fyne.Window, tabName []string) []*container.TabItem {
	var tabItems []*container.TabItem
	for _, n := range tabName {
		buttonList := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)))
		for _, b := range genButton(w, n) {
			buttonList.Add(b)
		}

		tabItems = append(tabItems, container.NewTabItem(n, container.NewScroll(buttonList)))
	}
	return tabItems
}

func genButton(w fyne.Window, category string) []*widget.Button {
	var buttons []*widget.Button
	for _, icon := range getIcon(category) {
		d := icon // 重新赋值，不然变量是最后一个没有变化

		b := widget.NewButtonWithIcon("", d.Resource, func() {
			img := canvas.NewImageFromResource(d.Resource)
			img.SetMinSize(fyne.NewSize(200, 200))

			dialog.ShowCustomConfirm("预览图标 "+d.FuncName, "复制", "取消", container.NewCenter(img), func(isCopy bool) {
				if isCopy {
					go w.Clipboard().SetContent(d.FuncName)
				}
			}, w)
		})
		buttons = append(buttons, b)
	}
	return buttons
}

func getIcon(category string) []Icon {
	var iconsC []Icon
	for _, icon := range icons {
		i := icon
		if i.Category == category {
			iconsC = append(iconsC, i)
		}
	}
	return iconsC
}

func settingsWindow() fyne.CanvasObject {
	themeText := canvas.NewText("主题设置", nil)
	dropdown := widget.NewSelect([]string{"Light", "Dark"}, parseTheme())
	themeSett := fyne.CurrentApp().Preferences().StringWithFallback("Theme", "Light")
	switch themeSett {
	case "Light":
		dropdown.PlaceHolder = "Light"
	case "Dark":
		dropdown.PlaceHolder = "Dark"
	}

	dropdown.Refresh()

	settings := container.NewVBox(themeText, dropdown)
	return settings
}

func parseTheme() func(string) {
	return func(t string) {
		switch t {
		case "Light":
			fyne.CurrentApp().Preferences().SetString("Theme", "Light")
			fyne.CurrentApp().Settings().SetTheme(&theme.MyTheme{Theme: "Light"})
		case "Dark":
			fyne.CurrentApp().Preferences().SetString("Theme", "Dark")
			fyne.CurrentApp().Settings().SetTheme(&theme.MyTheme{Theme: "Dark"})
		}
	}
}
