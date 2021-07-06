package dragon

import (
	"io"
	"log"
	"net/http"
	"os"
)

//suggest OSS (object storage service). file upload
func Upload(r *http.Request, file string, saveTo string) error {
	srcFile, _, err := r.FormFile(file)
	defer srcFile.Close()

	if err != nil {
		log.Println(err)
		return err
	}

	dstFile, err := os.Create(saveTo)
	defer dstFile.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
