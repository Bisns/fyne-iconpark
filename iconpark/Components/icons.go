// AUTO-GENERATED: DO NOT EDIT

package Components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	carousel     *theme.ThemedResource
	checklist    *theme.ThemedResource
	pagetemplate *theme.ThemedResource
	page         *theme.ThemedResource
	platte       *theme.ThemedResource
	switchbutton *theme.ThemedResource
	tablefile    *theme.ThemedResource
	treelist     *theme.ThemedResource
)

func init() {
	carousel = theme.NewThemedResource(resourceCarouselSvg)
	checklist = theme.NewThemedResource(resourceChecklistSvg)
	pagetemplate = theme.NewThemedResource(resourcePageTemplateSvg)
	page = theme.NewThemedResource(resourcePageSvg)
	platte = theme.NewThemedResource(resourcePlatteSvg)
	switchbutton = theme.NewThemedResource(resourceSwitchButtonSvg)
	tablefile = theme.NewThemedResource(resourceTableFileSvg)
	treelist = theme.NewThemedResource(resourceTreeListSvg)
}

// CarouselIcon returns Carousel icon resource
func CarouselIcon() fyne.Resource {
	return carousel
}

// ChecklistIcon returns Checklist icon resource
func ChecklistIcon() fyne.Resource {
	return checklist
}

// PagetemplateIcon returns Pagetemplate icon resource
func PagetemplateIcon() fyne.Resource {
	return pagetemplate
}

// PageIcon returns Page icon resource
func PageIcon() fyne.Resource {
	return page
}

// PlatteIcon returns Platte icon resource
func PlatteIcon() fyne.Resource {
	return platte
}

// SwitchbuttonIcon returns Switchbutton icon resource
func SwitchbuttonIcon() fyne.Resource {
	return switchbutton
}

// TablefileIcon returns Tablefile icon resource
func TablefileIcon() fyne.Resource {
	return tablefile
}

// TreelistIcon returns Treelist icon resource
func TreelistIcon() fyne.Resource {
	return treelist
}
