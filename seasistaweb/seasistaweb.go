package seasistaweb

import (
	"net/http"
	"text/template"
)

type MasterPage struct {
	Title             string
	View              string
	CssUrls           []string
	Javascripts       []string
	MasterJavascripts []string
	ContentHTML       string
}

type IndexPage struct {
	Title string
	Host  string
}

type GalleryImage struct {
	Name  string
	Title string
	Link  string
}

type Gallery struct {
	Title  string
	Menu   string
	Images []*GalleryImage
}

type Galleries struct {
	Title     string
	Galleries []*Gallery
}

func (g *Galleries) Add(hi *Gallery) {
	if g.Galleries == nil {
		g.Galleries = make([]*Gallery, 0, 4)
	}
	n := len(g.Galleries)
	if n+1 > cap(g.Galleries) {
		s := make([]*Gallery, n, 2*n+1)
		copy(s, g.Galleries)
		g.Galleries = s
	}
	g.Galleries = g.Galleries[0 : n+1]
	g.Galleries[n] = hi
}

func (g *Gallery) Add(hi *GalleryImage) {
	if g.Images == nil {
		g.Images = make([]*GalleryImage, 0, 4)
	}
	n := len(g.Images)
	if n+1 > cap(g.Images) {
		s := make([]*GalleryImage, n, 2*n+1)
		copy(s, g.Images)
		g.Images = s
	}
	g.Images = g.Images[0 : n+1]
	g.Images[n] = hi
	g.Images[n].Title = g.Title
}

func (ip *IndexPage) IsView(view string) (r bool) {
	return true
}

const viewPath = len("/views/")

func init() {
	http.HandleFunc("/", home)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index", root)
	http.HandleFunc("/samples", root)
	http.HandleFunc("/gallery", gallery)
	http.HandleFunc("/shop", root)
	http.HandleFunc("/about", root)
}

func getPageData(pageName string) (data interface{}) {
	pageData := &IndexPage{Title: pageName}
	return pageData
}

func root(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	if len(title) == 0 {
		title = "index"
	}
	pageData := getPageData(title)
	var pageScripts = []string{"viewmodels/" + title + "ViewModel"}
	pm := parseWithMasterPage(title, "master", pageData, pageScripts)
	w.Write(pm)

}

func gallery(w http.ResponseWriter, r *http.Request) {
	pageData := &Galleries{Title: "gallery"}

	pageData.Add(&Gallery{Title: "All Occasions", Menu: "beach"})
	pageData.Galleries[0].Add(&GalleryImage{Name: "005", Link: "http://www.etsy.com/listing/89866511/frangipani-greeting-card"})
	pageData.Galleries[0].Add(&GalleryImage{Name: "099", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	pageData.Galleries[0].Add(&GalleryImage{Name: "003", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	pageData.Galleries[0].Add(&GalleryImage{Name: "098", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	pageData.Galleries[0].Add(&GalleryImage{Name: "010", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	pageData.Galleries[0].Add(&GalleryImage{Name: "012", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	pageData.Galleries[0].Add(&GalleryImage{Name: "089", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	pageData.Galleries[0].Add(&GalleryImage{Name: "019", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	pageData.Galleries[0].Add(&GalleryImage{Name: "014", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	pageData.Galleries[0].Add(&GalleryImage{Name: "018", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	//pageData.Galleries[0].Add( &GalleryImage{Name:"098",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	//pageData.Galleries[0].Add( &GalleryImage{Name:"099",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	//pageData.Galleries[0].Add( &GalleryImage{Name:"100",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	//pageData.Galleries[0].Add( &GalleryImage{Name:"064",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	//pageData.Galleries[0].Add( &GalleryImage{Name:"065",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	//pageData.Galleries[0].Add( &GalleryImage{Name:"066",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	//pageData.Galleries[0].Add( &GalleryImage{Name:"067",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})
	//pageData.Galleries[0].Add( &GalleryImage{Name:"068",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10727682"})

	pageData.Add(&Gallery{Title: "Birthday", Menu: "birthday"})
	pageData.Galleries[1].Add(&GalleryImage{Name: "006", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"027",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"029",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"039",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"040",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"041",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"042",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"043",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"044",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"045",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"046",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"047",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"048",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"049",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	pageData.Galleries[1].Add(&GalleryImage{Name: "061", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	pageData.Galleries[1].Add(&GalleryImage{Name: "062", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	pageData.Galleries[1].Add(&GalleryImage{Name: "092", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"076",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"077",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"078",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"079",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"080",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	//pageData.Galleries[1].Add( &GalleryImage{Name:"088",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	pageData.Galleries[1].Add(&GalleryImage{Name: "131", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})
	pageData.Galleries[1].Add(&GalleryImage{Name: "132", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10888176"})

	pageData.Add(&Gallery{Title: "Love", Menu: "lovehearts"})
	pageData.Galleries[2].Add(&GalleryImage{Name: "050", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	pageData.Galleries[2].Add(&GalleryImage{Name: "090", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	pageData.Galleries[2].Add(&GalleryImage{Name: "095", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	pageData.Galleries[2].Add(&GalleryImage{Name: "089", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	pageData.Galleries[2].Add(&GalleryImage{Name: "026", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	pageData.Galleries[2].Add(&GalleryImage{Name: "051", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"053",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"054",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"069",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"072",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"073",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"075",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"089",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"090",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"091",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"093",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"094",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})
	//pageData.Galleries[2].Add( &GalleryImage{Name:"095",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10723117"})

	pageData.Add(&Gallery{Title: "New Baby", Menu: "baby"})
	pageData.Galleries[3].Add(&GalleryImage{Name: "007", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10844893"})
	pageData.Galleries[3].Add(&GalleryImage{Name: "008", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10844893"})
	pageData.Galleries[3].Add(&GalleryImage{Name: "009", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10844893"})
	pageData.Galleries[3].Add(&GalleryImage{Name: "030", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10844893"})
	pageData.Galleries[3].Add(&GalleryImage{Name: "031", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10844893"})
	pageData.Galleries[3].Add(&GalleryImage{Name: "032", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10844893"})
	pageData.Galleries[3].Add(&GalleryImage{Name: "033", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10844893"})
	pageData.Galleries[3].Add(&GalleryImage{Name: "034", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10844893"})
	pageData.Galleries[3].Add(&GalleryImage{Name: "035", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10844893"})

	// pageData.Add( &Gallery{Title:"Easter", Menu:"easter"})
	// pageData.Galleries[4].Add( &GalleryImage{Name:"000",Link:""})

	pageData.Add(&Gallery{Title: "St Patricks Day", Menu: "stpatrick"})
	pageData.Galleries[4].Add(&GalleryImage{Name: "036", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10950250"})
	pageData.Galleries[4].Add(&GalleryImage{Name: "037", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10950250"})
	pageData.Galleries[4].Add(&GalleryImage{Name: "038", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10950250"})

	pageData.Add(&Gallery{Title: "Easter", Menu: "chickenstar"})
	pageData.Galleries[5].Add(&GalleryImage{Name: "100", Link: ""})
	pageData.Galleries[5].Add(&GalleryImage{Name: "117", Link: ""})
	pageData.Galleries[5].Add(&GalleryImage{Name: "120", Link: ""})

	pageData.Add(&Gallery{Title: "Christmas", Menu: "xmas"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "015", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "016", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "021", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "022", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "023", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "024", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "025", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "055", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "056", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "057", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})
	pageData.Galleries[6].Add(&GalleryImage{Name: "086", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10723085"})

	pageData.Add(&Gallery{Title: "Photographic Prints", Menu: "baths"})
	//pageData.Galleries[7].Add( &GalleryImage{Name:"001",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "002", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "003", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "010", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "146", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "012", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	//pageData.Galleries[7].Add( &GalleryImage{Name:"017",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	//pageData.Galleries[7].Add( &GalleryImage{Name:"014",Link:"http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "004", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "059", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "060", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "005", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "082", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "097", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})
	pageData.Galleries[7].Add(&GalleryImage{Name: "098", Link: "http://www.etsy.com/shop/seasistaphotography?section_id=10797466"})

	tp := &templateParser{}
	t, _ := template.ParseFiles("views/gallery.html")
	t.Execute(tp, pageData)
	tt := []byte(tp.HTML)
	w.Write(tt)
}

func home(w http.ResponseWriter, r *http.Request) {
	pageData := &Gallery{Title: "index"}

	pageData.Add(&GalleryImage{Name: "012", Link: ""})
	pageData.Add(&GalleryImage{Name: "011", Link: ""})
	pageData.Add(&GalleryImage{Name: "150", Link: ""})
	pageData.Add(&GalleryImage{Name: "081", Link: ""})
	pageData.Add(&GalleryImage{Name: "096", Link: ""})
	pageData.Add(&GalleryImage{Name: "151", Link: ""})
	pageData.Add(&GalleryImage{Name: "010", Link: ""})
	pageData.Add(&GalleryImage{Name: "017", Link: ""})
	pageData.Add(&GalleryImage{Name: "082", Link: ""})
	pageData.Add(&GalleryImage{Name: "097", Link: ""})
	pageData.Add(&GalleryImage{Name: "157", Link: ""})
	pageData.Add(&GalleryImage{Name: "143", Link: ""})
	pageData.Add(&GalleryImage{Name: "060", Link: ""})

	//pageData.Add( &GalleryImage{Name:"012",Link:""})
	//pageData.Add( &GalleryImage{Name:"011",Link:""})
	//pageData.Add( &GalleryImage{Name:"051",Link:""})
	//pageData.Add( &GalleryImage{Name:"096",Link:""})
	//pageData.Add( &GalleryImage{Name:"014",Link:""})
	//pageData.Add( &GalleryImage{Name:"010",Link:""})
	//pageData.Add( &GalleryImage{Name:"017",Link:""})
	//pageData.Add( &GalleryImage{Name:"018",Link:""})
	//pageData.Add( &GalleryImage{Name:"020",Link:""})
	//pageData.Add( &GalleryImage{Name:"022",Link:""})
	//pageData.Add( &GalleryImage{Name:"023",Link:""})
	//pageData.Add( &GalleryImage{Name:"052",Link:""})

	tp := &templateParser{}
	t, _ := template.ParseFiles("views/home.html")
	t.Execute(tp, pageData)
	tt := []byte(tp.HTML)
	w.Write(tt)

}

func parseWithMasterPage(viewName string, masterPageName string, data interface{}, javascripts []string) []byte {
	// Build the content
	context := parseView(viewName, data)
	// Add the content to the master page
	mp := &MasterPage{Title: "seasista photography", ContentHTML: string(context), View: viewName}

	// Add the css templates
	var cssUrls = []string{"site"}
	mp.CssUrls = cssUrls

	// Add the javascripts, page and master scripts
	mp.Javascripts = javascripts
	var masterJavascripts = []string{"jquery-1.7.1.min"}
	mp.MasterJavascripts = masterJavascripts

	// Build the page
	tp := &templateParser{}
	t, _ := template.ParseFiles("views/" + masterPageName + ".html")
	t.Execute(tp, mp)
	return []byte(tp.HTML)
}

func parseView(viewName string, data interface{}) []byte {
	pageName := "views/" + viewName + ".html"
	tp := &templateParser{}
	t, _ := template.ParseFiles(pageName)
	t.Execute(tp, data)
	return []byte(tp.HTML)
}

type templateParser struct {
	HTML string
}

func (tP *templateParser) Write(p []byte) (n int, err error) {
	tP.HTML += string(p)
	return len(p), nil
}
