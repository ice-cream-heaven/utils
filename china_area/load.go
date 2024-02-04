package china_area

import (
	"github.com/ice-cream-heaven/log"
	"github.com/ice-cream-heaven/utils/common"
	"io"
	"net/http"
	"os"
)

var RemoteUrl string

var ChinaArea = NewAreaManage()

func download() error {
	client := &http.Client{}
	defer client.CloseIdleConnections()

	resp, err := client.Get(RemoteUrl)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("status code:%d", resp.StatusCode)
		return common.NewError(resp.StatusCode)
	}

	file, err := os.OpenFile("./resources/china_area.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}

func Load() error {
	return ChinaArea.Load()
}

func SearchProvince(province string) (*Area, bool) {
	return ChinaArea.SearchProvince(province)
}
