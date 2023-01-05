package collect

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/mattmattox/supportability-collector/modules/logging"
)

var log = logging.SetupLogging()

func CollectData() {

	// Create temporary directory for data collection
	tempDirRoot := CreateTmpDir()

	// Create timestamp file
	TimestampFile(tempDirRoot)

	CollectRancherData(tempDirRoot)
	CollectUpstreamCluster(tempDirRoot)

	// Tar up the temporary directory
	log.Infoln("Tarring up temporary directory")
	tarFile := tempDirRoot + ".tar.gz"
	err := TarGz(tempDirRoot, tarFile)
	if err != nil {
		log.Fatalln("Temporary directory tar failed")
	}
	log.Infoln("Successfully tarred up temporary directory")
	log.Infoln("Tar: " + tarFile)

	// Clean up temporary directory
	log.Infoln("Cleaning up temporary directory")
	err = os.RemoveAll(tempDirRoot)
	if err != nil {
		log.Fatalln("Temporary directory cleanup failed")
	}
	log.Infoln("Temporary directory cleanup successful")

	//Upload tar file to S3
	log.Infoln("Uploading tar file to S3")
	err = UploadToS3(tarFile)
	if err != nil {
		log.Fatalln("Tar file upload to S3 failed")
	}
	log.Infoln("Tar file upload to S3 successful")
}

func CreateTmpDir() string {
	log.Infoln("Creating temporary directory for data collection")
	tempDirRoot, err := os.MkdirTemp("", "supportability-")
	if err != nil {
		log.Fatalln("Temporary directory creation failed")
	}
	log.Infoln("Temporary directory created successfully")
	log.Infoln("Temporary directory: " + tempDirRoot)
	return tempDirRoot
}

func TimestampFile(dir string) {
	timestampFile := dir + "/timestamp"
	f, err := os.Create(timestampFile)
	if err != nil {
		log.Fatalln("Timestamp file creation failed")
	}
	fmt.Fprint(f, time.Now().Unix())
	f.Close()
	log.Infoln("Timestamp file created successfully")
}

func TarGz(src string, dst string) error {
	// Create a new file for the tar.gz archive
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Create a new gzip writer
	gz := gzip.NewWriter(out)
	defer gz.Close()

	// Create a new tar writer
	tw := tar.NewWriter(gz)
	defer tw.Close()

	// Walk the source directory and add each file to the tar.gz archive
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}

		// Create a new tar header for the current file
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// Set the name of the current file in the tar header
		header.Name = path

		// Write the tar header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// Open the current file
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		// Copy the contents of the file to the tar.gz archive
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}
		return nil
	})
}

func UploadToS3(tarFile string) error {
	// TODO
	return nil
}
