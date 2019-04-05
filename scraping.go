package main
import (
  "fmt"
  "log"
  "net/url"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  // スクレイピング対象のURL
  target := "http://localhost:8080/"

  // HTMLを取得する
  doc, err := goquery.NewDocument(target)
  if err != nil {
    log.Fatal("[Error] access failed\n", err)
  }

  // URLを分析する
  baseUrl, _ := url.Parse(target)

  // スライス宣言
  var titles []string
  var urls []string

   // タイトルとURL取得、スライスに格納する
  doc.Find("body #conL a").Each(func(i int, s *goquery.Selection) {
    // aタグのテキスト、hrefをそれぞれ取得する
    title := s.Text()
    href, _ := s.Attr("href")

    // 相対URLから絶対URLへ変換する
    absUrl := convertUrl(baseUrl, href)

    // スライスに格納する
    titles = append(titles, title)
    urls = append(urls, absUrl)
  })

  // タイトルとURLを出力する
  fmt.Println("Title:\n", titles, "\n")
  fmt.Println("URL:\n", urls)
}

func convertUrl(baseUrl *url.URL, webUrl string) string {
  // hrefの属性値を解析する
  targetUrl, err := url.Parse(webUrl)
  if err != nil {
    return ""
  }

  // hrefの属性値が相対URLだった場合、絶対URLへ変換する
  absUrl := baseUrl.ResolveReference(targetUrl)

  // 絶対URLを返却する
  return absUrl.String()
}