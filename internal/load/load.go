package load

import (
	"errors"
	"fmt"
	"strings"
	"video-tools/internal/parser"
	"video-tools/internal/parser/bilibili"
	"video-tools/internal/parser/douyin"
	"video-tools/internal/parser/kuaishou"
	"video-tools/internal/parser/pipixia"
	"video-tools/internal/parser/weishi"
	"video-tools/internal/parser/xiaohongshu"
	"video-tools/internal/parser/xigua"
	"video-tools/internal/parser/zuiyou"
	"video-tools/internal/svc"
)

func LoadParser(svcCtx *svc.ServiceContext, url string) (parser.Parser, error) {
	url = strings.Trim(url, " ")
	fmt.Println(url)
	switch {
	case strings.Contains(url, "weishi"):
		return weishi.NewParser(url), nil
	case strings.Contains(url, "pipix"):
		return pipixia.NewParser(url), nil
	case strings.Contains(url, "douyin"):
		return douyin.NewParser(url), nil
	case strings.Contains(url, "zuiyou") || strings.Contains(url, "xiaochuankeji"):
		return zuiyou.NewParser(url), nil
	case strings.Contains(url, "kuaishou"):
		return kuaishou.NewParser(svcCtx, url), nil
	case strings.Contains(url, "ixigua"):
		return xigua.NewParser(url), nil
	case strings.Contains(url, "xhslink.com") || strings.Contains(url, "xiaohongshu.com"):
		return xiaohongshu.NewParser(url), nil
	case strings.Contains(url, "b23.tv") || strings.Contains(url, "bilibili.com"):
		return bilibili.NewParser(svcCtx, url), nil
	default:
		return nil, errors.New("not support")
	}
}
