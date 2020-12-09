package screenshot

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/chromedp/chromedp"
)

// Screenshot -
type Screenshot interface {
	PickImage(url string, waitforElement string, width, height int, writer io.Writer) error
}

type screenshot struct {
}

// NewScreenshot -
func NewScreenshot() Screenshot {
	return &screenshot{}
}

func (s *screenshot) PickImage(url string, waitforElement string, width, height int, writer io.Writer) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoSandbox,
		chromedp.WindowSize(width, height),
	)
	allocCtx, cannel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cannel()
	taskCtx, cannel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cannel()

	var buf []byte
	err := chromedp.Run(taskCtx, s.pickElement(url, waitforElement, &buf))
	if err != nil {
		return fmt.Errorf("截图失败: %v", err)
	}
	_, err = writer.Write(buf)
	if err != nil {
		return fmt.Errorf("write failed. %v", err)
	}
	return nil
}

func (s *screenshot) pickElement(url string, waitforElement string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(waitforElement, chromedp.ByID),
		chromedp.Screenshot(waitforElement, res, chromedp.NodeVisible, chromedp.ByID),
	}
}
