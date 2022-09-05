package main

import (
	"flag"
	"log"
	"time"

	aw "github.com/deanishe/awgo"
)

type WorkflowConfig struct {
	SpaceID     string `env:"SPACE_ID"`
	Cookie      string `env:"COOKIE"`
	MaxCacheAge int    `env:"MAX_CACHE_AGE"`
}

// Workflow is the main API
var (
	cacheName   = "notion-cache"
	maxCacheAge = 180 //180 * time.Minute // How long to cache repo list for
	wfConfig    *WorkflowConfig
	wf          *aw.Workflow
)

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

func run() {
	wf.Args() // call to handle magic actions
	wfConfig = &WorkflowConfig{
		MaxCacheAge: maxCacheAge,
	}

	//si pas de cookie ou de spaceid alors error
	// Update config from environment variables
	if err := wf.Config.To(wfConfig); err != nil {
		panic(err)
	}

	log.Printf("loaded: %#v", wfConfig)

	if wfConfig.SpaceID == "" {
		wf.WarnEmpty("Config error", "SpaceID is empty")
		return
	}

	if wfConfig.Cookie == "" {
		wf.WarnEmpty("Config error", "Cookie is empty")
		return
	}

	flag.Parse()
	query := flag.Arg(0)
	if query == "" {
		wf.WarnEmpty("No query", "")
		return
	}
	log.Printf("query: %#v", query)

	links := []*AlfredLink{}
	cacheKey := cacheName + "-" + query
	if wf.Cache.Exists(cacheKey) {
		log.Printf("cache hit")
		if err := wf.Cache.LoadJSON(cacheKey, &links); err != nil {
			wf.FatalError(err)
		}
	}

	cacheAge := time.Duration(wfConfig.MaxCacheAge) * time.Minute
	log.Printf("cacheAge: %#v", cacheAge.String())

	if wf.Cache.Expired(cacheKey, cacheAge) {
		log.Printf("cache Expired")
		var err error
		links, err = callNotion(query, wfConfig.SpaceID, wfConfig.Cookie)
		if err != nil {
			wf.FatalError(err)
		}

		if len(links) > 0 {
			wf.Cache.StoreJSON(cacheKey, links)
		}
	}

	if len(links) == 0 {
		wf.NewItem("Open Notion - No results, empty query, or error").Arg(NotionUrl).Valid(true)
	}

	// Add results for cached repos
	for _, r := range links {
		wf.NewItem(r.Title).
			Subtitle(r.Subtitle).
			Arg(r.Link).
			UID(r.UID).
			Valid(true)
	}
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
