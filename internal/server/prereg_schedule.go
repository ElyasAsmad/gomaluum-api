package server

import (
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/gocolly/colly/v2"
	"github.com/mailru/easyjson"
	"github.com/nrmnqdds/gomaluum/internal/constants"
	"github.com/nrmnqdds/gomaluum/internal/dtos"
	"github.com/nrmnqdds/gomaluum/internal/errors"
	"github.com/nrmnqdds/gomaluum/internal/models"
	"github.com/nrmnqdds/gomaluum/pkg/utils"
)

func (s *Server) PreregScheduleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := &models.PreregScheduleRequest{}

	var (
		logger = s.log.GetLogger().Sugar()
		c      = colly.NewCollector(
		// colly.Async(),
		)
		scheduleUrl = constants.PreregSchedulePage
		// wg            sync.WaitGroup
		// stringBuilder strings.Builder
	)

	if err := easyjson.UnmarshalFromReader(r.Body, user); err != nil {
		logger.Errorf("Failed to unmarshal request body: %v", err)
		errors.Render(w, errors.ErrInvalidRequest)
		return
	}

	httpClient, err := CreateHTTPClient()
	if err != nil {
		logger.Errorf("Failed to create HTTP client: %v", err)
		errors.Render(w, errors.ErrFailedToCreateHTTPClient)
		return
	}

	// firstly, try to get the page to know pagination links
	// firstPageCollector := colly.NewCollector()

	// firstPageCollector.WithTransport(httpClient.Transport)

	// firstPageCollector.OnRequest(func(r *colly.Request) {
	// 	logger.Infof("Visiting URL: %v", r.URL)
	// 	r.Headers.Set("User-Agent", utils.GenerateUserAgent())
	// })

	// firstPageCollector.OnHTML("table[bgcolor='#cccccc'] tbody", func(e *colly.HTMLElement) {
	// 	links := e.ChildAttrs("tr[valign='middle'] td table tbody tr td[align='right'] a", "href")

	// })

	c.WithTransport(httpClient.Transport)

	c.OnRequest(func(r *colly.Request) {
		logger.Infof("Visiting URL: %v", r.URL)
		r.Headers.Set("User-Agent", utils.GenerateUserAgent())
	})

	c.OnHTML("table[bgcolor='#cccccc'] tbody", func(e *colly.HTMLElement) {

		// get tr entries
		e.ForEach("tr td", func(_ int, el *colly.HTMLElement) {

		})

		links := []string{}

		paginationLinks := e.ChildAttrs("tr[valign='middle'] td table tbody tr td[align='right'] a", "href")

		if len(paginationLinks) > 0 {
			// take only unique links
			links = make([]string, 0, len(paginationLinks))
			linksMap := make(map[string]struct{})
			for _, link := range paginationLinks {
				if _, exists := linksMap[link]; !exists {
					linksMap[link] = struct{}{}
					links = append(links, link)
				}
			}
			logger.Infof("Found %d unique pagination links", len(links))
		}

		for _, link := range links {
			logger.Infof("Found pagination link: %s", link)
		}
	})

	var payload = map[string]string{
		"kuly":   user.Kulliyyah,
		"sem":    fmt.Sprint(user.Semester),
		"ses":    user.Session,
		"search": "Submit",
		"ctype":  "<",
		"action": "view",
	}

	if err := c.Post(scheduleUrl, payload); err != nil {
		logger.Errorf("Failed to go to URL: %v", err)
		errors.Render(w, &errors.CustomError{
			Message:    "Failed to go to URL",
			StatusCode: 500,
		})
		return
	}

	response := &dtos.ResponseDTO{
		Message: "Successfully fetched prereg schedule",
		Data:    nil,
	}

	if err := sonic.ConfigFastest.NewEncoder(w).Encode(response); err != nil {
		logger.Errorf("Failed to encode response: %v", err)
		errors.Render(w, errors.ErrFailedToEncodeResponse)
	}
}
