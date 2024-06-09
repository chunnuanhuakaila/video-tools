package bilibili

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"video-tools/internal/consts"
	"video-tools/internal/svc"
	"video-tools/internal/types"
	"video-tools/internal/utils"
)

const (
	UserAgent   = "Mozilla/5.0 (Android 10; Mobile; rv:88.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.181 Mobile Safari/537.36"
	CidApi      = "https://api.bilibili.com/x/web-interface/view?bvid=%s"
	ContentType = "application/json"
	PlayApi     = "https://api.bilibili.com/x/player/playurl?otype=json&fnver=0&fnval=3&player=3&qn=64&bvid=%s&cid=%d&platform=html5&high_quality=1"
)

type Parser struct {
	url    string
	bvid   string
	svcCtx *svc.ServiceContext
}

func NewParser(svcCtx *svc.ServiceContext, url string) *Parser {
	return &Parser{url: url, svcCtx: svcCtx}
}

func (p *Parser) Parse() (*types.ParseResult, error) {
	p.getBvid()
	if p.bvid == "" {
		return nil, errors.New("bvid is empty")
	}
	baseInfo, err := p.getBaseInfo()
	if err != nil {
		return nil, err
	}
	playInfos, err := p.getAllPlayInfo(baseInfo.Data.Pages)
	if err != nil {
		return nil, err
	}
	return p.parseResult(baseInfo, playInfos)
}

func (p *Parser) getBaseInfo() (*BaseResponse, error) {
	cidApi := fmt.Sprintf(CidApi, p.bvid)
	header := map[string]string{"User-Agent": UserAgent, "Content-Type": ContentType, "Cookie": p.svcCtx.Config.BiliBiliCookie}
	data, err := utils.SendRequest(cidApi, header, nil)
	if err != nil {
		return nil, err
	}
	var baseInfo BaseResponse
	err = json.Unmarshal(data, &baseInfo)
	if err != nil {
		return nil, err
	}
	if baseInfo.Code != 0 {
		return nil, errors.New(baseInfo.Message)
	}
	return &baseInfo, nil
}

func (p *Parser) getBvid() {
	if strings.Contains(p.url, ".tv") {
		loc, err := utils.GetHeadersLocation(p.url)
		if err != nil {
			return
		}
		p.url = loc.URL.String()
	}
	u, err := url.Parse(p.url)
	if err != nil {
		return
	}
	s := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(s) > 0 {
		p.bvid = s[len(s)-1]
	}
}

func (p *Parser) getAllPlayInfo(pages []Page) ([]*PlayResponse, error) {
	wg := sync.WaitGroup{}
	resultChan := make(chan *PlayResponse, len(pages))
	for _, page := range pages {
		wg.Add(1)
		_page := page
		go func() {
			defer wg.Done()
			playInfo, err := p.getPlayInfo(_page)
			if err != nil {
				return
			}
			resultChan <- playInfo
		}()
	}
	wg.Wait()
	close(resultChan)
	var playInfos []*PlayResponse
	for result := range resultChan {
		playInfos = append(playInfos, result)
	}
	return playInfos, nil
}

func (p *Parser) getPlayInfo(page Page) (*PlayResponse, error) {
	header := map[string]string{"User-Agent": UserAgent, "Content-Type": ContentType, "Cookie": p.svcCtx.Config.BiliBiliCookie}
	playApi := fmt.Sprintf(PlayApi, p.bvid, page.Cid)
	data, err := utils.SendRequest(playApi, header, nil)
	if err != nil {
		return nil, err
	}
	var playInfo PlayResponse
	err = json.Unmarshal(data, &playInfo)
	if err != nil {
		return nil, err
	}
	if playInfo.Code != 0 {
		return nil, errors.New(playInfo.Message)
	}
	playInfo.Data.Durl[0].Page = page.Page
	return &playInfo, nil
}

// 解析结果
func (p *Parser) parseResult(baseInfo *BaseResponse, playInfos []*PlayResponse) (*types.ParseResult, error) {
	var (
		playMap = []string{}
	)
	for _, playInfo := range playInfos {
		if len(playInfo.Data.Durl) > 0 {
			playMap = append(playMap, playInfo.Data.Durl[0].Url)
		}
	}

	return &types.ParseResult{
		Data: &types.Data{
			Author: baseInfo.Data.Owner.Name,
			Avatar: baseInfo.Data.Owner.Face,
			Time:   utils.TimeStampToTime(baseInfo.Data.Pubdate, consts.TimeLayout),
			Title:  baseInfo.Data.Title,
			Cover:  baseInfo.Data.Pic,
			Url:    strings.Join(playMap, ","),
		},
	}, nil

}
