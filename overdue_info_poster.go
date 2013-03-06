//写微博发出每日的逾期信息
package main

import (
	"./dao"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	url := "http://10.209.134.40/index"
	msgs := [6]string{"麻烦有空像一只喜鹊一样叽叽喳喳海森地去还一下", "麻烦你像一只小猫一样喵喵乖巧地去还一下", "麻烦有空像一只小鹿一样欢快地跑过去还一下", "麻烦像一只小兔子一样蹦蹦跳跳地开开心心过去还一下么么嗒", "麻烦你像一直小蜗牛一样背着阳光悠哉悠哉好温暖地过去还一下", "请自己去还一下"}
	infos := dao.TimeLine("datediff(now(),start_time)>15 and is_back=0 ", 0, 20)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(infos); i++ {

		text := fmt.Sprintf("@%s 同学，你借阅 @%s 大牛 的书《%s》到期了，%s! 图书圈：%s", infos[i].Borrower.Name, infos[i].Owner.Name, infos[i].BookInfo.Name, msgs[r.Intn(len(msgs))], url)
		print(text)
		print("some thing =====")
		cmd := exec.Command("sh", "post.sh", "qlxk_2004@sina.cn", "195514374", "445670032", text)
		//cmd := exec.Command("curl -u %s:%s -d \"status=%s&source=%s\" \"http://api.weibo.com/2/statuses/update.json\"", "qlxk_2004@sina.cn","195514374","tttt","445670032")
		out, _ := cmd.Output()
		print(string(out))
		time.Sleep(10 * time.Second)
		// postStatus("qlxk_2004@sina.cn","195514374","tttt","445670032")
	}
}
func postStatus(user, pwd, text, source string) {
	client := &http.Client{}
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pwd)))
	url := "http://api.weibo.com/2/statuses/update.json?source=" + source + "&status=" + text
	print(url + auth)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", "Basic "+auth)
	req.ParseForm()
	print(req.Form)
	resp, _ := client.Do(req)
	print("-------over--------")
	print(resp)
}
